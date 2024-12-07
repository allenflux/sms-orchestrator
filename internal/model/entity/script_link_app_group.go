// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ScriptLinkAppGroup is the golang structure for table script_link_app_group.
type ScriptLinkAppGroup struct {
	Id         int         `json:"id"           orm:"id"           description:""`        //
	ScriptId   int         `json:"script_id"    orm:"script_id"    description:"脚本ID"`    // 脚本ID
	AppGroupId int         `json:"app_group_id" orm:"app_group_id" description:"app分组ID"` // app分组ID
	CreatedAt  *gtime.Time `json:"created_at"   orm:"created_at"   description:"创建时间"`    // 创建时间
}
