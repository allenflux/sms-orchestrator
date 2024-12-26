// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/sms"
)

type (
	IMainControllerProjectManagement interface {
		// GetProjectList retrieves the list of projects, optionally filtered by SubUserID.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing optional SubUserID for filtering.
		//
		// Returns:
		// - *sms.ProjectListForFrontRes: The response containing the project list.
		// - error: An error if the operation fails.
		GetProjectList(ctx context.Context, req *sms.ProjectListForFrontReq) (*sms.ProjectListForFrontRes, error)
		ProjectList(ctx context.Context, req *sms.ProjectListReq) (res *sms.ProjectListRes, err error)
		CreateProject(ctx context.Context, req *sms.ProjectCreateReq) (res *sms.ProjectCreateRes, err error)
		DeleteProject(ctx context.Context, req *sms.ProjectDeleteReq) (res *sms.ProjectDeleteRes, err error)
		UpdateProject(ctx context.Context, req *sms.ProjectUpdateReq) (res *sms.ProjectUpdateRes, err error)
		// AllocateAccount2Project assigns an account to a specific project and updates related device and project records.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing the ProjectID and AccountID.
		//
		// Returns:
		// - *sms.AllocateAccount2ProjectRes: The response indicating success.
		// - error: An error if the operation fails.
		AllocateAccount2Project(ctx context.Context, req *sms.AllocateAccount2ProjectReq) (*sms.AllocateAccount2ProjectRes, error)
	}
)

var (
	localMainControllerProjectManagement IMainControllerProjectManagement
)

func MainControllerProjectManagement() IMainControllerProjectManagement {
	if localMainControllerProjectManagement == nil {
		panic("implement not found for interface IMainControllerProjectManagement, forgot register?")
	}
	return localMainControllerProjectManagement
}

func RegisterMainControllerProjectManagement(i IMainControllerProjectManagement) {
	localMainControllerProjectManagement = i
}
