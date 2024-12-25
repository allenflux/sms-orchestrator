package subControllerDeviceManagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"sms_backend/utility"
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

// AllocateDevice2Group allocates a list of devices to a specified group.
// It validates the group ID, sub-user ID, and device IDs before updating the database.
// The function uses a transaction to ensure consistency.
//
// Parameters:
// - ctx: The context for handling the operation.
// - req: The request containing the group ID, sub-user ID, and device IDs.
//
// Returns:
// - *sms.AllocateDevice2GroupRes: The response indicating success of the allocation.
// - error: An error if the allocation fails at any point.
func (s *sSubControllerDeviceManagement) AllocateDevice2Group(ctx context.Context, req *sms.AllocateDevice2GroupReq) (*sms.AllocateDevice2GroupRes, error) {
	// Validate the device ID list
	if len(req.DeviceIdList) == 0 {
		return nil, errors.New("device ID list is empty")
	}

	// Validate and retrieve group details
	group, err := getGroupDetails(ctx, req.GroupID, req.SubUserID)
	if err != nil {
		return nil, err
	}

	// Begin database transaction
	db := g.DB()
	tx, err := db.Begin(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to start transaction: %v", err)
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			g.Log().Errorf(ctx, "Transaction rolled back due to panic: %v", p)
		}
	}()

	// Allocate devices to the group
	for _, deviceID := range req.DeviceIdList {
		// Validate device ownership and existence
		if err := validateDeviceOwnership(ctx, deviceID, req.SubUserID); err != nil {
			tx.Rollback()
			return nil, err
		}

		// Update the device with the group information
		if err := updateDeviceGroup(ctx, tx, deviceID, group.SubGroupName, group.Id); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		g.Log().Errorf(ctx, "Failed to commit transaction: %v", err)
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	// GetDeviceNumbersAndCleanCache retrieves device numbers by their IDs and cleans up associated caches.
	// Retrieve device numbers from the database using the given device IDs
	deviceNumbers, err := dao.DeviceList.
		Ctx(ctx).
		WhereIn("id", req.DeviceIdList).
		Fields("device_number").
		Array()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve device numbers: %w", err)
	}

	// Iterate through the retrieved device numbers and clean up their cache
	for _, deviceNumber := range deviceNumbers {
		if err := utility.CleanUpDeviceCache(ctx, utility.CachePrefixUnassignedDevice, deviceNumber.String()); err != nil {
			g.Log().Errorf(ctx, "Failed to clean up device cache for device_number '%s': %v", deviceNumber, err)
			return nil, fmt.Errorf("failed to clean up device cache for device_number '%s': %w", deviceNumber, err)
		}
	}

	g.Log().Infof(ctx, "Successfully cleaned up cache for %d devices", len(deviceNumbers))
	g.Log().Infof(ctx, "Successfully allocated devices to GroupID=%d for SubUserID=%d", req.GroupID, req.SubUserID)
	return &sms.AllocateDevice2GroupRes{}, nil
}

// Helper: Get group details and validate group ID and sub-user ID
func getGroupDetails(ctx context.Context, groupID, subUserID int) (*entity.SubGroup, error) {
	var group entity.SubGroup
	count := 0

	// Query group details
	if err := dao.SubGroup.Ctx(ctx).
		Where("id = ?", groupID).
		Where("sub_user_id = ?", subUserID).
		ScanAndCount(&group, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query SubGroup for GroupID=%d and SubUserID=%d: %v", groupID, subUserID, err)
		return nil, fmt.Errorf("failed to query SubGroup: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("invalid GroupID=%d or SubUserID=%d, please check the parameters", groupID, subUserID)
	}

	return &group, nil
}

// Helper: Validate device ownership and existence
func validateDeviceOwnership(ctx context.Context, deviceID, subUserID int) error {
	count, err := dao.DeviceList.Ctx(ctx).
		Where("id = ?", deviceID).
		Where("owner_account_id = ?", subUserID).
		Count()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to query DeviceList for DeviceID=%d and SubUserID=%d: %v", deviceID, subUserID, err)
		return fmt.Errorf("failed to validate device ownership: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("invalid DeviceID=%d or SubUserID=%d, please verify the input", deviceID, subUserID)
	}

	return nil
}

// Helper: Update device group details in the database
func updateDeviceGroup(ctx context.Context, tx gdb.TX, deviceID int, groupName string, groupID int) error {
	if _, err := tx.Model("device_list").
		Ctx(ctx).
		Data(g.Map{"group_name": groupName, "group_id": groupID}).
		Where("id = ?", deviceID).
		Update(); err != nil {
		g.Log().Errorf(ctx, "Failed to update DeviceList for DeviceID=%d with GroupName='%s' and GroupID=%d: %v", deviceID, groupName, groupID, err)
		return fmt.Errorf("failed to update device group details: %w", err)
	}

	g.Log().Infof(ctx, "Successfully updated DeviceID=%d with GroupName='%s' and GroupID=%d", deviceID, groupName, groupID)
	return nil
}
