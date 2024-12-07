// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/system"
	"sms_backend/internal/model/entity"
)

type (
	ISysRole interface {
		GetRoleListSearch(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error)
		// GetRoleList 获取角色列表
		GetRoleList(ctx context.Context) (list []*entity.SysRole, err error)
		// AddRoleRule 添加角色权限
		AddRoleRule(ctx context.Context, ruleIds []int, roleId int) (err error)
		// DelRoleRule 删除角色权限
		DelRoleRule(ctx context.Context, roleId int) (err error)
		AddRole(ctx context.Context, req *system.RoleAddReq) (err error)
		Get(ctx context.Context, id int) (res *entity.SysRole, err error)
		// GetFilteredNamedPolicy 获取角色关联的菜单规则
		GetFilteredNamedPolicy(ctx context.Context, id int) (gpSlice []int, err error)
		// EditRole 修改角色
		EditRole(ctx context.Context, req *system.RoleEditReq) (err error)
		// DeleteByIds 删除角色
		DeleteByIds(ctx context.Context, ids []int) (err error)
		GetIdNameMap(ctx context.Context) (res map[int]string, err error)
		GetList(ctx context.Context) (res []*entity.SysRole, err error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
