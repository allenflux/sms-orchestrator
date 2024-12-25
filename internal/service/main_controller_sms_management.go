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
		GetConversationRecordList(ctx context.Context, req *sms.ConversationListReq) (res *sms.ConversationListRes, err error)
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
		GetPendingTasks(ctx context.Context, req *sms.PendingTaskReq) (res *sms.PendingTaskRes, err error)
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
