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
		// TaskCreate handles the creation of a new task, including file parsing, validation, and distribution.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing task details, group ID, and the uploaded file.
		//
		// Returns:
		// - *sms.SubTaskCreateRes: The response containing the created task ID.
		// - error: An error if the operation fails.
		TaskCreate(ctx context.Context, req *sms.SubTaskCreateReq) (*sms.SubTaskCreateRes, error)
		TaskFileDownload(ctx context.Context, req *sms.TaskFileDownloadReq) (res *sms.TaskFileDownloadRes, err error)
		// TaskReportDelete handles the recall of a task by updating its status in the SmsMissionReport table.
		TaskReportDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (*sms.SubTaskDeleteRes, error)
		// GetSubGetConversationRecord retrieves a detailed conversation record list for a specific ChatLogID.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing the ChatLogID.
		//
		// Returns:
		// - *sms.SubGetConversationRecordRes: The response containing the conversation record list.
		// - error: An error if the operation fails.
		GetSubGetConversationRecord(ctx context.Context, req *sms.SubGetConversationRecordReq) (*sms.SubGetConversationRecordRes, error)
		// SubGetConversationRecordList retrieves a paginated list of conversation records for a specific project.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing pagination, project ID, and optional SubUserID.
		//
		// Returns:
		// - *sms.SubGetConversationRecordListRes: The response containing the conversation records.
		// - error: An error if the operation fails.
		SubGetConversationRecordList(ctx context.Context, req *sms.SubGetConversationRecordListReq) (*sms.SubGetConversationRecordListRes, error)
		GetTaskRecordList(ctx context.Context, req *sms.SubTaskRecordReq) (res *sms.SubTaskRecordRes, err error)
		// PostConversationRecord receives input from a Sub User and stores it in the cache as a priority task for a Device to execute.
		PostConversationRecord(ctx context.Context, req *sms.SubPostConversationRecordReq) (*sms.SubPostConversationRecordRes, error)
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
		// GetTaskDevices retrieves the list of devices assigned to a specific task by delegating the request to the main service.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing task ID and pagination details.
		//
		// Returns:
		// - *sms.SubTaskDevicesRes: The response containing the list of devices and pagination metadata.
		// - error: An error if the operation fails.
		GetTaskDevices(ctx context.Context, req *sms.SubTaskDevicesReq) (*sms.SubTaskDevicesRes, error)
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
