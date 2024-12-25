package mainControllerDeviceManagement

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

// AllocateDevice2Project allocates a list of devices to a project.
// It validates the device IDs, checks if the project exists, and updates the database accordingly.
//
// Parameters:
// - ctx: The context for handling the operation.
// - req: The request containing the project ID and list of device IDs to be allocated.
//
// Returns:
// - *sms.AllocateDevice2ProjectRes: The response indicating the success of the operation.
// - error: An error if the operation fails at any step.
func (s *sMainControllerDeviceManagement) AllocateDevice2Project(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (*sms.AllocateDevice2ProjectRes, error) {
	// Validate the request
	if len(req.DeviceIdList) == 0 {
		return nil, errors.New("device id list is empty")
	}

	// Check if the project exists and retrieve its details
	var project entity.ProjectList
	count := 0
	if err := dao.ProjectList.Ctx(ctx).Where("id = ?", req.ProjectID).ScanAndCount(&project, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query ProjectList for ProjectID=%d: %v", req.ProjectID, err)
		return nil, fmt.Errorf("failed to query project list: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("invalid ProjectID: %d", req.ProjectID)
	}

	// Begin a database transaction
	db := g.DB()
	tx, err := db.Begin(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to start transaction: %v", err)
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			g.Log().Errorf(ctx, "Transaction rolled back due to panic: %v", p)
		}
	}()

	g.Log().Info(ctx, "Transaction started for device allocation")

	// Process each device ID in the list
	for _, deviceID := range req.DeviceIdList {
		// Check if the device exists and is not already allocated
		_, err := validateAndRetrieveDevice(ctx, tx, deviceID)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		// Prepare the update data for the device
		updateData := g.Map{
			"assigned_items":    project.ProjectName,
			"assigned_items_id": project.Id,
		}
		if project.AssociatedAccountId != 0 {
			updateData["owner_account"] = project.AssociatedAccount
			updateData["owner_account_id"] = project.AssociatedAccountId
		}

		// Update the device record
		if _, err := tx.Model("device_list").Ctx(ctx).Data(updateData).Where("id = ?", deviceID).Update(); err != nil {
			g.Log().Errorf(ctx, "Failed to update DeviceList for DeviceID=%d: %v", deviceID, err)
			_ = tx.Rollback()
			return nil, fmt.Errorf("failed to update DeviceList for DeviceID=%d: %w", deviceID, err)
		}
	}

	// Update the project's device quantity
	if _, err := tx.Model("project_list").Ctx(ctx).
		Data("quantity_device = ?", project.QuantityDevice+len(req.DeviceIdList)).
		Where("id = ?", project.Id).Update(); err != nil {
		g.Log().Errorf(ctx, "Failed to update ProjectList quantity_device for ProjectID=%d: %v", project.Id, err)
		_ = tx.Rollback()
		return nil, fmt.Errorf("failed to update ProjectList quantity_device for ProjectID=%d: %w", project.Id, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		g.Log().Errorf(ctx, "Transaction commit failed: %v", err)
		return nil, fmt.Errorf("transaction commit failed: %w", err)
	}

	g.Log().Infof(ctx, "Successfully allocated devices to ProjectID=%d", req.ProjectID)
	return &sms.AllocateDevice2ProjectRes{}, nil
}

// Helper: Validate and retrieve a device by ID
func validateAndRetrieveDevice(ctx context.Context, tx gdb.TX, deviceID int) (*entity.DeviceList, error) {
	var device entity.DeviceList
	count := 0

	// Query the device from the database
	if err := dao.DeviceList.Ctx(ctx).Where("id = ?", deviceID).ScanAndCount(&device, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query DeviceList for DeviceID=%d: %v", deviceID, err)
		return nil, fmt.Errorf("failed to query DeviceList: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("invalid DeviceID: %d", deviceID)
	}

	// Check if the device is already assigned
	if device.GroupId != 0 {
		return nil, fmt.Errorf("device '%s' is already assigned to a group", device.DeviceNumber)
	}

	return &device, nil
}
