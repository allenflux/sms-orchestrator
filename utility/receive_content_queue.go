package utility

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sms_backend/api/v1/career"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"strconv"
)

type ReceivedSmsPayload struct {
	DeviceNumber      string
	TargetPhoneNumber string
	SmsContent        string
	StartTime         *gtime.Time
}

type ReportReceiveContentResPayload struct {
	Res *career.ReportReceiveContentRes
	Err error
}

const receiveTaskQueueLogger = "receiveTaskQueue"

// Define channels for received SMS tasks and their results

var ReceivedSmsPayloadChan = make(chan *ReceivedSmsPayload, 100)
var ReportReceiveContentResPayloadChan = make(chan *ReportReceiveContentResPayload, 100)

// ProcessReceivedSmsQueue continuously processes tasks from the received SMS channel.
// It listens for tasks and invokes HandleReceivedSmsMessage to process each task.
//
// Parameters:
// - ctx: The context for handling cancellation and task processing.
// - taskChan: The channel providing ReceivedSmsPayload tasks.
func ProcessReceivedSmsQueue(ctx context.Context, taskChan chan *ReceivedSmsPayload) {
	for {
		select {
		case taskData, ok := <-taskChan: // Fetch task from the channel
			if !ok {
				g.Log(receiveTaskQueueLogger).Info(ctx, "Task channel is closed. Stopping consumer.")
				return
			}

			// Process the received SMS message
			go processReceivedSmsTask(ctx, taskData)

		case <-ctx.Done(): // Handle context cancellation
			g.Log(receiveTaskQueueLogger).Info(ctx, "Consumer stopped due to context cancellation.")
			return
		}
	}
}

// processReceivedSmsTask handles the processing of a single SMS task.
//
// Parameters:
// - ctx: The context for handling the operation.
// - taskData: The payload containing the SMS task details.
func processReceivedSmsTask(ctx context.Context, taskData *ReceivedSmsPayload) {
	g.Log(receiveTaskQueueLogger).Infof(ctx, "Processing task: %+v", taskData)

	// Handle the received SMS message
	res, err := HandleReceivedSmsMessage(ctx, taskData.DeviceNumber, taskData.TargetPhoneNumber, taskData.SmsContent, taskData.StartTime)

	// Prepare the response payload
	payload := &ReportReceiveContentResPayload{
		Res: res,
		Err: err,
	}

	// Log the result of task processing
	if err != nil {
		g.Log(receiveTaskQueueLogger).Errorf(ctx, "Failed to process task for DeviceNumber=%s, TargetPhoneNumber=%s: %v", taskData.DeviceNumber, taskData.TargetPhoneNumber, err)
	} else {
		g.Log(receiveTaskQueueLogger).Infof(ctx, "Successfully processed task for DeviceNumber=%s, TargetPhoneNumber=%s", taskData.DeviceNumber, taskData.TargetPhoneNumber)
	}

	// Send the result to the result channel
	select {
	case ReportReceiveContentResPayloadChan <- payload:
		g.Log(receiveTaskQueueLogger).Info(ctx, "Result payload successfully sent to result channel.")
	case <-ctx.Done():
		g.Log(receiveTaskQueueLogger).Warning(ctx, "Failed to send result payload. Context was cancelled.")
	}
}

// HandleReceivedSmsMessage processes a received SMS message by fetching the task ID, validating the report, and saving a log entry.
//
// Parameters:
// - ctx: The context for handling the operation.
// - req: The request containing device number, target phone number, and SMS content.
//
// Returns:
// - *career.ReportReceiveContentRes: The response containing the log entry ID.
// - error: An error if any part of the operation fails.
func HandleReceivedSmsMessage(ctx context.Context, deviceNumber, targetPhoneNumber, smsContent string, startTime *gtime.Time) (*career.ReportReceiveContentRes, error) {
	// Fetch Task ID from Redis or Database
	taskId, err := fetchTaskID(ctx, deviceNumber, targetPhoneNumber)
	if err != nil {
		return nil, err
	}

	// Fetch the report from the database
	report, err := fetchReport(ctx, taskId)
	if err != nil {
		return nil, err
	}

	// Generate a unique hash to prevent duplicate entries
	rowHash := GenHash(
		report.TaskName, strconv.Itoa(taskId), targetPhoneNumber,
		deviceNumber, smsContent, "1", report.AssociatedAccount,
		strconv.Itoa(report.AssociatedAccountId), report.ProjectName,
		strconv.Itoa(report.ProjectId), startTime.String(), "2",
	)
	g.Log(receiveTaskQueueLogger).Infof(ctx, "Receive Data row hash -> %s", rowHash)

	// Check for duplicate log entries
	if exists, err := checkDuplicateLog(ctx, rowHash); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.New("current submission has duplicate records, please check if the device submitted twice")
	}

	// Save the log entry to the database
	rowID, err := saveLogEntry(ctx, report, targetPhoneNumber, deviceNumber, smsContent, rowHash)
	if err != nil {
		return nil, err
	}

	return &career.ReportReceiveContentRes{ID: rowID}, nil
}

// Helper: Fetch Task ID from Redis or Database
func fetchTaskID(ctx context.Context, deviceNumber, targetPhoneNumber string) (int, error) {
	// Try to get the task ID from Redis
	redisKey := MakeNameCachePrefixSmsTraceTask(deviceNumber, targetPhoneNumber)
	v, err := g.Redis().Get(ctx, redisKey)
	if err != nil {
		g.Log(receiveTaskQueueLogger).Error(ctx, "Failed to fetch task ID from Redis:", err)
		return 0, errors.New("error fetching task ID from Redis")
	}
	taskId := gconv.Int(v)
	if taskId != 0 {
		g.Log(receiveTaskQueueLogger).Info(ctx, "Task ID retrieved from Redis:", taskId)
		return taskId, nil
	}

	// If not in Redis, fetch from the database
	var traceData entity.SmsTraceTask
	count := 0
	err = dao.SmsTraceTask.Ctx(ctx).
		Where("device_number=?", deviceNumber).
		Where("target_number=?", targetPhoneNumber).
		ScanAndCount(&traceData, &count, false)
	if err != nil {
		g.Log(receiveTaskQueueLogger).Error(ctx, "Database query for SmsTraceTask failed:", err)
		return 0, errors.New("error querying SmsTraceTask from database")
	}
	if count == 0 {
		return 0, errors.New("no trace task found, please verify the platform task")
	}

	g.Log(receiveTaskQueueLogger).Info(ctx, "Task ID retrieved from Database:", traceData.TaskId)
	return traceData.TaskId, nil
}

// Helper: Fetch Report from the Database
func fetchReport(ctx context.Context, taskId int) (*entity.SmsMissionReport, error) {
	var report entity.SmsMissionReport
	count := 0
	err := dao.SmsMissionReport.Ctx(ctx).
		Where("id = ?", taskId).
		ScanAndCount(&report, &count, false)
	if err != nil {
		g.Log(receiveTaskQueueLogger).Error(ctx, "Failed to query SmsMissionReport:", err)
		return nil, errors.New("error querying SmsMissionReport")
	}
	if count == 0 {
		return nil, errors.New("no report found for the given task ID")
	}

	g.Log(receiveTaskQueueLogger).Info(ctx, "Report retrieved for Task ID:", taskId)
	return &report, nil
}

// Helper: Check for Duplicate Log Entries
func checkDuplicateLog(ctx context.Context, rowHash string) (bool, error) {
	count, err := dao.SmsChartLog.Ctx(ctx).
		Where("row_hash = ?", rowHash).
		Count()
	if err != nil {
		g.Log(receiveTaskQueueLogger).Error(ctx, "Failed to check duplicate log entries:", err)
		return false, errors.New("error checking for duplicate log entries")
	}
	return count > 0, nil
}

// Helper: Save Log Entry to the Database
func saveLogEntry(ctx context.Context, report *entity.SmsMissionReport, targetPhoneNumber, deviceNumber, smsContent, rowHash string) (int64, error) {
	data := entity.SmsChartLog{
		TaskId:              report.Id,
		ProjectName:         report.ProjectName,
		ProjectId:           report.ProjectId,
		TargetPhoneNumber:   targetPhoneNumber,
		DeviceNumber:        deviceNumber,
		SmsTopic:            "todo no topic for received SMS",
		SmsContent:          smsContent,
		SmsStatus:           1, // Status is always success for received messages
		AssociatedAccount:   report.AssociatedAccount,
		SentOrReceive:       2, // 2 indicates received message
		AssociatedAccountId: report.AssociatedAccountId,
		RowHash:             rowHash,
	}

	// Insert log entry and return its ID
	rowID, err := dao.SmsChartLog.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		g.Log(receiveTaskQueueLogger).Error(ctx, "Failed to insert SmsChartLog:", err)
		return 0, errors.New("error inserting log entry into SmsChartLog")
	}

	g.Log(receiveTaskQueueLogger).Info(ctx, "Log entry saved with ID:", rowID)
	return rowID, nil
}
