// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysLoginLog is the golang structure for table ent_sys_login_log.
type EntSysLoginLog struct {
	InfoId        int64       `json:"info_id"        orm:"info_id"        description:"访问ID"`          // 访问ID
	LoginName     string      `json:"login_name"     orm:"login_name"     description:"登录账号"`          // 登录账号
	Ipaddr        string      `json:"ipaddr"         orm:"ipaddr"         description:"登录IP地址"`        // 登录IP地址
	LoginLocation string      `json:"login_location" orm:"login_location" description:"登录地点"`          // 登录地点
	Browser       string      `json:"browser"        orm:"browser"        description:"浏览器类型"`         // 浏览器类型
	Os            string      `json:"os"             orm:"os"             description:"操作系统"`          // 操作系统
	Status        int         `json:"status"         orm:"status"         description:"登录状态（0成功 1失败）"` // 登录状态（0成功 1失败）
	Msg           string      `json:"msg"            orm:"msg"            description:"提示消息"`          // 提示消息
	LoginTime     *gtime.Time `json:"login_time"     orm:"login_time"     description:"登录时间"`          // 登录时间
	Module        string      `json:"module"         orm:"module"         description:"登录模块"`          // 登录模块
}
