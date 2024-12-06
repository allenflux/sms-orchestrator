// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"upay_backend/api/v1/admin"
)

type (
	IRecharge interface {
		RechargeList(ctx context.Context, req *admin.RechargeListReq) (res *admin.RechargeListRes, err error)
		RechargeListCreate(ctx context.Context, req *admin.RechargeListCreateParams) (err error)
	}
)

var (
	localRecharge IRecharge
)

func Recharge() IRecharge {
	if localRecharge == nil {
		panic("implement not found for interface IRecharge, forgot register?")
	}
	return localRecharge
}

func RegisterRecharge(i IRecharge) {
	localRecharge = i
}
