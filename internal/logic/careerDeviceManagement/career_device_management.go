package careerDeviceManagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/career"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"sms_backend/utility"
)

func New() *sCareerDeviceManagement {
	return &sCareerDeviceManagement{}
}

func init() {
	service.RegisterCareerDeviceManagement(New())
}

type sCareerDeviceManagement struct{}

const careerLogger = "career"

// DeviceRegister handles the registration of a new device.
// It checks if the device is already registered in the database or exists in the cache,
// and if not, registers it in the database and removes it from the cache.
//
// Parameters:
// - ctx: The context for handling the operation.
// - req: The request payload containing the device details (phone number, device number, and active time).
//
// Returns:
// - *career.RegisterRes: The response containing the registered device ID.
// - error: An error if the operation fails at any point.
func (s *sCareerDeviceManagement) DeviceRegister(ctx context.Context, req *career.RegisterReq) (*career.RegisterRes, error) {
	// Prepare the device entity for registration
	device := entity.DeviceList{
		Number:       req.PhoneNumber,
		DeviceNumber: req.DeviceNumber,
		ActiveTime:   req.ActiveTime,
		DeviceStatus: 2, // Normal
	}

	// Check if the device is already registered in the database
	count, err := dao.DeviceList.Ctx(ctx).Where("device_number = ?", device.DeviceNumber).Count()
	if err != nil {
		g.Log(careerLogger).Errorf(ctx, "Failed to check device registration for device_number '%s': %v", device.DeviceNumber, err)
		return nil, fmt.Errorf("failed to check device registration: %w", err)
	}
	if count > 0 {
		g.Log(careerLogger).Warningf(ctx, "Device already registered with device_number '%s'", device.DeviceNumber)
		return nil, errors.New("device already registered")
	}

	// Check if the device exists in the unregistered cache and remove it if found
	if err := utility.CleanUpDeviceCache(ctx, utility.CachePrefixUnregisteredDevice, device.DeviceNumber); err != nil {
		g.Log(careerLogger).Errorf(ctx, "Failed to clean up unregistered device cache for device_number '%s': %v", device.DeviceNumber, err)
		return nil, err
	}

	// Register the device in the database
	rowID, err := dao.DeviceList.Ctx(ctx).Data(device).InsertAndGetId()
	if err != nil {
		g.Log(careerLogger).Errorf(ctx, "Failed to register device in the database for device_number '%s': %v", device.DeviceNumber, err)
		return nil, fmt.Errorf("failed to register device: %w", err)
	}

	// Return the registration response
	g.Log(careerLogger).Infof(ctx, "Successfully registered device with ID '%d' and device_number '%s'", rowID, device.DeviceNumber)
	return &career.RegisterRes{
		ID: rowID,
	}, nil
}

type FileData struct {
	TargetPhoneNumber []string `json:"target_phone_number"`
	Content           []string `json:"content"`
	DeviceNumber      []string `json:"device_number"`
}

//var FetchTaskLen = g.Cfg().MustGet(context.Background(), "ChanLen.FetchTask").Int()
//var SEMAPHORE = make(chan struct{}, FetchTaskLen)

func (s *sCareerDeviceManagement) FetchTasks(ctx context.Context, req *career.FetchTaskReq) (*career.FetchTaskRes, error) {
	// Check if the device is already marked as unregistered in cache or Redis
	if err := checkDeviceStatusInCache(ctx, utility.CachePrefixUnregisteredDevice, req.DeviceNumber); err != nil {
		return nil, err
	}
	if err := checkDeviceStatusInCache(ctx, utility.CachePrefixUnassignedDevice, req.DeviceNumber); err != nil {
		return nil, err
	}

	// Send device payload to the processing channel
	devicePayload := &utility.SmsDeviceNumberPayload{DeviceNumber: req.DeviceNumber}
	select {
	case utility.SmsDeviceNumberPayloadChan <- devicePayload:
		g.Log(careerLogger).Infof(ctx, "Device payload sent to processing channel: %s", req.DeviceNumber)
	default:
		return nil, fmt.Errorf("failed to send device payload to channel for device: %s", req.DeviceNumber)
	}

	// Wait for the response from the task payload channel
	select {
	case resChan := <-utility.SmsFetchTaskPayloadChan:
		if resChan.Err != nil {
			g.Log(careerLogger).Errorf(ctx, "Failed to fetch tasks for device '%s': %v", req.DeviceNumber, resChan.Err)
			return nil, resChan.Err
		}
		g.Log(careerLogger).Infof(ctx, "Successfully fetched tasks for device '%s'", req.DeviceNumber)
		return resChan.Res, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("context canceled while waiting for task response for device: %s", req.DeviceNumber)
	}
}

// checkDeviceStatusInCache checks if a device is marked as unregistered or unassigned in the cache or Redis.
//
// Parameters:
// - ctx: The context for handling the operation.
// - cachePrefix: The prefix to distinguish between unregistered and unassigned devices.
// - deviceNumber: The device number to check.
//
// Returns:
// - error: An error if the device is already marked or if there is an issue during the check.
func checkDeviceStatusInCache(ctx context.Context, cachePrefix, deviceNumber string) error {
	cacheKey := cachePrefix
	hashKey := cacheKey + deviceNumber

	// Check existence in cache or Redis
	exists, err := utility.CheckMapExistsByCacheAndRedis(ctx, cacheKey, hashKey)
	if err != nil {
		g.Log(careerLogger).Errorf(ctx, "Failed to check device status in cache or Redis (cacheKey: %s, hashKey: %s): %v", cacheKey, hashKey, err)
		return fmt.Errorf("failed to check device status: %w", err)
	}

	// Log a warning if the device is already marked
	if exists {
		g.Log(careerLogger).Warningf(ctx, "Device already marked in cache or Redis (cacheKey: %s, hashKey: %s)", cacheKey, hashKey)
		return fmt.Errorf("device already marked for cacheKey: %s, hashKey: %s", cacheKey, hashKey)
	}

	return nil
}

func (s *sCareerDeviceManagement) ReportTaskResult(ctx context.Context, req *career.ReportTaskResultReq) (*career.ReportTaskResultRes, error) {
	// Validate the SendStatus
	if err := validateSendStatus(ctx, req.SendStatus, req.TaskId); err != nil {
		return nil, err
	}

	// Create a new SmsTaskPayload and send it to the channel
	taskPayload := utility.NewSmsTaskPayload(req.TaskId, req.TargetPhoneNumber, req.DeviceNumber, req.Content, req.SendStatus, req.StartTime, req.Reason)
	if err := sendToTaskChannel(ctx, taskPayload); err != nil {
		return nil, err
	}

	// Log success and return a response
	g.Log(careerLogger).Infof(ctx, "Successfully processed TaskId=%d for DeviceNumber=%s", req.TaskId, req.DeviceNumber)
	return &career.ReportTaskResultRes{}, nil
}

// Helper: Validate the SendStatus
func validateSendStatus(ctx context.Context, sendStatus int, taskId int) error {
	if !utility.SentStatus(sendStatus).IsValid() {
		g.Log(careerLogger).Errorf(ctx, "Invalid SendStatus=%d for TaskId=%d", sendStatus, taskId)
		return errors.New("SendStatus 验证错误，请提交合法的 SendStatus 值")
	}
	return nil
}

// Helper: Send the task payload to the channel
func sendToTaskChannel(ctx context.Context, taskPayload *utility.SmsTaskPayload) error {
	select {
	case utility.SmsTaskPayloadChan <- taskPayload:
		g.Log(careerLogger).Infof(ctx, "Task payload sent to channel successfully: TaskId=%d", taskPayload.TaskID)
		return nil
	case <-ctx.Done():
		g.Log(careerLogger).Errorf(ctx, "Failed to send task payload to channel: TaskId=%d, context canceled", taskPayload.TaskID)
		return fmt.Errorf("failed to send task payload to channel: TaskId=%d, context canceled", taskPayload.TaskID)
	default:
		g.Log(careerLogger).Errorf(ctx, "Task payload channel is full, failed to send TaskId=%d", taskPayload.TaskID)
		return fmt.Errorf("task payload channel is full, failed to send TaskId=%d", taskPayload.TaskID)
	}
}

// ReportReceiveContent handles the received SMS content by adding it to the processing queue
// and waiting for the result from the result channel.
//
// Parameters:
// - ctx: The context for handling the operation.
// - req: The request containing the received SMS content details.
//
// Returns:
// - *career.ReportReceiveContentRes: The result of the SMS processing.
// - error: An error if the processing fails.
func (s *sCareerDeviceManagement) ReportReceiveContent(ctx context.Context, req *career.ReportReceiveContentReq) (*career.ReportReceiveContentRes, error) {
	// Prepare the SMS payload for processing
	payload := &utility.ReceivedSmsPayload{
		DeviceNumber:      req.DeviceNumber,
		TargetPhoneNumber: req.TargetPhoneNumber,
		SmsContent:        req.SmsContent,
		StartTime:         req.StartTime,
	}

	// Send the payload to the SMS processing channel
	select {
	case utility.ReceivedSmsPayloadChan <- payload:
		g.Log(careerLogger).Infof(ctx, "Received SMS payload added to processing queue: %+v", payload)
	case <-ctx.Done():
		g.Log(careerLogger).Warning(ctx, "Context canceled before payload could be added to the processing queue.")
		return nil, fmt.Errorf("context canceled before adding payload to queue")
	}

	// Wait for the result from the result channel
	select {
	case resp := <-utility.ReportReceiveContentResPayloadChan:
		// Log and return the result
		if resp.Err != nil {
			g.Log(careerLogger).Errorf(ctx, "Error processing SMS payload: %+v, error: %v", payload, resp.Err)
		} else {
			g.Log(careerLogger).Infof(ctx, "Successfully processed SMS payload: %+v", payload)
		}
		return resp.Res, resp.Err

	case <-ctx.Done():
		g.Log(careerLogger).Warning(ctx, "Context canceled while waiting for processing result.")
		return nil, fmt.Errorf("context canceled while waiting for processing result")
	}
}
