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

func (s *sSubControllerSmsManagement) TaskCreate(ctx context.Context, req *sms.SubTaskCreateReq) (res *sms.SubTaskCreateRes, err error) {

	c := 0
	var group entity.SubGroup
	if err = dao.SubGroup.Ctx(ctx).Where("id = ?", req.GroupID).ScanAndCount(&group, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询SubGroup 错误")
	}
	if c == 0 {
		return nil, errors.New("未查询到相关group id，请检查 group id是否正确")
	}

	var devices []*entity.DeviceList
	if err = dao.DeviceList.Ctx(ctx).Where("group_id = ?", group.Id).ScanAndCount(&devices, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询DeviceList错误" + err.Error())
	}
	if c == 0 {
		return nil, errors.New("当前分组无可用设备 请分配设备后再上传任务")
	}

	var project entity.ProjectList
	c = 0
	if err = dao.ProjectList.Ctx(ctx).Where("id=?", group.ProjectId).ScanAndCount(&project, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询ProjectList错误")
	}
	if c == 0 {
		return nil, errors.New("所查询的Project 不存在 请检查 Project Id 是否正确")
	}
	if c, err := dao.SmsMissionReport.Ctx(ctx).Where("task_name = ?", req.TaskName).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	} else if c != 0 {
		return nil, errors.New("重复的任务名")
	}

	filename, err := req.File.Save(consts.TaskFilePath, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件存储错误")
	}
	// Check 文件名称是否重复
	if c, err := dao.SmsMissionReport.Ctx(ctx).Where("file_name = ?", filename).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	} else if c != 0 {
		return nil, errors.New("文件名称随机重复")
	}

	content, err := ioutil.ReadFile(consts.TaskFilePath + "/" + filename)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件打开错误")
	}
	//g.Log().Infof(ctx, "content 文件内容 = %s", string(content))
	// Now let's unmarshall the data into `payload`
	var payload FileData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件解析json错误")
	}

	if len(payload.Content) == 0 || len(payload.TargetPhoneNumber) == 0 || len(payload.Content) != len(payload.TargetPhoneNumber) {
		return nil, errors.New("文件格式错误")
	}

	data := entity.SmsMissionReport{
		ProjectId:       group.ProjectId,
		TaskName:        req.TaskName,
		GroupId:         req.GroupID,
		FileName:        filename,
		TaskStatus:      1,
		SmsQuantity:     len(payload.Content),
		SurplusQuantity: len(payload.TargetPhoneNumber),
		QuantitySent:    0,
		//todo 根据子账号id查询子账号名称
		AssociatedAccount:   "todo 根据子账号id查询子账号名称",
		IntervalTime:        req.IntervalTime,
		StartTime:           req.TimingStartTime,
		ProjectName:         project.ProjectName,
		AssociatedAccountId: req.SubUserId,
	}
	g.Log().Info(ctx, data)
	var mrID int64
	if mrID, err = dao.SmsMissionReport.Ctx(ctx).Data(data).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("插入DB SmsMissionReport 错误")
	}
	res = &sms.SubTaskCreateRes{
		ID: mrID,
	}

	//将文件存储到 Redis DB
	g.Log().Infof(ctx, "filename===%s", filename)
	// 将任务列表分割成小队列
	// 查询当前组下设备的的Device 信息，然后将任务正态分布到每台Device上
	devicesLen := len(devices)
	j := 0
	for i := range payload.TargetPhoneNumber {
		message := &sms.SubPostConversationRecordData{
			TaskID:            mrID,
			Content:           payload.Content[i],
			DeviceNumber:      devices[j].DeviceNumber,
			TargetPhoneNumber: payload.TargetPhoneNumber[i],
		}
		j += 1
		if j == devicesLen {
			j = 0
		}
		// 生成任务队列
		//keyList := filename + message.DeviceNumber
		keyList := filename

		if err = utility.AddValueToCacheAndRedisQueue(ctx, keyList, message); err != nil {
			return nil, err
		}
		//if err = utility.AddValueToSafeSliceAndRedis(ctx, keyList, message); err != nil {
		//	return nil, err
		//}
	}
	return

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

// Delete Task

func (s *sSubControllerSmsManagement) TaskReportDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (res *sms.SubTaskDeleteRes, err error) {
	// todo 撤销任务时做任务状态检查
	if _, err = dao.SmsMissionReport.Ctx(ctx).Where("id = ?", req.TaskID).Delete(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("删除记录失败 SmsMissionReport")
	}
	return
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
	)

	// Step 1: Query unique conversation records based on target phone number and device number
	query := dao.SmsChartLog.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Where("project_id = ?", req.ProjectID)

	if req.SubUserID != 0 {
		query = query.Where("associated_account_id = ?", req.SubUserID)
	}

	// Fetch the unique IDs for the latest conversation records
	if err := query.
		FieldMax("id", "id").
		Group("target_phone_number").
		Group("device_number").
		Scan(&chatLogsId); err != nil {
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

// 接收Sub User输入的内容放入 Cache，作为优先任务暴露给Device执行

func (s *sSubControllerSmsManagement) PostConversationRecord(ctx context.Context, req *sms.SubPostConversationRecordReq) (res *sms.SubPostConversationRecordRes, err error) {
	// search DB 获取chart log信息
	var chartLog entity.SmsChartLog
	c := 0
	if err = dao.SmsChartLog.Ctx(ctx).Where("id = ?", req.ChartLogID).ScanAndCount(&chartLog, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库操作错误 查询 Task ChartLogID 错误")
	}
	if c == 0 {
		return nil, errors.New("未查询到 Task ChartLog")
	}

	// 验证Sub User id
	//if req.SubUserID != chartLog.AssociatedAccountId {
	//	return nil, errors.New("验证sub user id失败，发送消息不属于原任务 report 的信息")
	//}
	// 生成任务信息
	// {TargetPhoneNumber: ... DeviceNumber: ... Content: ... TaskID ...}
	message := sms.SubPostConversationRecordData{
		TaskID:            int64(chartLog.TaskId),
		Content:           req.Content,
		DeviceNumber:      chartLog.DeviceNumber,
		TargetPhoneNumber: chartLog.TargetPhoneNumber,
		Interval:          "0",
		StartAt:           gtime.Now(),
	}
	// 生成任务队列
	var rIndex int64
	if rIndex, err = g.Redis().LPush(ctx, chartLog.DeviceNumber, message); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("插入redis list 错误")
	}
	res = &sms.SubPostConversationRecordRes{}
	res.TaskItemName = chartLog.DeviceNumber
	res.UnSendTaskNum = rIndex
	return
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
