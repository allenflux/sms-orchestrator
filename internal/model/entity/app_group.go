// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppGroup is the golang structure for table app_group.
type AppGroup struct {
	Id        int         `json:"id"         orm:"id"         description:""`     //
	Name      string      `json:"name"       orm:"name"       description:"分组名称"` // 分组名称
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`     //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`     //
	ClientId  int         `json:"client_id"  orm:"client_id"  description:"客户ID"` // 客户ID
	RoleId    int         `json:"role_id"    orm:"role_id"    description:"角色ID"` // 角色ID
}
