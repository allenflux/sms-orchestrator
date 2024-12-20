package subControllerDeviceManagement

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"strconv"
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
	dbTemper := dao.DeviceList.Ctx(ctx).Page(req.PageNum, req.PageSize).Order("id desc").Where("owner_account_id=", req.SubUserID)

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
		return nil, errors.New("查询 DeviceList 错误")
	}
	res = &sms.SubDeviceListRes{}
	res.Total = totalCount
	data := make([]sms.SubDeviceListResData, len(raw))
	currentTime := gtime.Now()
	for i := range raw {
		data[i] = sms.SubDeviceListResData{
			ID:            raw[i].Id,
			DeviceNumber:  raw[i].DeviceNumber,
			DeviceStatus:  raw[i].DeviceStatus,
			SentStatus:    raw[i].SentStatus,
			ProjectID:     raw[i].AssignedItemsId,
			Number:        raw[i].Number,
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

func (s *sSubControllerDeviceManagement) GroupCreate(ctx context.Context, req *sms.SubCreateGroupReq) (res *sms.SubCreateGroupRes, err error) {
	if count, err := dao.SubGroup.Ctx(ctx).Where("sub_user_id = ?", req.SubUserID).Where("sub_group_name = ?", req.GroupName).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("Count SubGroup DB 错误 ")
	} else if count > 0 {
		return nil, errors.New("此名称以存在 请更换分组名称")
	}

	var project entity.ProjectList
	c := 0
	if err := dao.ProjectList.Ctx(ctx).Where("id = ?", req.ProjectID).ScanAndCount(&project, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("Scan ProjectList DB 错误 ")
	} else if c == 0 {
		return nil, errors.New("Project ID 错误 所查询的Project 不存在 ")
	}
	if project.AssociatedAccountId == 0 {
		return nil, errors.New("非法创建 当前项目还未分配")
	}
	if project.AssociatedAccountId != req.SubUserID {
		return nil, errors.New("非法创建 这个项目所分配的子用户不是当前子用户")
	}

	var rawId int64
	if rawId, err = dao.SubGroup.Ctx(ctx).Data(g.Map{"sub_user_id": req.SubUserID, "sub_group_name": req.GroupName, "project_id": req.ProjectID}).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("SubGroup 分组创建失败")
	}
	res = &sms.SubCreateGroupRes{
		ID: rawId,
	}
	return
}

func (s *sSubControllerDeviceManagement) GroupUpdate(ctx context.Context, req *sms.SubUpdateGroupReq) (res *sms.SubUpdateGroupRes, err error) {
	if count, err := dao.SubGroup.Ctx(ctx).Where("sub_user_id = ?", req.SubUserID).Where("sub_group_name = ?", req.GroupName).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("SubGroup 查询错误")
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
	var data []*entity.SubGroup
	c := 0
	if err = dao.SubGroup.Ctx(ctx).Where("sub_user_id=?", req.SubUserID).ScanAndCount(&data, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("分组查询失败")
	}
	if c == 0 {
		return nil, errors.New("当前用户无分组")
	}
	res = &sms.SubGroupListRes{
		Data: make([]sms.SubGroupListResData, len(data)),
	}
	for i := range data {
		res.Data[i] = sms.SubGroupListResData{
			GroupName: data[i].SubGroupName,
			GroupID:   data[i].Id,
			ProjectID: data[i].ProjectId,
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
	c := 0
	if err = dao.SubGroup.Ctx(ctx).Where("id = ?", req.GroupID).Where("sub_user_id = ?", req.SubUserID).ScanAndCount(&group, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库查询SubGroup错误")
	}
	if c == 0 {
		return nil, errors.New("未查询到相关Group 请检查参数 Group id和Sub user id")
	}
	for i := range req.DeviceIdList {

		if count, err := dao.DeviceList.Ctx(ctx).Where("id=?", req.DeviceIdList[i]).Where("owner_account_id = ?", req.SubUserID).Count(); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("查询 DeviceList 错误 ")
		} else if count == 0 {
			return nil, errors.New("错误的Device ID或者Sub User ID， 请改正")
		}
		if _, err = dao.DeviceList.Ctx(ctx).Data(g.Map{"group_name": group.SubGroupName, "group_id": group.Id}).Where("id = ?", req.DeviceIdList[i]).Update(); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("更新 DeviceList DB group_name 和 group_id 错误")
		}
	}
	return
}
