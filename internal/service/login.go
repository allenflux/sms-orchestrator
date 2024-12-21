// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/allUser"
	"sms_backend/internal/model"
)

type (
	ILogin interface {
		Login(ctx context.Context, req *allUser.LoginReq) (res *allUser.LoginRes, err error)
		ChangePassword(ctx context.Context, req *allUser.ChangePasswordReq) (res *allUser.ChangePasswordRes, err error)
		Logout(ctx context.Context, req *allUser.LogoutReq) (res *allUser.LogoutRes, err error)
		UserListTree(pid int, list []*model.UserPermission) []*model.UserPermissionTree
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
