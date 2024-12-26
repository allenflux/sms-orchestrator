package utility

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/career"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
)

const FetchTaskFilterPrefix = "FetchTaskFilterPrefix"
const CachePrefixSmsTraceTask = "RedisPrefixSmsTraceTask"
const CachePrefixUnregisteredDevice = "CachePreUnregisteredDevice"
const CachePrefixUnassignedDevice = "CachePrefixUnassignedDevice"
const CachePrefixNoExecutableTask = "CachePrefixNoExecutableTask"
const fetchTaskQueueLogger = "fetchTaskQueue"

type SmsDeviceNumberPayload struct {
	DeviceNumber string `json:"device_number"` // 设备编号
}
type SmsFetchTaskPayload struct {
	Res *career.FetchTaskRes
	Err error
}

var SmsDeviceNumberPayloadChan = make(chan *SmsDeviceNumberPayload, 100)
var SmsFetchTaskPayloadChan = make(chan *SmsFetchTaskPayload, 100)

func ConsumerFetchTaskQueue(ctx context.Context, taskChan <-chan *SmsDeviceNumberPayload) {
	for {
		select {
		case taskData, ok := <-taskChan: // 从管道中获取任务
			if !ok {
				g.Log(fetchTaskQueueLogger).Info(ctx, "FetchTaskQueue channel is closed, stopping consumer")
				return
			}
			// 处理任务 taskID int, targetPhoneNumber, deviceNumber, content string, sendStatus int, startTime *gtime.Time, reason string
			var res *career.FetchTaskRes
			var err error
			if res, err = FetchTaskQueue(ctx, taskData.DeviceNumber); err != nil {
				g.Log(fetchTaskQueueLogger).Errorf(ctx, "FetchTaskQueue Error: %v", err)
			} else {

				g.Log(fetchTaskQueueLogger).Infof(ctx, "Task FetchTaskQueue successfully processed: %+v", taskData)
			}
			payload := &SmsFetchTaskPayload{Res: res, Err: err}
			SmsFetchTaskPayloadChan <- payload

		case <-ctx.Done(): // 监听 context 的取消信号
			g.Log(fetchTaskQueueLogger).Info(ctx, "FetchTaskQueue stopped due to context cancellation")
			return
		}
	}
}

// FetchTaskQueue fetches task queue based on device number.
func FetchTaskQueue(ctx context.Context, deviceNumber string) (*career.FetchTaskRes, error) {
	// Fetch the device information
	device, err := fetchDevice(ctx, deviceNumber)
	if err != nil {
		// Cache unregistered device to prevent duplicate processing
		if cacheErr := CacheDeviceStatus(ctx, CachePrefixUnregisteredDevice, deviceNumber); cacheErr != nil {
			g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to cache unregistered device '%s': %v", deviceNumber, cacheErr)
			return nil, cacheErr
		}
		return nil, err
	}

	// Check if the device exists in the unregistered cache and remove it if found
	if err := CleanUpDeviceCache(ctx, CachePrefixUnregisteredDevice, device.DeviceNumber); err != nil {
		g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to clean up unregistered device cache for device_number '%s': %v", device.DeviceNumber, err)
		return nil, err
	}

	// Check if the device is assigned to any group
	if device.GroupId == 0 {
		// Cache unassigned device to prevent duplicate processing
		if cacheErr := CacheDeviceStatus(ctx, CachePrefixUnassignedDevice, deviceNumber); cacheErr != nil {
			g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to cache unregistered device '%s': %v", deviceNumber, cacheErr)
			return nil, cacheErr
		}
		return nil, fmt.Errorf("device %s not assigned to any group", deviceNumber)
	}

	// Check if the device exists in the unassigned cache and remove it if found
	if err := CleanUpDeviceCache(ctx, CachePrefixUnassignedDevice, device.DeviceNumber); err != nil {
		g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to clean up unassigned device cache for device_number '%s': %v", device.DeviceNumber, err)
		return nil, err
	}

	// Handle tasks for the device
	taskRes, err := handleTasks(ctx, device)
	if err != nil {
		g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to handle tasks for device '%s': %v", deviceNumber, err)
		return nil, err
	}
	// Update the device status to "Occupied" (2)
	if err := updateDeviceStatus(ctx, deviceNumber, 2); err != nil {
		g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to update device status to 'Occupied' for DeviceNumber='%s': %v", deviceNumber, err)
		return nil, fmt.Errorf("failed to update device status for DeviceNumber='%s': %w", deviceNumber, err)
	}

	return taskRes, nil
}

// CacheDeviceStatus caches the status of a device in Redis and local cache.
// This is a generic function to cache various device states like "unregistered" or "unassigned".
//
// Parameters:
// - ctx: The context for handling the operation.
// - prefix: The cache prefix to distinguish the type of status (e.g., "unregistered" or "unassigned").
// - deviceNumber: The device number to be cached.
//
// Returns:
// - error: An error if the caching operation fails.
func CacheDeviceStatus(ctx context.Context, prefix, deviceNumber string) error {
	cacheKey := prefix + deviceNumber
	if err := SetMapValueToCacheAndRedis(ctx, prefix, cacheKey, 1); err != nil {
		g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to set cache for device '%s' with prefix '%s': %v", deviceNumber, prefix, err)
		return fmt.Errorf("failed to cache device '%s' with prefix '%s': %w", deviceNumber, prefix, err)
	}
	g.Log().Infof(ctx, "Successfully cached device '%s' with prefix '%s'", deviceNumber, prefix)
	return nil
}

func fetchDevice(ctx context.Context, deviceNumber string) (*entity.DeviceList, error) {
	var device entity.DeviceList
	count := 0
	err := dao.DeviceList.Ctx(ctx).Where("device_number = ?", deviceNumber).ScanAndCount(&device, &count, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve device info for DeviceNumber=%s: %w", deviceNumber, err)
	}
	if count == 0 {
		return nil, fmt.Errorf("no device found for DeviceNumber=%s", deviceNumber)
	}
	return &device, nil
}

func handleTasks(ctx context.Context, device *entity.DeviceList) (*career.FetchTaskRes, error) {
	// 优先查看对话任务
	if taskRes, err := checkRedisForTasks(ctx, device.DeviceNumber); err != nil || taskRes != nil {
		return taskRes, err
	}
	// 处理从数据库中获取的任务
	return handleDatabaseTasks(ctx, device)
}

func checkRedisForTasks(ctx context.Context, deviceNumber string) (*career.FetchTaskRes, error) {
	redisTaskCount, err := g.Redis().LLen(ctx, deviceNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get Redis task length for DeviceNumber=%s: %w", deviceNumber, err)
	}
	if redisTaskCount > 0 {
		return fetchTaskFromRedis(ctx, deviceNumber)
	}
	return nil, nil
}

// 直接从redis返回来自对话的任务
func fetchTaskFromRedis(ctx context.Context, deviceNumber string) (*career.FetchTaskRes, error) {
	messageData, err := g.Redis().LPop(ctx, deviceNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to pop task from Redis for DeviceNumber=%s: %w", deviceNumber, err)
	}
	var subMessageData sms.SubPostConversationRecordData
	if err = messageData.Scan(&subMessageData); err != nil {
		return nil, fmt.Errorf("failed to map Redis data for DeviceNumber=%s: %w", deviceNumber, err)
	}
	return constructTaskResponse(&subMessageData, deviceNumber), nil
}

func handleDatabaseTasks(ctx context.Context, device *entity.DeviceList) (*career.FetchTaskRes, error) {
	// Fetch job from the database
	job, err := fetchJobFromDB(ctx, device.GroupId)
	if err != nil {
		return nil, err
	}

	// Generate cache keys
	//listKey, cacheKey, hashKey := generateCacheKeys(job.FileName, job.SurplusQuantity)
	listKey, _, _ := generateCacheKeys(job.FileName, job.SurplusQuantity)

	// Check if the task already exists in cache or Redis
	//if exists, err := checkTaskExistence(ctx, cacheKey, hashKey); err != nil {
	//	return nil, err
	//} else if exists {
	//	return nil, fmt.Errorf("task already exists for cacheKey: %s, hashKey: %s", cacheKey, hashKey)
	//}

	// Fetch the task value from cache or Redis
	//value, err := RetrieveValueFromCacheAndRedis(ctx, listKey, job.SurplusQuantity-1)
	value, err := GetValueFromCacheOrRedisQueue(ctx, listKey)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve value for listKey: %s, index: %d: %w", listKey, job.SurplusQuantity-1, err)
	}
	if value == nil {
		return nil, errors.New("empty value retrieved from cache or Redis")
	}

	// Parse the task value
	subMessageData, err := ParseCacheValueToRecord(value)
	if err != nil {
		return nil, fmt.Errorf("failed to parse cache value for listKey: %s: %w", listKey, err)
	}

	// Update cache to prevent duplicate processing
	//if err := SetMapValueToCacheAndRedis(ctx, cacheKey, hashKey, 1); err != nil {
	//	Debug(ctx, err.Error(), "fetch_task_err.txt")
	//}

	//Interval:          subMessageData.Interval,
	//TaskId:            subMessageData.TaskID,
	//StartAt:           subMessageData.StartAt,
	subMessageData.Interval = job.IntervalTime
	subMessageData.StartAt = job.StartTime
	// Construct the task response
	taskRes := constructTaskResponse(subMessageData, device.DeviceNumber)

	// Begin a database transaction to handle task updates
	if err := processTaskTransaction(ctx, job, taskRes, listKey); err != nil {
		return nil, err
	}

	return taskRes, nil
}

func generateCacheKeys(fileName string, surplusQuantity int) (listKey, cacheKey, hashKey string) {
	index := surplusQuantity - 1
	if index < 0 {
		index = 0
	}
	listKey = fileName
	cacheKey = FetchTaskFilterPrefix + fileName
	hashKey = fmt.Sprintf("fetch_task %d", index)
	return
}

// updateDeviceStatus updates the sent_status of a device.
//
// Parameters:
// - ctx: The context for handling the operation.
// - deviceNumber: The unique identifier of the device.
// - status: The new sent status (1 - Abnormal, 2 - Occupied, 3 - Idle).
//
// Returns:
// - error: An error if the operation fails.
func updateDeviceStatus(ctx context.Context, deviceNumber string, status int) error {
	// Validate the sent_status
	// Validate the sent_status
	if status < 1 || status > 3 {
		return fmt.Errorf("invalid sent_status '%d' for DeviceNumber='%s'", status, deviceNumber)
	}

	// Update the sent_status in the database
	if _, err := dao.DeviceList.Ctx(ctx).
		Data(g.Map{"sent_status": status}).
		Where("device_number = ?", deviceNumber).
		Update(); err != nil {
		g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to update sent_status to '%d' for DeviceNumber='%s': %v", status, deviceNumber, err)
		return fmt.Errorf("failed to update sent_status to '%d' for DeviceNumber='%s': %w", status, deviceNumber, err)
	}

	g.Log(fetchTaskQueueLogger).Infof(ctx, "Successfully updated sent_status to '%d' for DeviceNumber='%s'", status, deviceNumber)
	return nil
}

func constructTaskResponse(subMessageData *sms.SubPostConversationRecordData, deviceNumber string) *career.FetchTaskRes {
	return &career.FetchTaskRes{
		TargetPhoneNumber: subMessageData.TargetPhoneNumber,
		Content:           subMessageData.Content,
		DeviceNumber:      deviceNumber,
		Interval:          subMessageData.Interval,
		TaskId:            subMessageData.TaskID,
		StartAt:           subMessageData.StartAt,
	}
}

func processTaskTransaction(ctx context.Context, job *entity.SmsMissionReport, taskRes *career.FetchTaskRes, listKey string) error {
	db := g.DB()
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			g.Log(fetchTaskQueueLogger).Errorf(ctx, "Transaction rolled back due to panic: %v", p)
		}
	}()

	// Update the task status and surplus quantity
	surplusQuantity := job.SurplusQuantity - 1
	if surplusQuantity < 0 {
		return errors.New("remaining task quantity cannot be negative")
	}
	updateData := g.Map{
		"task_status":      2,
		"surplus_quantity": surplusQuantity,
	}
	if surplusQuantity == 0 {
		updateData["task_status"] = 3
	}
	if _, err := tx.Model("sms_mission_report").Ctx(ctx).
		Data(updateData).
		Where("id = ?", job.Id).
		Update(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update task status: %w", err)
	}

	// Clear cache if surplus quantity is zero
	if surplusQuantity == 0 {
		if err := clearCache(ctx, listKey); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Add task tracking
	if err := handleTaskTracking(ctx, tx, taskRes); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to handle task tracking: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}

func clearCache(ctx context.Context, listKey string) error {
	if err := RemoveCacheAndRedisKey(ctx, listKey); err != nil {
		g.Log(fetchTaskQueueLogger).Errorf(ctx, "Failed to clear cache for listKey: %s: %v", listKey, err)
		return fmt.Errorf("failed to clear cache for listKey: %w", err)
	}
	//if err := RemoveCacheAndRedisKey(ctx, cacheKey); err != nil {
	//	g.Log().Errorf(ctx, "Failed to clear cache for cacheKey: %s: %v", cacheKey, err)
	//	return fmt.Errorf("failed to clear cache for cacheKey: %w", err)
	//}
	return nil
}

func fetchJobFromDB(ctx context.Context, groupId int) (*entity.SmsMissionReport, error) {
	var job entity.SmsMissionReport
	count := 0
	err := dao.SmsMissionReport.Ctx(ctx).
		Where("group_id = ?", groupId).
		Where("task_status IN (?)", []int{1, 2}).
		Limit(1).
		ScanAndCount(&job, &count, false)
	if err != nil {
		return nil, fmt.Errorf("failed to query SmsMissionReport for GroupId=%d: %w", groupId, err)
	}
	if count == 0 {
		return nil, fmt.Errorf("no executable tasks found for GroupId=%d", groupId)
	}
	return &job, nil
}

func MakeNameCachePrefixSmsTraceTask(deviceNumber, targetPhoneNumber string) string {
	return CachePrefixSmsTraceTask + deviceNumber + targetPhoneNumber
}

func handleTaskTracking(ctx context.Context, tx gdb.TX, res *career.FetchTaskRes) error {
	var traceTask entity.SmsTraceTask
	var count int

	// 查询任务追踪记录
	if err := dao.SmsTraceTask.Ctx(ctx).
		Where("target_number = ?", res.TargetPhoneNumber).
		Where("device_number = ?", res.DeviceNumber).
		ScanAndCount(&traceTask, &count, false); err != nil {
		return fmt.Errorf("查询 SmsTraceTask 错误: %w", err)
	}

	if count != 0 {
		// 更新任务追踪记录
		if _, err := tx.Model("sms_trace_task").Ctx(ctx).Data("task_id = ?", res.TaskId).Update(); err != nil {
			return fmt.Errorf("更新 SmsTraceTask 错误: %w", err)
		}
	} else {
		// 插入新的任务追踪记录
		if _, err := tx.Model("sms_trace_task").Ctx(ctx).Data(entity.SmsTraceTask{
			TargetNumber: res.TargetPhoneNumber,
			DeviceNumber: res.DeviceNumber,
			TaskId:       int(res.TaskId),
		}).Insert(); err != nil {
			return fmt.Errorf("插入 SmsTraceTask 错误: %w", err)
		}
	}

	return nil
}
