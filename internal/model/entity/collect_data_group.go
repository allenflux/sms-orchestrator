// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectDataGroup is the golang structure for table collect_data_group.
type CollectDataGroup struct {
	Id        uint        `json:"id"         orm:"id"         description:""`       //
	Name      string      `json:"name"       orm:"name"       description:"名称"`     // 名称
	ClientId  int         `json:"client_id"  orm:"client_id"  description:"创建客户id"` // 创建客户id
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`   // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"修改时间"`   // 修改时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`   // 删除时间
}
