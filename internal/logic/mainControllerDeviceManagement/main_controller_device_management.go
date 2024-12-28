package mainControllerDeviceManagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	commonApi "sms_backend/api/v1/common"
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
const PrefixDeviceID = "DeviceID"

// ConvertNoPhoneNumber converts "NoPhoneNumber" to an empty string; otherwise, returns the input string.
func ConvertNoPhoneNumber(str string) string {
	if str == NoPhoneNumber {
		return ""
	}
	return str
}

// GetDeviceList retrieves a paginated list of devices based on the given filters.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing the filters for querying the device list.
//
// Returns:
// - *sms.DeviceListRes: The response containing the paginated device list and metadata.
// - error: An error if the operation fails.
func (s *sMainControllerDeviceManagement) GetDeviceList(ctx context.Context, req *sms.DeviceListReq) (*sms.DeviceListRes, error) {
	// Initialize database query with pagination and ordering
	dbQuery := dao.DeviceList.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Order("id DESC")

	// Apply dynamic filters to the query
	dbQuery = applyDeviceListFilters(dbQuery, req)

	// Execute the query and count the total records
	var rawDevices []*entity.DeviceList
	var totalCount int
	if err := dbQuery.ScanAndCount(&rawDevices, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query DeviceList: %v", err)
		return nil, fmt.Errorf("failed to query DeviceList: %w", err)
	}

	// Transform raw database results into response format
	res := &sms.DeviceListRes{
		ListRes: commonApi.ListRes{
			Total: totalCount,
		},
		Data: transformDeviceListToResponse(rawDevices),
	}

	g.Log().Infof(ctx, "Successfully retrieved %d devices for request: %+v", totalCount, req)
	return res, nil
}

// applyDeviceListFilters dynamically applies filters to the database query based on the request parameters.
//
// Parameters:
// - dbQuery: The initial database query.
// - req: The request containing filter parameters.
//
// Returns:
// - *gdb.Model: The updated database query with applied filters.
func applyDeviceListFilters(dbQuery *gdb.Model, req *sms.DeviceListReq) *gdb.Model {
	if req.DeviceID != "" {
		dbQuery = dbQuery.Where("id ?", req.DeviceID)
	}
	if req.TaskName != "" {
		dbQuery = dbQuery.Where("task_name LIKE ?", "%"+req.TaskName+"%")
	}
	if req.SentStatus != 0 {
		dbQuery = dbQuery.Where("sent_status = ?", req.SentStatus)
	}
	if req.Number != "" {
		dbQuery = dbQuery.Where("number LIKE ?", "%"+req.Number+"%")
	}
	if req.DeviceNumber != "" {
		dbQuery = dbQuery.Where("device_number LIKE ?", "%"+req.DeviceNumber+"%")
	}
	if len(req.SentDateRange) == 2 {
		dbQuery = dbQuery.Where("start_at BETWEEN ? AND ?", req.SentDateRange[0], req.SentDateRange[1])
	}
	if len(req.CreateDateRange) == 2 {
		dbQuery = dbQuery.Where("create_at BETWEEN ? AND ?", req.CreateDateRange[0], req.CreateDateRange[1])
	}
	return dbQuery
}

// transformDeviceListToResponse converts raw database entities into response data for the API.
//
// Parameters:
// - rawDevices: The raw database entities retrieved from the query.
//
// Returns:
// - []sms.DeviceListResData: The transformed response data.
func transformDeviceListToResponse(rawDevices []*entity.DeviceList) []sms.DeviceListResData {
	responseData := make([]sms.DeviceListResData, len(rawDevices))
	currentTime := gtime.Now()

	for i, device := range rawDevices {
		responseData[i] = sms.DeviceListResData{
			ID:            device.Id,
			DeviceNumber:  device.DeviceNumber,
			DeviceID:      PrefixDeviceID + strconv.Itoa(device.Id), // Default value, can be populated later if needed
			DeviceStatus:  device.DeviceStatus,
			SentStatus:    device.SentStatus,
			ProjectID:     device.AssignedItemsId,
			Number:        ConvertNoPhoneNumber(device.Number),
			ActiveDays:    int(currentTime.Sub(device.ActiveTime).Hours() / 24),
			OwnerAccount:  device.OwnerAccount,
			AssignedItems: device.AssignedItems,
			QuantitySent:  strconv.Itoa(device.QuantitySent),
			ActiveTime:    device.ActiveTime.String(),
			GroupID:       device.GroupId,
			GroupName:     device.GroupName,
		}
	}

	return responseData
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
