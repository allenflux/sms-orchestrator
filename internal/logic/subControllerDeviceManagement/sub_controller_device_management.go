package subControllerDeviceManagement

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
	"sms_backend/internal/logic/mainControllerDeviceManagement"
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

// GetDeviceList retrieves a paginated list of devices for a specific sub-user.
func (s *sSubControllerDeviceManagement) GetDeviceList(ctx context.Context, req *sms.SubDeviceListReq) (*sms.SubDeviceListRes, error) {
	// Initialize query builder with pagination and sorting.
	query := dao.DeviceList.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Order("id DESC").
		Where("owner_account_id = ?", req.SubUserID)

	// Dynamically apply filters if provided.
	query = applyDeviceFilters(query, req)

	// Fetch data and count total records.
	var devices []*entity.DeviceList
	var totalCount int
	if err := query.ScanAndCount(&devices, &totalCount, false); err != nil {
		g.Log().Error(ctx, "Error querying DeviceList:", err)
		return nil, errors.New("failed to query DeviceList")
	}

	// Prepare response data.
	response := &sms.SubDeviceListRes{
		ListRes: commonApi.ListRes{
			Total: totalCount,
		},
		Data: transformDeviceListData(devices),
	}

	return response, nil
}

// applyDeviceFilters dynamically adds filters to the query based on the request.
func applyDeviceFilters(query *gdb.Model, req *sms.SubDeviceListReq) *gdb.Model {
	if req.TaskName != "" {
		query = query.Where("task_name LIKE ?", "%"+req.TaskName+"%")
	}
	if req.SentStatus != 0 {
		query = query.Where("sent_status = ?", req.SentStatus)
	}
	if len(req.Number) > 0 {
		query = query.Where("number LIKE ?", "%"+req.Number+"%")
	}
	if len(req.DeviceNumber) > 0 {
		query = query.Where("device_number LIKE ?", "%"+req.DeviceNumber+"%")
	}
	if len(req.SentDateRange) > 0 {
		query = query.Where("start_at BETWEEN ? AND ?", req.SentDateRange[0], req.SentDateRange[1])
	}
	if len(req.CreateDateRange) > 0 {
		query = query.Where("create_at BETWEEN ? AND ?", req.CreateDateRange[0], req.CreateDateRange[1])
	}
	return query
}

// transformDeviceListData converts raw device data into response format.
func transformDeviceListData(devices []*entity.DeviceList) []sms.SubDeviceListResData {
	data := make([]sms.SubDeviceListResData, len(devices))
	currentTime := gtime.Now()

	for i, device := range devices {
		data[i] = sms.SubDeviceListResData{
			DeviceListResData: sms.DeviceListResData{
				ID:            device.Id,
				DeviceNumber:  device.DeviceNumber,
				DeviceStatus:  device.DeviceStatus,
				SentStatus:    device.SentStatus,
				ProjectID:     device.AssignedItemsId,
				Number:        mainControllerDeviceManagement.ConvertNoPhoneNumber(device.Number),
				ActiveDays:    int(currentTime.Sub(device.ActiveTime).Hours() / 24),
				OwnerAccount:  device.OwnerAccount,
				AssignedItems: device.AssignedItems,
				QuantitySent:  strconv.Itoa(device.QuantitySent),
				ActiveTime:    device.ActiveTime.String(),
				GroupName:     device.GroupName,
				GroupID:       device.GroupId,
				DeviceID:      mainControllerDeviceManagement.PrefixDeviceID + strconv.Itoa(device.Id),
			},
		}
	}

	return data
}

// GroupCreate handles the creation of a new group for a specific sub-user.
func (s *sSubControllerDeviceManagement) GroupCreate(ctx context.Context, req *sms.SubCreateGroupReq) (*sms.SubCreateGroupRes, error) {
	// Check if the group name already exists for the sub-user
	if exists, err := isGroupNameExists(ctx, req.SubUserID, req.GroupName); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("failed to check group name in SubGroup database")
	} else if exists {
		return nil, errors.New("group name already exists, please choose another name")
	}

	// Validate the project details
	//project, err := validateProjectForGroup(ctx, req.ProjectID, req.SubUserID)
	_, err := validateProjectForGroup(ctx, req.ProjectID, req.SubUserID)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// Create the group
	groupID, err := createGroupInDatabase(ctx, req.SubUserID, req.GroupName, req.ProjectID)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("failed to create group in SubGroup database")
	}

	// Return the response
	return &sms.SubCreateGroupRes{
		ID: groupID,
	}, nil
}

// isGroupNameExists checks if a group name already exists for the specified sub-user.
func isGroupNameExists(ctx context.Context, subUserID int64, groupName string) (bool, error) {
	count, err := dao.SubGroup.Ctx(ctx).
		Where("sub_user_id = ?", subUserID).
		Where("sub_group_name = ?", groupName).
		Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// validateProjectForGroup validates if the project exists and is assigned to the correct sub-user.
func validateProjectForGroup(ctx context.Context, projectID, subUserID int64) (*entity.ProjectList, error) {
	var project entity.ProjectList
	var count int
	err := dao.ProjectList.Ctx(ctx).
		Where("id = ?", projectID).
		ScanAndCount(&project, &count, false)
	if err != nil {
		return nil, errors.New("failed to scan ProjectList database")
	}
	if count == 0 {
		return nil, errors.New("invalid Project ID: project does not exist")
	}
	if project.AssociatedAccountId == 0 {
		return nil, errors.New("invalid operation: the project has not been assigned")
	}
	if int64(project.AssociatedAccountId) != subUserID {
		return nil, errors.New("invalid operation: the project is not assigned to the current sub-user")
	}
	return &project, nil
}

// createGroupInDatabase inserts a new group into the SubGroup database.
func createGroupInDatabase(ctx context.Context, subUserID int64, groupName string, projectID int64) (int64, error) {
	groupID, err := dao.SubGroup.Ctx(ctx).
		Data(g.Map{"sub_user_id": subUserID, "sub_group_name": groupName, "project_id": projectID}).
		InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return groupID, nil
}

func (s *sSubControllerDeviceManagement) UpdateGroup(ctx context.Context, req *sms.SubUpdateGroupReq) (res *sms.SubUpdateGroupRes, err error) {
	// Check if the group name already exists for the user
	count, err := dao.SubGroup.Ctx(ctx).
		Where("sub_user_id = ?", req.SubUserID).
		Where("sub_group_name = ?", req.GroupName).
		Count()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to query SubGroup for user_id='%d' and group_name='%s': %v", req.SubUserID, req.GroupName, err)
		return nil, fmt.Errorf("failed to query SubGroup: %w", err)
	}
	if count > 0 {
		return nil, errors.New("the group name already exists, please choose a different name")
	}

	// Update the group name
	_, err = dao.SubGroup.Ctx(ctx).
		Data(g.Map{"sub_group_name": req.GroupName}).
		Where("id = ?", req.GroupID).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to update SubGroup for group_id='%d' with group_name='%s': %v", req.GroupID, req.GroupName, err)
		return nil, fmt.Errorf("failed to update group: %w", err)
	}

	// Return success response
	g.Log().Infof(ctx, "Successfully updated SubGroup for group_id='%d' with new group_name='%s'", req.GroupID, req.GroupName)
	return &sms.SubUpdateGroupRes{}, nil
}

// DeleteGroup deletes a group if there are no devices associated with it.
func (s *sSubControllerDeviceManagement) DeleteGroup(ctx context.Context, req *sms.SubDeleteGroupReq) (res *sms.SubDeleteGroupRes, err error) {
	// Check if any devices are associated with the group
	deviceCount, err := dao.DeviceList.Ctx(ctx).Where("group_id = ?", req.GroupID).Count()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to count devices for GroupID=%d: %v", req.GroupID, err)
		return nil, fmt.Errorf("failed to count devices for the group: %w", err)
	}

	// If devices exist, prevent deletion
	if deviceCount > 0 {
		g.Log().Warningf(ctx, "GroupID=%d cannot be deleted as it has associated devices", req.GroupID)
		return nil, errors.New("group cannot be deleted as it contains associated devices")
	}

	// Proceed to delete the group
	if _, err = dao.SubGroup.Ctx(ctx).Where("id = ?", req.GroupID).Delete(); err != nil {
		g.Log().Errorf(ctx, "Failed to delete GroupID=%d: %v", req.GroupID, err)
		return nil, fmt.Errorf("failed to delete group: %w", err)
	}

	g.Log().Infof(ctx, "Successfully deleted GroupID=%d", req.GroupID)
	return &sms.SubDeleteGroupRes{}, nil
}

// ListUserGroups retrieves the list of groups for a specific sub-user.
func (s *sSubControllerDeviceManagement) ListUserGroups(ctx context.Context, req *sms.SubGroupListReq) (*sms.SubGroupListRes, error) {
	// Declare a slice to store the group data.
	var groups []*entity.SubGroup
	var totalGroups int
	// Query the database for groups associated with the SubUserID.
	err := dao.SubGroup.Ctx(ctx).
		Where("sub_user_id = ?", req.SubUserID).
		ScanAndCount(&groups, &totalGroups, false)
	if err != nil {
		g.Log().Error(ctx, "Failed to query groups:", err)
		return nil, errors.New("failed to query user groups")
	}

	// Check if no groups were found.
	//if totalGroups == 0 {
	//	return nil, errors.New("no groups found for the current user")
	//}

	// Prepare the response structure.
	response := &sms.SubGroupListRes{
		Data: make([]sms.SubGroupListResData, len(groups)),
	}

	// Populate the response data.
	for i, group := range groups {
		response.Data[i] = sms.SubGroupListResData{
			GroupName: group.SubGroupName,
			GroupID:   group.Id,
			ProjectID: group.ProjectId,
		}
	}

	return response, nil
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
		if err := validateDeviceOwnership(ctx, deviceID, req.SubUserID, group.ProjectId); err != nil {
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
func validateDeviceOwnership(ctx context.Context, deviceID, subUserID, projectID int) error {
	count, err := dao.DeviceList.Ctx(ctx).
		Where("id = ?", deviceID).
		Where("owner_account_id = ?", subUserID).
		Where("assigned_items_id = ?", projectID).
		Count()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to query DeviceList for DeviceID=%d and SubUserID=%d: %v", deviceID, subUserID, err)
		return fmt.Errorf("failed to validate device ownership: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("invalid ProjectID=%d DeviceID=%d or SubUserID=%d, please verify the input", projectID, deviceID, subUserID)
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

// GetProjectList retrieves the list of projects for the front-end, scoped by the SubUserID.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing the SubUserID.
//
// Returns:
// - *sms.SubProjectListForFrontRes: The response containing the list of projects.
// - error: An error if the operation fails.
func (s *sSubControllerDeviceManagement) GetProjectList(ctx context.Context, req *sms.SubProjectListForFrontReq) (*sms.SubProjectListForFrontRes, error) {
	// Step 1: Prepare the main controller request
	mainReq := &sms.ProjectListForFrontReq{
		SubUserID: req.SubUserID,
	}

	// Step 2: Delegate to the main controller's service
	mainRes, err := service.MainControllerProjectManagement().GetProjectList(ctx, mainReq)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to retrieve project list for SubUserID=%d: %v", req.SubUserID, err)
		return nil, fmt.Errorf("failed to retrieve project list: %w", err)
	}

	// Step 3: Prepare the response
	res := &sms.SubProjectListForFrontRes{
		ProjectListForFrontRes: mainRes,
	}

	g.Log().Infof(ctx, "Successfully retrieved project list for SubUserID=%d", req.SubUserID)
	return res, nil
}
