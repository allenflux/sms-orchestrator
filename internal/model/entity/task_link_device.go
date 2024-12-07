// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskLinkDevice is the golang structure for table task_link_device.
type TaskLinkDevice struct {
	Id         int         `json:"id"          orm:"id"          description:""`                                  //
	TaskId     int         `json:"task_id"     orm:"task_id"     description:"任务ID"`                              // 任务ID
	DeviceId   int         `json:"device_id"   orm:"device_id"   description:"任务ID"`                              // 任务ID
	Num        int         `json:"num"         orm:"num"         description:"执行次数"`                              // 执行次数
	Status     int         `json:"status"      orm:"status"      description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CompleteAt *gtime.Time `json:"complete_at" orm:"complete_at" description:"完成时间"`                              // 完成时间
	CreatedAt  *gtime.Time `json:"created_at"  orm:"created_at"  description:""`                                  //
	UpdatedAt  *gtime.Time `json:"updated_at"  orm:"updated_at"  description:""`                                  //
}
