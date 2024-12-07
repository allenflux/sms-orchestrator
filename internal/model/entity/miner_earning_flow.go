// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerEarningFlow is the golang structure for table miner_earning_flow.
type MinerEarningFlow struct {
	Id          uint        `json:"id"            orm:"id"            description:""`            //
	DeviceDevId string      `json:"device_dev_id" orm:"device_dev_id" description:"设备本机id"`      // 设备本机id
	Platform    string      `json:"platform"      orm:"platform"      description:"平台"`          // 平台
	Account     string      `json:"account"       orm:"account"       description:"账号"`          // 账号
	Source      int         `json:"source"        orm:"source"        description:"来源0代表无1视频任务"` // 来源0代表无1视频任务
	Type        int         `json:"type"          orm:"type"          description:"类型1收益2提现"`    // 类型1收益2提现
	Money       float64     `json:"money"         orm:"money"         description:"金额"`          // 金额
	OldBalance  float64     `json:"old_balance"   orm:"old_balance"   description:"变更前金额"`       // 变更前金额
	NewBalance  float64     `json:"new_balance"   orm:"new_balance"   description:"变更后金额"`       // 变更后金额
	Remark      string      `json:"remark"        orm:"remark"        description:"备注"`          // 备注
	CreatedAt   *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`        // 创建时间
	CreatedBy   int         `json:"created_by"    orm:"created_by"    description:"创建者id"`       // 创建者id
}
