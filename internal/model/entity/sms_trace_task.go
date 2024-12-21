// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsTraceTask is the golang structure for table sms_trace_task.
type SmsTraceTask struct {
	Id           int         `json:"id"            orm:"id"            description:""` //
	TargetNumber string      `json:"target_number" orm:"target_number" description:""` //
	DeviceNumber string      `json:"device_number" orm:"device_number" description:""` //
	TaskId       int         `json:"task_id"       orm:"task_id"       description:""` //
	CreateAt     *gtime.Time `json:"create_at"     orm:"create_at"     description:""` //
	UpdateAt     *gtime.Time `json:"update_at"     orm:"update_at"     description:""` //
	DeleteAt     *gtime.Time `json:"delete_at"     orm:"delete_at"     description:""` //
}
