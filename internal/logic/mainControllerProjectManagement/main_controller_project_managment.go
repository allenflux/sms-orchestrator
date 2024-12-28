package mainControllerProjectManagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
)

func New() *sMainControllerProjectManagement {
	return &sMainControllerProjectManagement{}
}

func init() {
	service.RegisterMainControllerProjectManagement(New())
}

type sMainControllerProjectManagement struct{}

// GetProjectList retrieves the list of projects, optionally filtered by SubUserID.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing optional SubUserID for filtering.
//
// Returns:
// - *sms.ProjectListForFrontRes: The response containing the project list.
// - error: An error if the operation fails.
func (s *sMainControllerProjectManagement) GetProjectList(ctx context.Context, req *sms.ProjectListForFrontReq) (*sms.ProjectListForFrontRes, error) {
	r := ghttp.RequestFromCtx(ctx)
	currentRoute := r.Router.Uri
	g.Log().Info(ctx, "Current Route:", currentRoute)
	if currentRoute == "/api/v1/sms/project/items" {
		req.SubUserID = 0
	}

	// Step 1: Initialize the database query
	query := dao.ProjectList.Ctx(ctx)

	// Step 2: Apply filters based on SubUserID, if provided
	if req.SubUserID != 0 {
		g.Log().Infof(ctx, "Fetching projects for SubUserID=%d", req.SubUserID)
		query = query.Where("associated_account_id = ?", req.SubUserID)
	}

	// Step 3: Fetch project list and count
	var projects []*entity.ProjectList
	count := 0
	if err := query.ScanAndCount(&projects, &count, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query ProjectList: %v", err)
		return nil, fmt.Errorf("failed to query project list: %w", err)
	}

	// Step 4: Map database results to response structure
	projectData := make([]*sms.ProjectListForFrontResData, len(projects))
	for i, project := range projects {
		projectData[i] = &sms.ProjectListForFrontResData{
			ProjectName: project.ProjectName,
			ProjectId:   project.Id,
		}
	}

	// Step 5: Prepare and return the response
	response := &sms.ProjectListForFrontRes{
		Data: projectData,
	}

	g.Log().Infof(ctx, "Successfully fetched %d projects for SubUserID=%d", count, req.SubUserID)
	return response, nil
}

// Project List

func (s *sMainControllerProjectManagement) ProjectList(ctx context.Context, req *sms.ProjectListReq) (res *sms.ProjectListRes, err error) {
	// Fetch and filter project list with pagination and keyword search
	dbQuery := dao.ProjectList.Ctx(ctx).
		Page(req.PageNum, req.PageSize).
		Order("id desc")

	// Apply keyword search filter if provided
	if req.KeyWordSearch != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+req.KeyWordSearch+"%")
	}

	// Retrieve data and count
	var projects []entity.ProjectList
	var totalCount int
	if err := dbQuery.ScanAndCount(&projects, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count ProjectList: %v", err)
		return nil, fmt.Errorf("failed to fetch project list: %w", err)
	}

	// Prepare response
	res = &sms.ProjectListRes{
		ListRes: commonApi.ListRes{Total: totalCount},
		Data:    make([]sms.ProjectListResData, len(projects)),
	}

	// Map data to response structure
	for i, project := range projects {
		res.Data[i] = sms.ProjectListResData{
			ID:                  project.Id,
			Name:                project.ProjectName,
			QuantityDevice:      project.QuantityDevice,
			AssociatedAccount:   project.AssociatedAccount,
			AssociatedAccountId: project.AssociatedAccountId,
			Note:                project.Note,
			UpdateTime:          project.UpdateAt.String(),
		}
	}

	return res, nil
}

// Create Project

// CreateProject handles the creation of a new project.
func (s *sMainControllerProjectManagement) CreateProject(ctx context.Context, req *sms.ProjectCreateReq) (res *sms.ProjectCreateRes, err error) {
	// Prepare the project entity
	project := &entity.ProjectList{
		ProjectName: req.ProjectName,
		Note:        req.Note,
	}

	// Check if the project name already exists
	count, err := dao.ProjectList.Ctx(ctx).
		Where("project_name = ?", req.ProjectName).
		Count()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to check for duplicate project name '%s': %v", req.ProjectName, err)
		return nil, fmt.Errorf("failed to check for duplicate project name: %w", err)
	}
	if count > 0 {
		g.Log().Warningf(ctx, "Project name '%s' already exists", req.ProjectName)
		return nil, errors.New("project name already exists")
	}

	// Insert the new project into the database
	rowID, err := dao.ProjectList.Ctx(ctx).
		Data(project).
		InsertAndGetId()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to insert new project '%s': %v", req.ProjectName, err)
		return nil, fmt.Errorf("failed to insert new project: %w", err)
	}

	// Prepare the response
	g.Log().Infof(ctx, "Successfully created project '%s' with ID=%d", req.ProjectName, rowID)
	res = &sms.ProjectCreateRes{
		ID: rowID,
	}

	return res, nil
}

// Delete Project

// DeleteProject handles the deletion of a project after checking constraints.
func (s *sMainControllerProjectManagement) DeleteProject(ctx context.Context, req *sms.ProjectDeleteReq) (res *sms.ProjectDeleteRes, err error) {
	// Fetch project details
	var project entity.ProjectList
	if err := dao.ProjectList.Ctx(ctx).Where("id = ?", req.ProjectID).Scan(&project); err != nil {
		g.Log().Errorf(ctx, "Failed to fetch project details for ProjectID=%d: %v", req.ProjectID, err)
		return nil, fmt.Errorf("failed to fetch project details for ProjectID=%d: %w", req.ProjectID, err)
	}

	// Check if the project is associated with an account
	if project.AssociatedAccountId != 0 {
		g.Log().Warningf(ctx, "Cannot delete project with ProjectID=%d as it is assigned to an account", req.ProjectID)
		return nil, errors.New("the project is assigned to a sub-account and cannot be deleted")
	}

	// Check if the project has any devices associated
	if project.QuantityDevice > 0 {
		g.Log().Warningf(ctx, "Cannot delete project with ProjectID=%d as it has associated devices", req.ProjectID)
		return nil, errors.New("the project has associated devices and cannot be deleted")
	}

	// Proceed to delete the project
	if _, err := dao.ProjectList.Ctx(ctx).Where("id = ?", req.ProjectID).Delete(); err != nil {
		g.Log().Errorf(ctx, "Failed to delete project with ProjectID=%d: %v", req.ProjectID, err)
		return nil, fmt.Errorf("failed to delete project with ProjectID=%d: %w", req.ProjectID, err)
	}

	g.Log().Infof(ctx, "Successfully deleted project with ProjectID=%d", req.ProjectID)
	return &sms.ProjectDeleteRes{}, nil
}

// UpdateProject handles updating the details of a project.
func (s *sMainControllerProjectManagement) UpdateProject(ctx context.Context, req *sms.ProjectUpdateReq) (res *sms.ProjectUpdateRes, err error) {
	// Check if the project name already exists
	existingProjects := make([]entity.ProjectList, 0)
	if err := dao.ProjectList.Ctx(ctx).
		Where("project_name = ?", req.ProjectName).
		Scan(&existingProjects); err != nil {
		g.Log().Errorf(ctx, "Failed to check for duplicate project name '%s': %v", req.ProjectName, err)
		return nil, fmt.Errorf("failed to check for duplicate project name: %w", err)
	}

	// Validate if the name already exists
	if len(existingProjects) > 0 {
		g.Log().Warningf(ctx, "Project name '%s' already exists", req.ProjectName)
		return nil, errors.New("project name already exists")
	}

	// Update the project details in the database
	if _, err := dao.ProjectList.Ctx(ctx).
		Data(g.Map{
			"project_name": req.ProjectName,
			"note":         req.Note,
		}).
		Where("id = ?", req.ProjectID).
		Update(); err != nil {
		g.Log().Errorf(ctx, "Failed to update project with ID=%d: %v", req.ProjectID, err)
		return nil, fmt.Errorf("failed to update project: %w", err)
	}

	// Log success and prepare response
	g.Log().Infof(ctx, "Successfully updated project with ID=%d", req.ProjectID)
	return &sms.ProjectUpdateRes{}, nil
}

// AllocateAccount2Project

// AllocateAccount2Project assigns an account to a specific project and updates related device and project records.
//
// Parameters:
// - ctx: The context for handling the request.
// - req: The request containing the ProjectID and AccountID.
//
// Returns:
// - *sms.AllocateAccount2ProjectRes: The response indicating success.
// - error: An error if the operation fails.
func (s *sMainControllerProjectManagement) AllocateAccount2Project(ctx context.Context, req *sms.AllocateAccount2ProjectReq) (*sms.AllocateAccount2ProjectRes, error) {
	// Step 1: Fetch project details
	var project entity.ProjectList
	if err := dao.ProjectList.Ctx(ctx).
		Where("id = ?", req.ProjectId).
		Scan(&project); err != nil {
		g.Log().Errorf(ctx, "Failed to query ProjectList for ProjectId=%d: %v", req.ProjectId, err)
		return nil, fmt.Errorf("failed to query project details: %w", err)
	}

	// Step 2: Fetch account details
	var account entity.User
	if err := dao.User.Ctx(ctx).
		Where("id = ?", req.AccountId).
		Scan(&account); err != nil {
		g.Log().Errorf(ctx, "Failed to query User for AccountId=%d: %v", req.AccountId, err)
		return nil, fmt.Errorf("failed to query account details: %w", err)
	}
	accountName := account.Name

	// Step 3: Update device records for the assigned project
	if _, err := dao.DeviceList.Ctx(ctx).
		Data(g.Map{
			"owner_account":    accountName,
			"owner_account_id": req.AccountId,
		}).
		Where("assigned_items_id = ?", req.ProjectId).
		Update(); err != nil {
		g.Log().Errorf(ctx, "Failed to update DeviceList for ProjectId=%d and AccountId=%d: %v", req.ProjectId, req.AccountId, err)
		return nil, fmt.Errorf("failed to update device records: %w", err)
	}

	// Step 4: Update project details with associated account
	if _, err := dao.ProjectList.Ctx(ctx).
		Data(g.Map{
			"associated_account_id": req.AccountId,
			"associated_account":    accountName,
		}).
		Where("id = ?", req.ProjectId).
		Update(); err != nil {
		g.Log().Errorf(ctx, "Failed to update ProjectList for ProjectId=%d and AccountId=%d: %v", req.ProjectId, req.AccountId, err)
		return nil, fmt.Errorf("failed to update project details: %w", err)
	}

	// Step 5: Log success and return response
	g.Log().Infof(ctx, "Successfully allocated AccountId=%d to ProjectId=%d", req.AccountId, req.ProjectId)
	return &sms.AllocateAccount2ProjectRes{}, nil
}
