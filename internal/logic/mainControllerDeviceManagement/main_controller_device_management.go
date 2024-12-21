package mainControllerDeviceManagement

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"strconv"
)

func New() *sMainControllerDeviceManagement {
	return &sMainControllerDeviceManagement{}
}

func init() {
	service.RegisterMainControllerDeviceManagement(New())
}

type sMainControllerDeviceManagement struct{}

const NoPhoneNumber = "NoPhoneNumber"

func ConverNoPhoneNumber(str string) string {
	if str == NoPhoneNumber {
		return ""
	}
	return str
}

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
		return nil, errors.New("ScanAndCount错误 Table DeviceList")
	}
	res = &sms.DeviceListRes{}
	res.Total = totalCount
	data := make([]sms.DeviceListResData, len(raw))

	currentTime := gtime.Now()
	for i := range raw {
		data[i] = sms.DeviceListResData{
			ID:            raw[i].Id,
			DeviceNumber:  raw[i].DeviceNumber,
			DeviceID:      "",
			DeviceStatus:  raw[i].DeviceStatus,
			SentStatus:    raw[i].SentStatus,
			ProjectID:     raw[i].AssignedItemsId,
			Number:        ConverNoPhoneNumber(raw[i].Number),
			ActiveDays:    int(currentTime.Sub(raw[i].ActiveTime).Hours() / 24),
			OwnerAccount:  raw[i].OwnerAccount,
			AssignedItems: raw[i].AssignedItems,
			QuantitySent:  strconv.Itoa(raw[i].QuantitySent),
			ActiveTime:    raw[i].ActiveTime.String(),
		}
	}
	res.Data = data
	return
}

func (s *sMainControllerDeviceManagement) AllocateDevice2Project(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (res *sms.AllocateDevice2ProjectRes, err error) {
	if len(req.DeviceIdList) == 0 {
		return nil, errors.New("device id list is empty")
	}
	// get project name
	var project entity.ProjectList
	c := 0
	if err = dao.ProjectList.Ctx(ctx).Where("id = ?", req.ProjectID).ScanAndCount(&project, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库查询ProjectList错误" + err.Error())
	}
	if c == 0 {
		return nil, errors.New("错误的ProjectID")
	}

	db := g.DB()
	var tx gdb.TX
	if tx, err = db.Begin(ctx); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("开启事务操作失败")
	}
	g.Log().Info(ctx, "开启事务操作")

	for i := range req.DeviceIdList {
		c := 0
		dl := &entity.DeviceList{}
		if err = dao.DeviceList.Ctx(ctx).Where("id=?", req.DeviceIdList[i]).ScanAndCount(dl, &c, false); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("查询DeviceList错误")
		} else if c == 0 {
			return nil, errors.New("不存在的Device， 请验证Device ID是否正确")
		} else if dl.GroupId != 0 {
			return nil, errors.New("当前Device {" + dl.DeviceNumber + "} 已经被分配 请重新选择")
		}

		if _, err = tx.Model("device_list").Ctx(ctx).Data(g.Map{"assigned_items": project.ProjectName, "assigned_items_id": project.Id}).Where("id = ?", req.DeviceIdList[i]).Update(); err != nil {
			g.Log().Error(ctx, err)
			if err = tx.Rollback(); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("rollback Error")
			}
			return nil, errors.New("更新 DeviceList DB assigned_items 和 assigned_items_id 错误")
		}
	}
	if _, err = tx.Model("project_list").Ctx(ctx).Data("quantity_device = ? ", project.QuantityDevice+len(req.DeviceIdList)).Where("id = ?", project.Id).Update(); err != nil {
		g.Log().Error(ctx, err)
		if err = tx.Rollback(); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("rollback Error")
		}
		return nil, errors.New("更新 project_list DB quantity_device  错误")
	}
	if err = tx.Commit(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("Commit  Error")
	}
	return
}
