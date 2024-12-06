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
		CetCardList(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (res *sms.AllocateDevice2ProjectRes, err error)
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
