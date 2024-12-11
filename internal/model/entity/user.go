// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"         orm:"id"         description:""`               //
	Name      string      `json:"name"       orm:"name"       description:"用户名"`            // 用户名
	Password  string      `json:"password"   orm:"password"   description:"用户密码"`           // 用户密码
	Salt      string      `json:"salt"       orm:"salt"       description:"密码盐"`            // 密码盐
	Status    int         `json:"status"     orm:"status"     description:"账号状态，1-停用 2-启用"` // 账号状态，1-停用 2-启用
	RoleId    int         `json:"role_id"    orm:"role_id"    description:"角色权限id"`         // 角色权限id
	SystemId  int         `json:"system_id"  orm:"system_id"  description:"所属后台id"`         // 所属后台id
	Note      string      `json:"note"       orm:"note"       description:"备注"`             // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`           // 创建时间
	UpdateAt  *gtime.Time `json:"update_at"  orm:"update_at"  description:"更新时间"`           // 更新时间
	DeleteAt  *gtime.Time `json:"delete_at"  orm:"delete_at"  description:"删除时间"`           // 删除时间
}
