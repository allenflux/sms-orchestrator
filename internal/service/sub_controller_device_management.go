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
	ISubControllerDeviceManagement interface {
		GetDeviceList(ctx context.Context, req *sms.SubDeviceListReq) (res *sms.SubDeviceListRes, err error)
		GroupCreate(ctx context.Context, req *sms.SubCreateGroupReq) (res *sms.SubCreateGroupRes, err error)
		GroupUpdate(ctx context.Context, req *sms.SubUpdateGroupReq) (res *sms.SubUpdateGroupRes, err error)
		// DeleteGroup deletes a group if there are no devices associated with it.
		DeleteGroup(ctx context.Context, req *sms.SubDeleteGroupReq) (res *sms.SubDeleteGroupRes, err error)
		GroupList(ctx context.Context, req *sms.SubGroupListReq) (res *sms.SubGroupListRes, err error)
		// AllocateDevice2Group allocates a list of devices to a specified group.
		// It validates the group ID, sub-user ID, and device IDs before updating the database.
		// The function uses a transaction to ensure consistency.
		//
		// Parameters:
		// - ctx: The context for handling the operation.
		// - req: The request containing the group ID, sub-user ID, and device IDs.
		//
		// Returns:
		// - *sms.AllocateDevice2GroupRes: The response indicating success of the allocation.
		// - error: An error if the allocation fails at any point.
		AllocateDevice2Group(ctx context.Context, req *sms.AllocateDevice2GroupReq) (*sms.AllocateDevice2GroupRes, error)
		// GetProjectList retrieves the list of projects for the front-end, scoped by the SubUserID.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing the SubUserID.
		//
		// Returns:
		// - *sms.SubProjectListForFrontRes: The response containing the list of projects.
		// - error: An error if the operation fails.
		GetProjectList(ctx context.Context, req *sms.SubProjectListForFrontReq) (*sms.SubProjectListForFrontRes, error)
	}
)

var (
	localSubControllerDeviceManagement ISubControllerDeviceManagement
)

func SubControllerDeviceManagement() ISubControllerDeviceManagement {
	if localSubControllerDeviceManagement == nil {
		panic("implement not found for interface ISubControllerDeviceManagement, forgot register?")
	}
	return localSubControllerDeviceManagement
}

func RegisterSubControllerDeviceManagement(i ISubControllerDeviceManagement) {
	localSubControllerDeviceManagement = i
}
