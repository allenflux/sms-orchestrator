// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceRemarks is the golang structure for table device_remarks.
type DeviceRemarks struct {
	Id        int         `json:"id"         orm:"id"         description:""`     //
	DeviceId  int         `json:"device_id"  orm:"device_id"  description:"设备ID"` // 设备ID
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`   // 备注
	CreatedBy int         `json:"created_by" orm:"created_by" description:"创建者"`  // 创建者
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
}
