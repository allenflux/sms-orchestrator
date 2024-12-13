package mainControllerDeviceManagement

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"sms_backend/utility"
	"strconv"
	"time"
)

func New() *sMainControllerDeviceManagement {
	return &sMainControllerDeviceManagement{}
}

func init() {
	service.RegisterMainControllerDeviceManagement(New())
}

type sMainControllerDeviceManagement struct{}

func (s *sMainControllerDeviceManagement) GetDeviceList(ctx context.Context, req *sms.DeviceListReq) (res *sms.DeviceListRes, err error) {
	raw := make([]*entity.DeviceList, 0)
	dbTemper := dao.DeviceList.Ctx(ctx).Page(req.PageNum, req.PageSize).Order("id desc")

	if req.TaskName != "" {
		dbTemper = dbTemper.Where("task_name like ?", "%"+req.TaskName+"%")
	}

	if req.SentStatus != 0 {
		dbTemper = dbTemper.Where("sent_status = ?", req.SentStatus)
	}

	if len(req.Number) != 0 {
		dbTemper = dbTemper.Where("number like ?", "%"+req.Number+"%")
	}

	if len(req.DeviceNumber) != 0 {
		dbTemper = dbTemper.Where("device_number like ?", "%"+req.DeviceNumber+"%")
	}

	if len(req.SentDateRange) > 0 {
		dbTemper = dbTemper.Where("start_at >= ? AND start_at <= ?", req.SentDateRange[0], req.SentDateRange[1])
	}

	if len(req.CreateDateRange) > 0 {
		dbTemper = dbTemper.Where("create_at >= ? AND create_at <= ?", req.CreateDateRange[0], req.CreateDateRange[1])
	}

	var totalCount int
	if err := dbTemper.ScanAndCount(&raw, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res.Total = totalCount
	data := make([]sms.DeviceListResData, len(raw))
	currentTime := time.Now().Second()
	for i := range raw {
		data[i] = sms.DeviceListResData{
			ID:            raw[i].Id,
			DeviceNumber:  raw[i].DeviceNumber,
			DeviceStatus:  raw[i].DeviceStatus,
			SentStatus:    raw[i].SentStatus,
			DeviceID:      raw[i].DeviceId,
			Number:        raw[i].Number,
			ActiveDays:    utility.Second2Day(currentTime - raw[i].ActiveTime.Second()),
			OwnerAccount:  raw[i].OwnerAccount,
			AssignedItems: raw[i].AssignedItems,
			QuantitySent:  strconv.Itoa(raw[i].QuantitySent),
			ActiveTime:    raw[i].ActiveTime.String(),
		}
	}

	return
}

func (s *sMainControllerDeviceManagement) AllocateDevice2Project(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (res *sms.AllocateDevice2ProjectRes, err error) {
	if len(req.DeviceIdList) == 0 {
		return nil, errors.New("device id list is empty")
	}
	// get project name
	var project entity.ProjectList
	if err = dao.ProjectList.Ctx(ctx).Where("id = ?", req.ProjectID).Scan(&project); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库查询ProjectList错误")
	}

	for i := range req.DeviceIdList {
		if _, err = dao.DeviceList.Ctx(ctx).Data(g.Map{"assigned_items": project.ProjectName, "assigned_items_id": project.Id}).Where("id = ?", req.DeviceIdList[i]).Update(); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("更新 DeviceList DB assigned_items 和 assigned_items_id 错误")
		}
	}
	return
}
