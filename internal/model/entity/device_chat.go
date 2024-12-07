// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceChat is the golang structure for table device_chat.
type DeviceChat struct {
	Id           uint        `json:"id"             orm:"id"             description:""`   //
	DeviceId     int         `json:"device_id"      orm:"device_id"      description:"id"` // id
	DeviceDevId  string      `json:"device_dev_id"  orm:"device_dev_id"  description:"id"` // id
	LastChatTime *gtime.Time `json:"last_chat_time" orm:"last_chat_time" description:""`   //
}
