// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerSession is the golang structure for table miner_session.
type MinerSession struct {
	Id       uint        `json:"id"        orm:"id"        description:""`         //
	DevId    string      `json:"dev_id"    orm:"dev_id"    description:"设备dev id"` // 设备dev id
	LoginAt  *gtime.Time `json:"login_at"  orm:"login_at"  description:"登录时间"`     // 登录时间
	LogoffAt *gtime.Time `json:"logoff_at" orm:"logoff_at" description:"登出时间"`     // 登出时间
}
