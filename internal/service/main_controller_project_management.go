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
		ProjectList(ctx context.Context, req *sms.ProjectListReq) (res *sms.ProjectListRes, err error)
		CreateProject(ctx context.Context, req *sms.ProjectCreateReq) (res *sms.ProjectCreateRes, err error)
		DeleteProject(ctx context.Context, req *sms.ProjectDeleteReq) (res *sms.ProjectDeleteRes, err error)
		UpdateProject(ctx context.Context, req *sms.ProjectUpdateReq) (res *sms.ProjectUpdateRes, err error)
		AllocateAccount2Project(ctx context.Context, req *sms.AllocateAccount2ProjectReq) (res *sms.AllocateAccount2ProjectRes, err error)
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
