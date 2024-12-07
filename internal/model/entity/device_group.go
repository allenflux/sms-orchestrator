// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceGroup is the golang structure for table device_group.
type DeviceGroup struct {
	Id        int         `json:"id"         orm:"id"         description:""`     //
	Name      string      `json:"name"       orm:"name"       description:"分组名称"` // 分组名称
	DeviceNum int         `json:"device_num" orm:"device_num" description:"设备数量"` // 设备数量
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`   // 备注
	ClientId  int         `json:"client_id"  orm:"client_id"  description:"客户id"` // 客户id
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`     //
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`     //
}
