package mainControllerSmsManagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
	"sms_backend/internal/model"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
)

func New() *sMainControllerSmsManagement {
	return &sMainControllerSmsManagement{}
}

func init() {
	service.RegisterMainControllerSmsManagement(New())
}

type sMainControllerSmsManagement struct{}

func (s *sMainControllerSmsManagement) GetTaskList(ctx context.Context, req *sms.TaskListReq) (res *sms.TaskListRes, err error) {
	sand := dao.SmsMissionReport.Ctx(ctx).Page(req.PageNum, req.PageSize)
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
	res = &sms.TaskListRes{}
	res.Total = totalCount
	res.Data = make([]sms.TaskListResData, len(data))
	for i := range data {
		res.Data[i] = sms.TaskListResData{
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

// GetTaskRecordList retrieves a paginated list of task records based on the provided filters.
func (s *sMainControllerSmsManagement) GetTaskRecordList(ctx context.Context, req *sms.TaskRecordReq) (*sms.TaskRecordRes, error) {
	// Initialize the query builder
	query := dao.SmsMissionRecord.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		OrderDesc("id")

	// Apply dynamic filters
	query = applyTaskRecordFilters(query, req)

	// Fetch data and count
	var records []entity.SmsMissionRecord
	var totalCount int
	if err := query.ScanAndCount(&records, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query SmsMissionRecord: %v", err)
		return nil, errors.New("failed to query SmsMissionRecord")
	}

	// Transform data for the response
	res := &sms.TaskRecordRes{
		ListRes: commonApi.ListRes{
			Total: totalCount,
		},
		Data: transformTaskRecordData(records),
	}

	return res, nil
}

// applyTaskRecordFilters dynamically applies filters to the query.
func applyTaskRecordFilters(query *gdb.Model, req *sms.TaskRecordReq) *gdb.Model {
	if req.ProjectID != 0 {
		query = query.Where("project_id = ?", req.ProjectID)
	}
	if req.SmsStatus != 0 {
		query = query.Where("sms_status = ?", req.SmsStatus)
	}
	if len(req.TaskName) != 0 {
		query = query.Where("task_name LIKE ?", "%"+req.TaskName+"%")
	}
	if len(req.TargetPhoneNumber) != 0 {
		query = query.Where("target_phone_number LIKE ?", "%"+req.TargetPhoneNumber+"%")
	}
	if len(req.DeviceNumber) != 0 {
		query = query.Where("device_number LIKE ?", "%"+req.DeviceNumber+"%")
	}
	if len(req.SentDateRange) > 0 {
		query = query.Where("start_time >= ? AND start_time <= ?", req.SentDateRange[0], req.SentDateRange[1])
	}
	if len(req.CreateDateRange) > 0 {
		query = query.Where("created_at >= ? AND created_at <= ?", req.CreateDateRange[0], req.CreateDateRange[1])
	}
	return query
}

// transformTaskRecordData transforms the raw entity data into response data.
func transformTaskRecordData(records []entity.SmsMissionRecord) []sms.TaskRecordResData {
	data := make([]sms.TaskRecordResData, len(records))
	for i, record := range records {
		data[i] = sms.TaskRecordResData{
			ID:                record.Id,
			TaskName:          record.TaskName,
			SubTaskID:         record.SubTaskId,
			TargetPhoneNumber: record.TargetPhoneNumber,
			DeviceNumber:      record.DeviceNumber,
			SmsTopic:          record.SmsTopic,
			SmsContent:        record.SmsContent,
			SmsStatus:         record.SmsStatus,
			Reason:            record.Reason,
			AssociatedAccount: record.AssociatedAccount,
			ProjectName:       record.ProjectName,
			StartTime:         record.StartTime.String(),
			CreateTime:        record.CreatedAt.String(),
		}
	}
	return data
}

// GetSubGetConversationRecord retrieves a specific conversation record by delegating the request to the sub-controller service.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing the ChatLogID.
//
// Returns:
// - *sms.ConversationRecordRes: The response containing the conversation record.
// - error: An error if the operation fails.
func (s *sMainControllerSmsManagement) GetSubGetConversationRecord(ctx context.Context, req *sms.ConversationRecordReq) (*sms.ConversationRecordRes, error) {
	// Prepare the sub-controller request
	subReq := &sms.SubGetConversationRecordReq{
		ChatLogID: req.ChatLogID,
	}
	// Delegate the request to the sub-controller service
	subRes, err := service.SubControllerSmsManagement().GetSubGetConversationRecord(ctx, subReq)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to fetch conversation record for ChatLogID=%d: %v", req.ChatLogID, err)
		return nil, fmt.Errorf("failed to fetch conversation record: %w", err)
	}

	// Prepare the response
	res := &sms.ConversationRecordRes{
		SubGetConversationRecordRes: subRes,
	}

	g.Log().Infof(ctx, "Successfully fetched conversation record for ChatLogID=%d", req.ChatLogID)
	return res, nil
}

// GetConversationRecordList retrieves the list of conversation records by delegating the request to the sub-controller service.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing pagination details and the ProjectID.
//
// Returns:
// - *sms.ConversationListRes: The response containing the conversation records.
// - error: An error if the operation fails.
func (s *sMainControllerSmsManagement) GetConversationRecordList(ctx context.Context, req *sms.ConversationListReq) (*sms.ConversationListRes, error) {
	// Prepare the sub-controller request
	subReq := &sms.SubGetConversationRecordListReq{
		PageReq: model.PageReq{
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
		},
		SubUserID: 0, // Default value
		ProjectID: req.ProjectID,
	}

	// Delegate the request to the sub-controller service
	subRes, err := service.SubControllerSmsManagement().SubGetConversationRecordList(ctx, subReq)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to fetch conversation records for ProjectID=%d: %v", req.ProjectID, err)
		return nil, fmt.Errorf("failed to fetch conversation records: %w", err)
	}

	// Prepare the response
	res := &sms.ConversationListRes{
		SubGetConversationRecordListRes: subRes,
	}

	g.Log().Infof(ctx, "Successfully fetched conversation records for ProjectID=%d, Total=%d", req.ProjectID, subRes.Total)
	return res, nil
}

// PostConversationRecord handles posting conversation records and forwards the request to the sub-controller.
func (s *sMainControllerSmsManagement) PostConversationRecord(ctx context.Context, req *sms.PostConversationRecordReq) (*sms.PostConversationRecordRes, error) {
	// Forward the request to the sub-controller
	subRes, err := forwardConversationRecordRequest(ctx, req.SubPostConversationRecordReq)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to post conversation record: %v", err)
		return nil, fmt.Errorf("failed to post conversation record: %w", err)
	}

	// Prepare and return the response
	return &sms.PostConversationRecordRes{
		SubPostConversationRecordRes: subRes,
	}, nil
}

// forwardConversationRecordRequest forwards the conversation record request to the sub-controller.
func forwardConversationRecordRequest(ctx context.Context, subReq *sms.SubPostConversationRecordReq) (*sms.SubPostConversationRecordRes, error) {
	subRes, err := service.SubControllerSmsManagement().PostConversationRecord(ctx, subReq)
	if err != nil {
		return nil, fmt.Errorf("error from sub-controller: %w", err)
	}

	g.Log().Infof(ctx, "Successfully forwarded conversation record request: %+v", subReq)
	return subRes, nil
}

// GetTaskDevices retrieves the list of devices assigned to a specific task (group ID) with pagination support.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing the task ID and pagination details.
//
// Returns:
// - *sms.TaskDevicesRes: The response containing the list of devices and pagination metadata.
// - error: An error if the operation fails.
func (s *sMainControllerSmsManagement) GetTaskDevices(ctx context.Context, req *sms.TaskDevicesReq) (*sms.TaskDevicesRes, error) {
	var (
		devices []*entity.DeviceList
		group   entity.SmsMissionReport
		count   int
	)

	// Query Group ID by Task ID
	record, err := dao.SmsMissionReport.Ctx(ctx).
		Where("id = ?", req.TaskID).
		Fields("group_id").
		One()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to fetch group ID for TaskID=%d: %v", req.TaskID, err)
		return nil, fmt.Errorf("failed to fetch group ID: %w", err)
	}

	// Map the record to the SmsMissionReport struct
	if err := record.Struct(&group); err != nil {
		g.Log().Errorf(ctx, "Failed to map group data for TaskID=%d: %v", req.TaskID, err)
		return nil, fmt.Errorf("failed to map group data: %w", err)
	}

	// Query the device list with pagination
	if err := dao.DeviceList.Ctx(ctx).
		Where("group_id = ?", group.GroupId).
		Page(req.PageNum, req.PageSize).
		ScanAndCount(&devices, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query device list for GroupID=%d (TaskID=%d): %v", group.GroupId, req.TaskID, err)
		return nil, fmt.Errorf("failed to query device list: %w", err)
	}

	// Transform the device list into response data
	data := make([]sms.TaskDevicesResData, len(devices))
	for i, device := range devices {
		data[i] = sms.TaskDevicesResData{
			ID:        device.Id, // Using the device's ID instead of the task ID
			TaskID:    req.TaskID,
			DeviceNum: device.DeviceNumber,
		}
	}

	// Prepare the response with pagination metadata
	return &sms.TaskDevicesRes{
		ListRes: commonApi.ListRes{
			Total:       count,
			CurrentPage: req.PageNum,
		},
		Data: data,
	}, nil
}

// GetPendingTasks retrieves pending tasks by delegating the request to the sub-controller service.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing the TaskID.
//
// Returns:
// - *sms.PendingTaskRes: The response containing the pending tasks and metadata.
// - error: An error if the operation fails.
func (s *sMainControllerSmsManagement) GetPendingTasks(ctx context.Context, req *sms.PendingTaskReq) (*sms.PendingTaskRes, error) {
	// Prepare the sub-controller request
	subReq := &sms.SubPendingTaskReq{
		TaskID: req.TaskID,
	}

	// Delegate the request to the sub-controller service
	subRes, err := service.SubControllerSmsManagement().GetPendingTasks(ctx, subReq)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to fetch pending tasks for TaskID=%d: %v", req.TaskID, err)
		return nil, fmt.Errorf("failed to fetch pending tasks: %w", err)
	}

	// Prepare the response
	res := &sms.PendingTaskRes{
		ListRes: commonApi.ListRes{
			Total: subRes.Total,
		},
		Data: subRes.Data,
	}

	g.Log().Infof(ctx, "Successfully fetched pending tasks for TaskID=%d, Total=%d", req.TaskID, subRes.Total)
	return res, nil
}
