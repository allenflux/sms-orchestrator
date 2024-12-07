// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceLabel is the golang structure for table device_label.
type DeviceLabel struct {
	Id        uint        `json:"id"         orm:"id"         description:""`     //
	Name      string      `json:"name"       orm:"name"       description:"标签名称"` // 标签名称
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`     //
}
