// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/role"
)

type (
	IRole interface {
		CreatedRole(ctx context.Context, req *role.CreatedReq) (res *role.CreatedRes, err error)
		GetList(ctx context.Context, req *role.ListReq) (res *role.ListRes, err error)
		UpdatedRole(ctx context.Context, req *role.UpdatedReq) (res *role.UpdatedRes, err error)
		DeleteRole(ctx context.Context, req *role.DeletedReq) (res *role.DeletedRes, err error)
		GetPermissionList(ctx context.Context, req *role.GetPermissionReq) (res *role.GetPermissionRes, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
