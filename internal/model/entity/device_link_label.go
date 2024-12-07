// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceLinkLabel is the golang structure for table device_link_label.
type DeviceLinkLabel struct {
	Id            int         `json:"id"              orm:"id"              description:""`     //
	DeviceLabelId int         `json:"device_label_id" orm:"device_label_id" description:"标签ID"` // 标签ID
	DevId         string      `json:"dev_id"          orm:"dev_id"          description:"设备ID"` // 设备ID
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:""`     //
	DeletedAt     *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:""`     //
}
