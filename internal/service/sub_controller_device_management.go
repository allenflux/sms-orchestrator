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
		GroupDelete(ctx context.Context, req *sms.SubDeleteGroupReq) (res *sms.SubDeleteGroupRes, err error)
		GroupList(ctx context.Context, req *sms.SubGroupListReq) (res *sms.SubGroupListRes, err error)
		AllocateDevice2Group(ctx context.Context, req *sms.AllocateDevice2GroupReq) (res *sms.AllocateDevice2GroupRes, err error)
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
