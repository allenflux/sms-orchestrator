// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceGrant is the golang structure for table device_grant.
type DeviceGrant struct {
	Id        uint        `json:"id"         orm:"id"         description:""`       //
	DevId     string      `json:"dev_id"     orm:"dev_id"     description:"设备本机id"` // 设备本机id
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`     // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`   // 创建时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`   // 删除时间
}
