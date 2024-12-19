// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        int         `json:"id"         orm:"id"         description:""`     //
	Name      string      `json:"name"       orm:"name"       description:"角色名称"` // 角色名称
	Note      string      `json:"note"       orm:"note"       description:"备注"`   // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	UpdateAt  *gtime.Time `json:"update_at"  orm:"update_at"  description:"修改时间"` // 修改时间
	DeleteAt  *gtime.Time `json:"delete_at"  orm:"delete_at"  description:"删除时间"` // 删除时间
}
