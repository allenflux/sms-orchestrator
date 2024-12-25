// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/career"
)

type (
	ICareerDeviceManagement interface {
		// DeviceRegister handles the registration of a new device.
		// It checks if the device is already registered in the database or exists in the cache,
		// and if not, registers it in the database and removes it from the cache.
		//
		// Parameters:
		// - ctx: The context for handling the operation.
		// - req: The request payload containing the device details (phone number, device number, and active time).
		//
		// Returns:
		// - *career.RegisterRes: The response containing the registered device ID.
		// - error: An error if the operation fails at any point.
		DeviceRegister(ctx context.Context, req *career.RegisterReq) (*career.RegisterRes, error)
		FetchTasks(ctx context.Context, req *career.FetchTaskReq) (*career.FetchTaskRes, error)
		ReportTaskResult(ctx context.Context, req *career.ReportTaskResultReq) (*career.ReportTaskResultRes, error)
		// ReportReceiveContent handles the received SMS content by adding it to the processing queue
		// and waiting for the result from the result channel.
		//
		// Parameters:
		// - ctx: The context for handling the operation.
		// - req: The request containing the received SMS content details.
		//
		// Returns:
		// - *career.ReportReceiveContentRes: The result of the SMS processing.
		// - error: An error if the processing fails.
		ReportReceiveContent(ctx context.Context, req *career.ReportReceiveContentReq) (*career.ReportReceiveContentRes, error)
	}
)

var (
	localCareerDeviceManagement ICareerDeviceManagement
)

func CareerDeviceManagement() ICareerDeviceManagement {
	if localCareerDeviceManagement == nil {
		panic("implement not found for interface ICareerDeviceManagement, forgot register?")
	}
	return localCareerDeviceManagement
}

func RegisterCareerDeviceManagement(i ICareerDeviceManagement) {
	localCareerDeviceManagement = i
}
