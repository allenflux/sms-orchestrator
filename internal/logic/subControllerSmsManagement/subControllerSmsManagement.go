package subControllerSmsManagement

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"io/ioutil"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"sms_backend/utility"
)

func New() *sSubControllerSmsManagement {
	return &sSubControllerSmsManagement{}
}

func init() {
	service.RegisterSubControllerSmsManagement(New())
}

type sSubControllerSmsManagement struct{}

func (s *sSubControllerSmsManagement) GetTaskList(ctx context.Context, req *sms.SubTaskListReq) (res *sms.SubTaskListRes, err error) {
	sand := dao.SmsMissionReport.Ctx(ctx).Page(req.PageNum, req.PageSize).Where("associated_account_id = ?", req.SubUserID)
	if req.ProjectID != 0 {
		sand = sand.Where("project_id = ?", req.ProjectID)
	}
	if len(req.TaskName) != 0 {
		sand = sand.Where("task_name like ?", "%"+req.TaskName+"%")
	}

	if len(req.DateRange) > 0 {
		sand = sand.Where("create_at >= ? AND create_at <= ?", req.DateRange[0], req.DateRange[1])
	}
	var data []entity.SmsMissionReport
	var totalCount int
	if err = sand.ScanAndCount(&data, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询DB SmsMissionReport 错误")
	}
	res = &sms.SubTaskListRes{}
	res.Total = totalCount
	res.Data = make([]sms.SubTaskListResData, len(data))
	for i := range data {
		res.Data[i] = sms.SubTaskListResData{
			ID:                data[i].Id,
			ProjectID:         data[i].ProjectId,
			ProjectName:       data[i].ProjectName,
			TaskName:          data[i].TaskName,
			FileName:          data[i].FileName,
			DeviceQuota:       data[i].DeviceQuota,
			TaskStatus:        data[i].TaskStatus,
			SmsQuantity:       data[i].SmsQuantity,
			SurplusQuantity:   data[i].SurplusQuantity,
			QuantitySent:      data[i].QuantitySent,
			AssociatedAccount: data[i].AssociatedAccount,
			IntervalTime:      data[i].IntervalTime,
			StartTime:         data[i].StartTime.String(),
			CreateTime:        data[i].CreatedAt.String(),
		}
	}
	return
}

// Create Task

type FileData struct {
	TargetPhoneNumber []string `json:"target_phone_number"`
	Content           []string `json:"content"`
}

// TaskCreate handles the creation of a new task, including file parsing, validation, and distribution.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing task details, group ID, and the uploaded file.
//
// Returns:
// - *sms.SubTaskCreateRes: The response containing the created task ID.
// - error: An error if the operation fails.
func (s *sSubControllerSmsManagement) TaskCreate(ctx context.Context, req *sms.SubTaskCreateReq) (*sms.SubTaskCreateRes, error) {
	// Step 1: Fetch account details
	var account entity.User
	if err := dao.User.Ctx(ctx).Where("id = ?", req.SubUserId).Scan(&account); err != nil {
		g.Log().Errorf(ctx, "Failed to query User for AccountId=%d: %v", req.SubUserId, err)
		return nil, fmt.Errorf("failed to query account details: %w", err)
	}
	accountName := account.Name

	// Step 2: Validate group information
	var group entity.SubGroup
	count := 0
	if err := dao.SubGroup.Ctx(ctx).Where("id = ?", req.GroupID).ScanAndCount(&group, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query SubGroup for GroupID=%d: %v", req.GroupID, err)
		return nil, fmt.Errorf("failed to query group details: %w", err)
	} else if count == 0 {
		return nil, errors.New("invalid group ID: no matching record found")
	}

	// Step 3: Validate devices in the group
	var devices []*entity.DeviceList
	if err := dao.DeviceList.Ctx(ctx).Where("group_id = ?", group.Id).ScanAndCount(&devices, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query DeviceList for GroupID=%d: %v", group.Id, err)
		return nil, fmt.Errorf("failed to query device list: %w", err)
	} else if count == 0 {
		return nil, errors.New("no available devices in the selected group")
	}

	// Step 4: Validate project information
	var project entity.ProjectList
	if err := dao.ProjectList.Ctx(ctx).Where("id=?", group.ProjectId).ScanAndCount(&project, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query ProjectList for ProjectId=%d: %v", group.ProjectId, err)
		return nil, fmt.Errorf("failed to query project details: %w", err)
	} else if count == 0 {
		return nil, errors.New("invalid project ID: no matching record found")
	}

	// Step 5: Validate task name and file
	if count, err := dao.SmsMissionReport.Ctx(ctx).Where("task_name = ?", req.TaskName).Count(); err != nil {
		g.Log().Errorf(ctx, "Failed to validate task name '%s': %v", req.TaskName, err)
		return nil, fmt.Errorf("failed to validate task name: %w", err)
	} else if count > 0 {
		return nil, errors.New("duplicate task name")
	}

	// Save uploaded file
	filename, err := req.File.Save(consts.TaskFilePath, true)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to save file: %v", err)
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	// Validate file name uniqueness
	if count, err := dao.SmsMissionReport.Ctx(ctx).Where("file_name = ?", filename).Count(); err != nil {
		g.Log().Errorf(ctx, "Failed to validate file name '%s': %v", filename, err)
		return nil, fmt.Errorf("failed to validate file name: %w", err)
	} else if count > 0 {
		return nil, errors.New("duplicate file name")
	}

	// Step 6: Parse and validate file content
	content, err := ioutil.ReadFile(consts.TaskFilePath + "/" + filename)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to read file '%s': %v", filename, err)
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var payload FileData
	if err = json.Unmarshal(content, &payload); err != nil {
		g.Log().Errorf(ctx, "Failed to parse JSON content in file '%s': %v", filename, err)
		return nil, fmt.Errorf("failed to parse JSON content: %w", err)
	}

	if len(payload.Content) == 0 || len(payload.TargetPhoneNumber) == 0 || len(payload.Content) != len(payload.TargetPhoneNumber) {
		return nil, errors.New("invalid file format: content and phone numbers mismatch")
	}

	// Step 7: Create and save task
	task := entity.SmsMissionReport{
		ProjectId:           group.ProjectId,
		TaskName:            req.TaskName,
		GroupId:             req.GroupID,
		FileName:            filename,
		TaskStatus:          1,
		SmsQuantity:         len(payload.Content),
		SurplusQuantity:     len(payload.TargetPhoneNumber),
		QuantitySent:        0,
		AssociatedAccount:   accountName,
		IntervalTime:        req.IntervalTime,
		StartTime:           req.TimingStartTime,
		ProjectName:         project.ProjectName,
		AssociatedAccountId: req.SubUserId,
	}

	taskID, err := dao.SmsMissionReport.Ctx(ctx).Data(task).InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to insert task into DB: %v", err)
		return nil, fmt.Errorf("failed to insert task: %w", err)
	}

	// Step 8: Distribute tasks to devices
	devicesLen := len(devices)
	for i, phoneNumber := range payload.TargetPhoneNumber {
		message := &sms.SubPostConversationRecordData{
			TaskID:            taskID,
			Content:           payload.Content[i],
			DeviceNumber:      devices[i%devicesLen].DeviceNumber,
			TargetPhoneNumber: phoneNumber,
		}

		if err := utility.AddValueToCacheAndRedisQueue(ctx, filename, message); err != nil {
			g.Log().Errorf(ctx, "Failed to distribute task to Redis queue for DeviceNumber='%s': %v", message.DeviceNumber, err)
			return nil, fmt.Errorf("failed to distribute task to Redis queue: %w", err)
		}
	}

	g.Log().Infof(ctx, "Successfully created task '%s' with ID=%d", req.TaskName, taskID)
	return &sms.SubTaskCreateRes{ID: taskID}, nil
}

// Download File

func (s *sSubControllerSmsManagement) TaskFileDownload(ctx context.Context, req *sms.TaskFileDownloadReq) (res *sms.TaskFileDownloadRes, err error) {
	g.Log().Infof(ctx, "FileDownloadReq: %v", req.FileName)
	var content []byte
	if content, err = ioutil.ReadFile(consts.TaskFilePath + "/" + req.FileName); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("下载时存储的读取文件失败")
	}
	res = &sms.TaskFileDownloadRes{
		JsonData: string(content),
	}
	return
}

// TaskReportDelete handles the recall of a task by updating its status in the SmsMissionReport table.
func (s *sSubControllerSmsManagement) TaskReportDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (*sms.SubTaskDeleteRes, error) {
	var report entity.SmsMissionReport
	var count int

	// Fetch the task report based on TaskID
	if err := dao.SmsMissionReport.Ctx(ctx).
		Where("id = ?", req.TaskID).
		ScanAndCount(&report, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to fetch SmsMissionReport for TaskID=%d: %v", req.TaskID, err)
		return nil, fmt.Errorf("failed to fetch task record: %w", err)
	}

	// Check if the task exists
	if count == 0 {
		g.Log().Warningf(ctx, "Invalid TaskID=%d: No record found", req.TaskID)
		return nil, errors.New("invalid task ID: no record found")
	}

	// Check if the task is already completed
	if report.TaskStatus == 3 {
		g.Log().Warningf(ctx, "Invalid TaskID=%d: Task is already completed", req.TaskID)
		return nil, errors.New("invalid task ID: task already completed")
	}

	// Update the task status to "recalled" (4)
	if _, err := dao.SmsMissionReport.Ctx(ctx).
		Data(g.Map{"task_status": 4}).
		Where("id = ?", req.TaskID).
		Update(); err != nil {
		g.Log().Errorf(ctx, "Failed to update task status to 'recalled' for TaskID=%d: %v", req.TaskID, err)
		return nil, fmt.Errorf("failed to recall task: %w", err)
	}

	g.Log().Infof(ctx, "Successfully recalled task with TaskID=%d", req.TaskID)
	return &sms.SubTaskDeleteRes{}, nil
}

// GetSubGetConversationRecord retrieves a detailed conversation record list for a specific ChatLogID.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing the ChatLogID.
//
// Returns:
// - *sms.SubGetConversationRecordRes: The response containing the conversation record list.
// - error: An error if the operation fails.
func (s *sSubControllerSmsManagement) GetSubGetConversationRecord(ctx context.Context, req *sms.SubGetConversationRecordReq) (*sms.SubGetConversationRecordRes, error) {
	var (
		chatLog     entity.SmsChartLog
		chatLogList []*entity.SmsChartLog
	)

	// Step 1: Fetch the main chat log entry based on ChatLogID
	count := 0
	err := dao.SmsChartLog.Ctx(ctx).
		Where("id = ?", req.ChatLogID).
		ScanAndCount(&chatLog, &count, false)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to query DB SmsChartLog for ChatLogID=%d: %v", req.ChatLogID, err)
		return nil, fmt.Errorf("failed to query DB SmsChartLog: %w", err)
	}
	if count == 0 {
		g.Log().Warningf(ctx, "Invalid ChatLogID=%d: No matching record found", req.ChatLogID)
		return nil, errors.New("invalid ChatLogID: no matching record found")
	}

	// Step 2: Fetch all chat logs for the same TargetPhoneNumber and DeviceNumber
	if err := dao.SmsChartLog.Ctx(ctx).
		Where("target_phone_number = ?", chatLog.TargetPhoneNumber).
		Where("device_number = ?", chatLog.DeviceNumber).
		OrderDesc("id").
		Scan(&chatLogList); err != nil {
		g.Log().Errorf(ctx, "Failed to query related SmsChartLog records for TargetPhoneNumber=%s, DeviceNumber=%s: %v", chatLog.TargetPhoneNumber, chatLog.DeviceNumber, err)
		return nil, fmt.Errorf("failed to query related chat logs: %w", err)
	}

	// Step 3: Map the chat logs to the response structure
	data := make([]sms.SubGetConversationRecordListResData, len(chatLogList))
	for i, log := range chatLogList {
		data[i] = sms.SubGetConversationRecordListResData{
			ChatLogID:         log.Id,
			RecordTime:        log.CreatedAt,
			Content:           log.SmsContent,
			TargetPhoneNumber: log.TargetPhoneNumber,
			SentOrReceive:     log.SentOrReceive,
			DeviceNumber:      log.DeviceNumber,
		}
	}

	// Step 4: Prepare and return the response
	res := &sms.SubGetConversationRecordRes{
		Data: data,
	}

	g.Log().Infof(ctx, "Successfully fetched conversation record list for ChatLogID=%d (Total=%d)", req.ChatLogID, len(data))
	return res, nil
}

// SubGetConversationRecordList retrieves a paginated list of conversation records for a specific project.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing pagination, project ID, and optional SubUserID.
//
// Returns:
// - *sms.SubGetConversationRecordListRes: The response containing the conversation records.
// - error: An error if the operation fails.
func (s *sSubControllerSmsManagement) SubGetConversationRecordList(ctx context.Context, req *sms.SubGetConversationRecordListReq) (*sms.SubGetConversationRecordListRes, error) {
	var (
		chatLogsId []*entity.SmsChartLog
		chatLogs   []*entity.SmsChartLog
		totalCount int
	)

	// Step 1: Query unique conversation records based on target phone number and device number
	query := dao.SmsChartLog.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Where("project_id = ?", req.ProjectID).
		OrderDesc("id")

	if req.SubUserID != 0 {
		query = query.Where("associated_account_id = ?", req.SubUserID)
	}

	// Fetch the unique IDs for the latest conversation records
	if err := query.
		FieldMax("id", "id").
		Group("target_phone_number").
		Group("device_number").
		ScanAndCount(&chatLogsId, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query SmsChartLog IDs for ProjectID=%d: %v", req.ProjectID, err)
		return nil, fmt.Errorf("failed to query unique conversation record IDs: %w", err)
	}

	// Step 2: Extract IDs and prepare for fetching full records
	idList := make([]int, len(chatLogsId))
	for i, log := range chatLogsId {
		idList[i] = log.Id
	}

	// Step 3: Query full conversation records using the fetched IDs
	if err := dao.SmsChartLog.Ctx(ctx).
		WhereIn("id", idList).
		Scan(&chatLogs); err != nil {
		g.Log().Errorf(ctx, "Failed to query SmsChartLog for IDs: %v", err)
		return nil, fmt.Errorf("failed to query conversation records: %w", err)
	}

	// Step 4: Transform the query results into the response format
	data := make([]sms.SubGetConversationRecordListResData, len(chatLogs))
	for i, log := range chatLogs {
		data[i] = sms.SubGetConversationRecordListResData{
			ChatLogID:         log.Id,
			RecordTime:        log.CreatedAt,
			Content:           log.SmsContent,
			TargetPhoneNumber: log.TargetPhoneNumber,
			SentOrReceive:     log.SentOrReceive,
		}
	}

	// Step 5: Prepare the response
	res := &sms.SubGetConversationRecordListRes{
		ListRes: commonApi.ListRes{
			CurrentPage: req.PageNum,
			Total:       totalCount,
		},
		Data: data,
	}

	g.Log().Infof(ctx, "Successfully fetched %d conversation records for ProjectID=%d", len(data), req.ProjectID)
	return res, nil
}

func (s *sSubControllerSmsManagement) GetTaskRecordList(ctx context.Context, req *sms.SubTaskRecordReq) (res *sms.SubTaskRecordRes, err error) {
	sand := dao.SmsMissionRecord.Ctx(ctx).Page(req.PageNum, req.PageSize).Where("associated_account_id = ?", req.SubUserID)
	if req.ProjectID != 0 {
		sand = sand.Where("project_id = ?", req.ProjectID)
	}

	if req.SmsStatus != 0 {
		sand = sand.Where("sms_status = ?", req.SmsStatus)
	}

	if len(req.TaskName) != 0 {
		sand = sand.Where("task_name like ?", "%"+req.TaskName+"%")
	}

	if len(req.TargetPhoneNumber) != 0 {
		sand = sand.Where("target_phone_number like ?", "%"+req.TargetPhoneNumber+"%")
	}

	if len(req.DeviceNumber) != 0 {
		sand = sand.Where("device_number like ?", "%"+req.DeviceNumber+"%")
	}

	if len(req.SentDateRange) > 0 {
		sand = sand.Where("start_time >= ? AND start_time <= ?", req.SentDateRange[0], req.SentDateRange[1])
	}

	if len(req.CreateDateRange) > 0 {
		sand = sand.Where("created_at >= ? AND created_at <= ?", req.CreateDateRange[0], req.CreateDateRange[1])
	}

	var data []entity.SmsMissionRecord
	var totalCount int
	if err = sand.ScanAndCount(&data, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询DB SmsMissionRecord 错误")
	}
	res = &sms.SubTaskRecordRes{}
	res.Total = totalCount
	res.Data = make([]sms.SubTaskRecordResData, len(data))
	for i := range data {
		res.Data[i] = sms.SubTaskRecordResData{
			ID:                data[i].Id,
			TaskName:          data[i].TaskName,
			SubTaskID:         data[i].SubTaskId,
			TargetPhoneNumber: data[i].TargetPhoneNumber,
			DeviceNumber:      data[i].DeviceNumber,
			SmsTopic:          data[i].SmsTopic,
			SmsContent:        data[i].SmsContent,
			SmsStatus:         data[i].SmsStatus,
			AssociatedAccount: data[i].AssociatedAccount,
			ProjectName:       data[i].ProjectName,
			Reason:            data[i].Reason,
			StartTime:         data[i].StartTime.String(),
			CreateTime:        data[i].CreatedAt.String(),
		}
	}
	return
}

// PostConversationRecord receives input from a Sub User and stores it in the cache as a priority task for a Device to execute.
func (s *sSubControllerSmsManagement) PostConversationRecord(ctx context.Context, req *sms.SubPostConversationRecordReq) (*sms.SubPostConversationRecordRes, error) {
	// Step 1: Fetch SmsChartLog record by ChartLogID
	chartLog, err := fetchSmsChartLogByID(ctx, req.ChatLogID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch SmsChartLog: %w", err)
	}

	// Step 2: Generate task message
	message := generateTaskMessage(chartLog, req.Content)

	// Step 3: Push task to Redis queue
	unSentTaskCount, err := pushTaskToRedis(ctx, chartLog.DeviceNumber, message)
	if err != nil {
		return nil, fmt.Errorf("failed to push task to Redis queue: %w", err)
	}

	// Step 4: Construct response
	return &sms.SubPostConversationRecordRes{
		TaskItemName:  chartLog.DeviceNumber,
		UnSendTaskNum: unSentTaskCount,
	}, nil
}

// fetchSmsChartLogByID retrieves SmsChartLog record by its ID.
func fetchSmsChartLogByID(ctx context.Context, chartLogID int64) (*entity.SmsChartLog, error) {
	var chartLog entity.SmsChartLog
	count := 0
	err := dao.SmsChartLog.Ctx(ctx).
		Where("id = ?", chartLogID).
		ScanAndCount(&chartLog, &count, false)

	if err != nil {
		g.Log().Errorf(ctx, "Failed to query SmsChartLog with ID '%d': %v", chartLogID, err)
		return nil, errors.New("database query error while fetching SmsChartLog")
	}

	if count == 0 {
		return nil, errors.New("no SmsChartLog found for the given ID")
	}

	return &chartLog, nil
}

// generateTaskMessage creates a task message based on the provided chartLog and content.
func generateTaskMessage(chartLog *entity.SmsChartLog, content string) sms.SubPostConversationRecordData {
	return sms.SubPostConversationRecordData{
		TaskID:            int64(chartLog.TaskId),
		Content:           content,
		DeviceNumber:      chartLog.DeviceNumber,
		TargetPhoneNumber: chartLog.TargetPhoneNumber,
		Interval:          "0",
		StartAt:           gtime.Now(),
	}
}

// pushTaskToRedis pushes the given task message to the Redis queue associated with the device.
func pushTaskToRedis(ctx context.Context, deviceNumber string, message sms.SubPostConversationRecordData) (int64, error) {
	taskIndex, err := g.Redis().LPush(ctx, deviceNumber, message)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to push task to Redis for device '%s': %v", deviceNumber, err)
		return 0, errors.New("failed to insert task into Redis list")
	}

	g.Log().Infof(ctx, "Task successfully pushed to Redis for device '%s', Unsent Task Count: %d", deviceNumber, taskIndex)
	return taskIndex, nil
}

// GetPendingTasks retrieves the pending tasks for a specific SMS mission.
//
// Parameters:
// - ctx: The context for the request.
// - req: The request containing the TaskID.
//
// Returns:
// - *sms.SubPendingTaskRes: The response containing the list of pending tasks.
// - error: An error if the operation fails.
func (s *sSubControllerSmsManagement) GetPendingTasks(ctx context.Context, req *sms.SubPendingTaskReq) (*sms.SubPendingTaskRes, error) {
	// Fetch task details by TaskID
	var task entity.SmsMissionReport
	if err := dao.SmsMissionReport.Ctx(ctx).
		Where("id = ?", req.TaskID).
		Scan(&task); err != nil {
		g.Log().Errorf(ctx, "Failed to fetch task details for TaskID=%d: %v", req.TaskID, err)
		return nil, fmt.Errorf("failed to fetch task details: %w", err)
	}

	// Log task details
	g.Log().Infof(ctx, "FileName for TaskID=%d: %s", req.TaskID, task.FileName)

	// Read the task file from the file system
	filePath := fmt.Sprintf("%s/%s", consts.TaskFilePath, task.FileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to read file '%s': %v", filePath, err)
		return nil, errors.New("failed to read task file")
	}

	// Unmarshal the file content into a structured payload
	var payload FileData
	if err := json.Unmarshal(content, &payload); err != nil {
		g.Log().Errorf(ctx, "Failed to unmarshal file content for TaskID=%d: %v", req.TaskID, err)
		return nil, fmt.Errorf("failed to parse task file content: %w", err)
	}

	// Generate the list of pending tasks based on the surplus quantity
	length := task.SurplusQuantity
	if length > len(payload.TargetPhoneNumber) || length > len(payload.Content) {
		g.Log().Errorf(ctx, "SurplusQuantity exceeds file content length for TaskID=%d", req.TaskID)
		return nil, errors.New("invalid surplus quantity in task file")
	}

	data := make([]sms.PendingTaskResData, length)
	for i := 0; i < length; i++ {
		data[i] = sms.PendingTaskResData{
			TaskID:            req.TaskID,
			TargetPhoneNumber: payload.TargetPhoneNumber[i],
			Content:           payload.Content[i],
			StartAt:           task.StartTime,
		}
	}

	// Prepare the response
	res := &sms.SubPendingTaskRes{
		ListRes: commonApi.ListRes{
			Total: length,
		},
		Data: data,
	}

	return res, nil
}

// GetTaskDevices retrieves the list of devices assigned to a specific task by delegating the request to the main service.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing task ID and pagination details.
//
// Returns:
// - *sms.SubTaskDevicesRes: The response containing the list of devices and pagination metadata.
// - error: An error if the operation fails.
func (s *sSubControllerSmsManagement) GetTaskDevices(ctx context.Context, req *sms.SubTaskDevicesReq) (*sms.SubTaskDevicesRes, error) {
	// Prepare the main service request
	mainReq := &sms.TaskDevicesReq{
		PageReq: model.PageReq{
			PageSize: req.PageSize,
			PageNum:  req.PageNum,
		},
		TaskID: req.TaskID,
	}

	// Delegate the request to the main service
	mainRes, err := service.MainControllerSmsManagement().GetTaskDevices(ctx, mainReq)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get task devices for TaskID=%d: %v", req.TaskID, err)
		return nil, err
	}

	// Prepare the response for the sub-controller
	res := &sms.SubTaskDevicesRes{
		ListRes: commonApi.ListRes{
			Total: mainRes.Total,
		},
		Data: mainRes.Data,
	}

	return res, nil
}
