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
	ISubControllerSmsManagement interface {
		GetTaskList(ctx context.Context, req *sms.SubTaskListReq) (res *sms.SubTaskListRes, err error)
		TaskCreate(ctx context.Context, req *sms.SubTaskCreateReq) (res *sms.SubTaskListRes, err error)
		// Download File
		TaskFileDownload(ctx context.Context, req *sms.TaskFileDownloadReq) (res *sms.TaskFileDownloadRes, err error)
		TaskReportDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (res *sms.SubTaskDeleteRes, err error)
		GetTaskRecordList(ctx context.Context, req *sms.SubTaskRecordReq) (res *sms.SubTaskRecordRes, err error)
	}
)

var (
	localSubControllerSmsManagement ISubControllerSmsManagement
)

func SubControllerSmsManagement() ISubControllerSmsManagement {
	if localSubControllerSmsManagement == nil {
		panic("implement not found for interface ISubControllerSmsManagement, forgot register?")
	}
	return localSubControllerSmsManagement
}

func RegisterSubControllerSmsManagement(i ISubControllerSmsManagement) {
	localSubControllerSmsManagement = i
}
