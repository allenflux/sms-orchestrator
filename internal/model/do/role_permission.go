// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermission is the golang structure of table role_permission for DAO operations like Where/Data.
type RolePermission struct {
	g.Meta       `orm:"table:role_permission, do:true"`
	Id           interface{} //
	RoleId       interface{} // 角色id
	PermissionId interface{} // 权限id
	CreatedAt    *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 修改时间
	DeleteAt     *gtime.Time // 删除时间
}
