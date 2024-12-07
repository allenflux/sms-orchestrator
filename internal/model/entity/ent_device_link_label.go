// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntDeviceLinkLabel is the golang structure for table ent_device_link_label.
type EntDeviceLinkLabel struct {
	Id            int         `json:"id"              orm:"id"              description:""`     //
	DeviceLabelId int         `json:"device_label_id" orm:"device_label_id" description:"标签ID"` // 标签ID
	DevId         string      `json:"dev_id"          orm:"dev_id"          description:"设备ID"` // 设备ID
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:""`     //
}
