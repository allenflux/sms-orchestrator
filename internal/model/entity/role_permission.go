// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermission is the golang structure for table role_permission.
type RolePermission struct {
	Id           int         `json:"id"            orm:"id"            description:""`     //
	RoleId       int         `json:"role_id"       orm:"role_id"       description:"角色id"` // 角色id
	PermissionId int         `json:"permission_id" orm:"permission_id" description:"权限id"` // 权限id
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"` // 创建时间
	UpdateAt     *gtime.Time `json:"update_at"     orm:"update_at"     description:"修改时间"` // 修改时间
	DeleteAt     *gtime.Time `json:"delete_at"     orm:"delete_at"     description:"删除时间"` // 删除时间
}
