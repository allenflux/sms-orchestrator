// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceLinkGroup is the golang structure for table device_link_group.
type DeviceLinkGroup struct {
	Id        int         `json:"id"         orm:"id"         description:""`       //
	DeviceId  int         `json:"device_id"  orm:"device_id"  description:"设备ID"`   // 设备ID
	GroupId   int         `json:"group_id"   orm:"group_id"   description:"设备分组ID"` // 设备分组ID
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`   // 创建时间
	ClientId  int         `json:"client_id"  orm:"client_id"  description:"客户ID"`   // 客户ID
}
