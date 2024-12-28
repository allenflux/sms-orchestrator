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
	IMainControllerDeviceManagement interface {
		// GetDeviceList retrieves a paginated list of devices based on the given filters.
		//
		// Parameters:
		// - ctx: The context for handling the request.
		// - req: The request containing the filters for querying the device list.
		//
		// Returns:
		// - *sms.DeviceListRes: The response containing the paginated device list and metadata.
		// - error: An error if the operation fails.
		GetDeviceList(ctx context.Context, req *sms.DeviceListReq) (*sms.DeviceListRes, error)
		// AllocateDevice2Project allocates a list of devices to a project.
		// It validates the device IDs, checks if the project exists, and updates the database accordingly.
		//
		// Parameters:
		// - ctx: The context for handling the operation.
		// - req: The request containing the project ID and list of device IDs to be allocated.
		//
		// Returns:
		// - *sms.AllocateDevice2ProjectRes: The response indicating the success of the operation.
		// - error: An error if the operation fails at any step.
		AllocateDevice2Project(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (*sms.AllocateDevice2ProjectRes, error)
	}
)

var (
	localMainControllerDeviceManagement IMainControllerDeviceManagement
)

func MainControllerDeviceManagement() IMainControllerDeviceManagement {
	if localMainControllerDeviceManagement == nil {
		panic("implement not found for interface IMainControllerDeviceManagement, forgot register?")
	}
	return localMainControllerDeviceManagement
}

func RegisterMainControllerDeviceManagement(i IMainControllerDeviceManagement) {
	localMainControllerDeviceManagement = i
}
