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
		DeviceRegister(ctx context.Context, req *career.RegisterReq) (res *career.RegisterRes, err error)
		FetchTasks(ctx context.Context, req *career.FetchTaskReq) (res *career.FetchTaskRes, err error)
		ReportTaskResult(ctx context.Context, req *career.ReportTaskResultReq) (res *career.ReportTaskResultRes, err error)
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
