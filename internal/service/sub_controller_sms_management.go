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
		TaskCreate(ctx context.Context, req *sms.SubTaskCreateReq) (res *sms.SubTaskCreateRes, err error)
		// Download File
		TaskFileDownload(ctx context.Context, req *sms.TaskFileDownloadReq) (res *sms.TaskFileDownloadRes, err error)
		TaskReportDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (res *sms.SubTaskDeleteRes, err error)
		GetSubGetConversationRecord(ctx context.Context, req *sms.SubGetConversationRecordReq) (res *sms.SubGetConversationRecordRes, err error)
		SubGetConversationRecordList(ctx context.Context, req *sms.SubGetConversationRecordListReq) (res *sms.SubGetConversationRecordListRes, err error)
		GetTaskRecordList(ctx context.Context, req *sms.SubTaskRecordReq) (res *sms.SubTaskRecordRes, err error)
		PostConversationRecord(ctx context.Context, req *sms.SubPostConversationRecordReq) (res *sms.SubPostConversationRecordRes, err error)
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
