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
	IMainControllerSmsManagement interface {
		GetTaskList(ctx context.Context, req *sms.TaskListReq) (res *sms.TaskListRes, err error)
		GetTaskRecordList(ctx context.Context, req *sms.TaskRecordReq) (res *sms.TaskRecordRes, err error)
	}
)

var (
	localMainControllerSmsManagement IMainControllerSmsManagement
)

func MainControllerSmsManagement() IMainControllerSmsManagement {
	if localMainControllerSmsManagement == nil {
		panic("implement not found for interface IMainControllerSmsManagement, forgot register?")
	}
	return localMainControllerSmsManagement
}

func RegisterMainControllerSmsManagement(i IMainControllerSmsManagement) {
	localMainControllerSmsManagement = i
}
