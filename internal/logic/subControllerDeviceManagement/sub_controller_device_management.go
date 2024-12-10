package subControllerDeviceManagement

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

func New() *sSubControllerDeviceManagement {
	return &sSubControllerDeviceManagement{}
}

func init() {
	service.RegisterSubControllerDeviceManagement(New())
}

type sSubControllerDeviceManagement struct{}

func (s *sSubControllerDeviceManagement) GetDeviceList(ctx context.Context, req *sms.SubDeviceListReq) (res *sms.SubDeviceListRes, err error) {
	raw := make([]*entity.DeviceList, 0)
	dbTemper := dao.DeviceList.Ctx(ctx).Page(req.PageNum, req.PageSize).Order("id desc").Where("owner_account_id = ？", req.SubUserID)

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

func (s *sSubControllerDeviceManagement) GroupCreate(ctx context.Context, req *sms.SubCreateGroupReq) (res *sms.SubCreateGroupRes, err error) {
	if count, err := dao.SubGroup.Ctx(ctx).Where("sub_user_id = ?", req.SubUserID).Where("sub_group_name = ?", req.GroupName).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	} else if count > 0 {
		return nil, errors.New("此名称以存在 请更换分组名称")
	}
	if _, err = dao.SubGroup.Ctx(ctx).Data(g.Map{"sub_user": req.SubUserID, "sub_group_name": req.GroupName}).Insert(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("分组创建失败")
	}
	return
}

func (s *sSubControllerDeviceManagement) GroupUpdate(ctx context.Context, req *sms.SubUpdateGroupReq) (res *sms.SubUpdateGroupRes, err error) {
	if count, err := dao.SubGroup.Ctx(ctx).Where("sub_user_id = ?", req.SubUserID).Where("sub_group_name = ?", req.GroupName).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	} else if count > 0 {
		return nil, errors.New("此名称以存在 请更换分组名称再更新")
	}

	if _, err = dao.SubGroup.Ctx(ctx).Data(g.Map{"sub_group_name": req.GroupName}).Where("id=?", req.GroupID).Update(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("分组更新失败")
	}
	return
}

func (s *sSubControllerDeviceManagement) GroupDelete(ctx context.Context, req *sms.SubDeleteGroupReq) (res *sms.SubDeleteGroupRes, err error) {
	// todo 设备分组被删除后，被删除分组的设备自动重置为未分组
	// todo 设备存在待发送、发送中的任务，则发送状态为占用
	if _, err = dao.SubGroup.Ctx(ctx).Where("id=?", req.GroupID).Delete(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("分组删除失败")
	}
	return
}

// Get Group List

func (s *sSubControllerDeviceManagement) GroupList(ctx context.Context, req *sms.SubGroupListReq) (res *sms.SubGroupListRes, err error) {
	data := make([]*entity.SubGroup, 0)
	if err = dao.SubGroup.Ctx(ctx).Where("sub_user_id=?", req.SubUserID).Scan(data); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("分组查询失败")
	}
	res.Data = make([]sms.SubGroupListResData, len(data))
	for i := range data {
		res.Data[i] = sms.SubGroupListResData{
			GroupName: data[i].SubGroupName,
			GroupID:   data[i].Id,
		}
	}
	return
}

func (s *sSubControllerDeviceManagement) AllocateDevice2Group(ctx context.Context, req *sms.AllocateDevice2GroupReq) (res *sms.AllocateDevice2GroupRes, err error) {
	if len(req.DeviceIdList) == 0 {
		return nil, errors.New("device id list is empty")
	}
	// get Group name
	var group entity.SubGroup
	if err = dao.SubGroup.Ctx(ctx).Where("id = ?", req.GroupID).Scan(&group); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库查询SubGroup错误")
	}
	for i := range req.DeviceIdList {
		if _, err = dao.DeviceList.Ctx(ctx).Data(g.Map{"group_name": group.SubGroupName, "group_id": group.Id}).Where("id = ?", req.DeviceIdList[i]).Update(); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("更新 DeviceList DB group_name 和 group_id 错误")
		}
	}
	return
}
