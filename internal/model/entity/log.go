// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Log is the golang structure for table log.
type Log struct {
	Id        int         `json:"id"         orm:"id"         description:""`         //
	UserId    int         `json:"user_id"    orm:"user_id"    description:"操作者的用户id"` // 操作者的用户id
	UserName  string      `json:"user_name"  orm:"user_name"  description:"操作者的用户名"`  // 操作者的用户名
	ClientIp  string      `json:"client_ip"  orm:"client_ip"  description:"操作ip"`     // 操作ip
	Function  string      `json:"function"   orm:"function"   description:"操作功能"`     // 操作功能
	Note      string      `json:"note"       orm:"note"       description:"操作内容"`     // 操作内容
	SystemId  int         `json:"system_id"  orm:"system_id"  description:"后台id"`     // 后台id
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`     // 创建时间
	UpdateAt  *gtime.Time `json:"update_at"  orm:"update_at"  description:"修改时间"`     // 修改时间
	DeleteAt  *gtime.Time `json:"delete_at"  orm:"delete_at"  description:"删除时间"`     // 删除时间
}
