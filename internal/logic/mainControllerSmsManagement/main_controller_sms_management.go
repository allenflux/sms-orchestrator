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
	return
}
