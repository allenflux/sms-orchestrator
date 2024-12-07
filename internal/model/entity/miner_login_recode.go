// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerLoginRecode is the golang structure for table miner_login_recode.
type MinerLoginRecode struct {
	Id        int         `json:"id"         orm:"id"         description:""`            //
	DeviceId  int         `json:"device_id"  orm:"device_id"  description:"设备主键ID"`      // 设备主键ID
	Platform  string      `json:"platform"   orm:"platform"   description:"平台"`          // 平台
	Status    int         `json:"status"     orm:"status"     description:"1 未登录 2 已登录"` // 1 未登录 2 已登录
	ClientId  int         `json:"client_id"  orm:"client_id"  description:"客户ID"`        // 客户ID
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`          // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`            //
	DevId     string      `json:"dev_id"     orm:"dev_id"     description:"设备ID"`        // 设备ID
}
