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
		// GetTaskRecordList retrieves a paginated list of task records based on the provided filters.
		GetTaskRecordList(ctx context.Context, req *sms.TaskRecordReq) (*sms.TaskRecordRes, error)
		// GetSubGetConversationRecord retrieves a specific conversation record by delegating the request to the sub-controller service.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing the ChatLogID.
		//
		// Returns:
		// - *sms.ConversationRecordRes: The response containing the conversation record.
		// - error: An error if the operation fails.
		GetSubGetConversationRecord(ctx context.Context, req *sms.ConversationRecordReq) (*sms.ConversationRecordRes, error)
		// GetConversationRecordList retrieves the list of conversation records by delegating the request to the sub-controller service.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing pagination details and the ProjectID.
		//
		// Returns:
		// - *sms.ConversationListRes: The response containing the conversation records.
		// - error: An error if the operation fails.
		GetConversationRecordList(ctx context.Context, req *sms.ConversationListReq) (*sms.ConversationListRes, error)
		// PostConversationRecord handles posting conversation records and forwards the request to the sub-controller.
		PostConversationRecord(ctx context.Context, req *sms.PostConversationRecordReq) (*sms.PostConversationRecordRes, error)
		// GetTaskDevices retrieves the list of devices assigned to a specific task (group ID) with pagination support.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing the task ID and pagination details.
		//
		// Returns:
		// - *sms.TaskDevicesRes: The response containing the list of devices and pagination metadata.
		// - error: An error if the operation fails.
		GetTaskDevices(ctx context.Context, req *sms.TaskDevicesReq) (*sms.TaskDevicesRes, error)
		// GetPendingTasks retrieves pending tasks by delegating the request to the sub-controller service.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing the TaskID.
		//
		// Returns:
		// - *sms.PendingTaskRes: The response containing the pending tasks and metadata.
		// - error: An error if the operation fails.
		GetPendingTasks(ctx context.Context, req *sms.PendingTaskReq) (*sms.PendingTaskRes, error)
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
