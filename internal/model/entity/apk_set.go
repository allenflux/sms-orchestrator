// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ApkSet is the golang structure for table apk_set.
type ApkSet struct {
	Id            uint        `json:"id"              orm:"id"              description:""`         //
	SetId         string      `json:"set_id"          orm:"set_id"          description:"集合id"`     // 集合id
	ScriptId      int         `json:"script_id"       orm:"script_id"       description:"脚本id"`     // 脚本id
	ScriptAppName string      `json:"script_app_name" orm:"script_app_name" description:"脚本包名"`     // 脚本包名
	Status        int         `json:"status"          orm:"status"          description:"1未安装2已安装"` // 1未安装2已安装
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`     // 创建时间
	UpdatedAt     *gtime.Time `json:"updated_at"      orm:"updated_at"      description:"修改时间"`     // 修改时间
}
