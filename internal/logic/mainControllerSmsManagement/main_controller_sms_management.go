package mainControllerSmsManagement

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
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

func (s *sMainControllerSmsManagement) GetTaskRecordList(ctx context.Context, req *sms.TaskRecordReq) (res *sms.TaskRecordRes, err error) {
	sand := dao.SmsMissionRecord.Ctx(ctx).Page(req.PageNum, req.PageSize)
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
	res = &sms.TaskRecordRes{}
	res.Total = totalCount
	res.Data = make([]sms.TaskRecordResData, len(data))
	for i := range data {
		res.Data[i] = sms.TaskRecordResData{
			ID:                data[i].Id,
			TaskName:          data[i].TaskName,
			SubTaskID:         data[i].SubTaskId,
			TargetPhoneNumber: data[i].TargetPhoneNumber,
			DeviceNumber:      data[i].DeviceNumber,
			SmsTopic:          data[i].SmsTopic,
			SmsContent:        data[i].SmsContent,
			SmsStatus:         data[i].SmsStatus,
			Reason:            data[i].Reason,
			AssociatedAccount: data[i].AssociatedAccount,
			ProjectName:       data[i].ProjectName,
			StartTime:         data[i].StartTime.String(),
			CreateTime:        data[i].CreatedAt.String(),
		}
	}
	return
}

func (s *sMainControllerSmsManagement) GetConversationRecordList(ctx context.Context, req *sms.ConversationListReq) (res *sms.ConversationListRes, err error) {
	var chatLogsId []*entity.SmsChartLog
	var chatLogs []*entity.SmsChartLog
	if err = dao.SmsChartLog.Ctx(ctx).Where("project_id=?", req.ProjectID).FieldMax("id", "id").Page(req.PageNum, req.PageSize).Group("target_phone_number").Group("device_number").Scan(&chatLogsId); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("Main查询SmsChartLog IDs错误")
	}
	idList := make([]int, len(chatLogsId))
	for i := range chatLogsId {
		idList[i] = chatLogsId[i].Id
	}

	if err = dao.SmsChartLog.Ctx(ctx).WhereIn("id", idList).Scan(&chatLogs); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("Main查询SmsChartLog错误")
	}
	data := make([]sms.ConversationRecordListResData, len(chatLogs))
	for i := range chatLogs {
		data[i].ChatLogID = chatLogs[i].Id
		data[i].RecordTime = chatLogs[i].CreatedAt
		data[i].Content = chatLogs[i].SmsContent
		data[i].TargetPhoneNumber = chatLogs[i].TargetPhoneNumber
		data[i].SentOrReceive = chatLogs[i].SentOrReceive
	}
	res = &sms.ConversationListRes{
		Data: data,
	}
	return
}
