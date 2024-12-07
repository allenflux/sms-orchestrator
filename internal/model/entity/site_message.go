// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteMessage is the golang structure for table site_message.
type SiteMessage struct {
	Id            int         `json:"id"               orm:"id"               description:""`             //
	Title         string      `json:"title"            orm:"title"            description:"标题"`           // 标题
	Content       string      `json:"content"          orm:"content"          description:"消息"`           // 消息
	AppUserIdStrs string      `json:"app_user_id_strs" orm:"app_user_id_strs" description:"APP用户ID，逗号隔开"` // APP用户ID，逗号隔开
	AdminUsername string      `json:"admin_username"   orm:"admin_username"   description:"管理员账号"`        // 管理员账号
	HasAll        int         `json:"has_all"          orm:"has_all"          description:"0-非所有  1-所有"`  // 0-非所有  1-所有
	CreatedAt     *gtime.Time `json:"created_at"       orm:"created_at"       description:""`             //
	UpdatedAt     *gtime.Time `json:"updated_at"       orm:"updated_at"       description:""`             //
}
