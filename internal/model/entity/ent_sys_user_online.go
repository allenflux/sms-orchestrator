// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysUserOnline is the golang structure for table ent_sys_user_online.
type EntSysUserOnline struct {
	Id         int64       `json:"id"          orm:"id"          description:""`        //
	Uuid       string      `json:"uuid"        orm:"uuid"        description:"用户标识"`    // 用户标识
	Token      string      `json:"token"       orm:"token"       description:"用户token"` // 用户token
	CreateTime *gtime.Time `json:"create_time" orm:"create_time" description:"登录时间"`    // 登录时间
	UserName   string      `json:"user_name"   orm:"user_name"   description:"用户名"`     // 用户名
	Ip         string      `json:"ip"          orm:"ip"          description:"登录ip"`    // 登录ip
	Explorer   string      `json:"explorer"    orm:"explorer"    description:"浏览器"`     // 浏览器
	Os         string      `json:"os"          orm:"os"          description:"操作系统"`    // 操作系统
}
