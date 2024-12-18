// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/user"
)

type (
	IUser interface {
		GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error)
		CreatedUser(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error)
		UpdateUser(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error)
		DeleteUser(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error)
		ChangeStatus(ctx context.Context, req *user.ChangeStatusReq) (res *user.ChangeStatusRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
