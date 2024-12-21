// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SmsTraceTask is the golang structure of table sms_trace_task for DAO operations like Where/Data.
type SmsTraceTask struct {
	g.Meta       `orm:"table:sms_trace_task, do:true"`
	Id           interface{} //
	TargetNumber interface{} //
	DeviceNumber interface{} //
	TaskId       interface{} //
	CreateAt     *gtime.Time //
	UpdateAt     *gtime.Time //
	DeleteAt     *gtime.Time //
}
