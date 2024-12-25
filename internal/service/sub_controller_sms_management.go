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
		TaskFileDownload(ctx context.Context, req *sms.TaskFileDownloadReq) (res *sms.TaskFileDownloadRes, err error)
		TaskReportDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (res *sms.SubTaskDeleteRes, err error)
		GetSubGetConversationRecord(ctx context.Context, req *sms.SubGetConversationRecordReq) (res *sms.SubGetConversationRecordRes, err error)
		SubGetConversationRecordList(ctx context.Context, req *sms.SubGetConversationRecordListReq) (res *sms.SubGetConversationRecordListRes, err error)
		GetTaskRecordList(ctx context.Context, req *sms.SubTaskRecordReq) (res *sms.SubTaskRecordRes, err error)
		PostConversationRecord(ctx context.Context, req *sms.SubPostConversationRecordReq) (res *sms.SubPostConversationRecordRes, err error)
		// GetPendingTasks retrieves the pending tasks for a specific SMS mission.
		//
		// Parameters:
		// - ctx: The context for the request.
		// - req: The request containing the TaskID.
		//
		// Returns:
		// - *sms.SubPendingTaskRes: The response containing the list of pending tasks.
		// - error: An error if the operation fails.
		GetPendingTasks(ctx context.Context, req *sms.SubPendingTaskReq) (*sms.SubPendingTaskRes, error)
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
