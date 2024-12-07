// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskLinkDeviceGroup is the golang structure for table task_link_device_group.
type TaskLinkDeviceGroup struct {
	Id        uint        `json:"id"         orm:"id"         description:""`       //
	TaskId    int         `json:"task_id"    orm:"task_id"    description:"任务id"`   // 任务id
	GroupId   int         `json:"group_id"   orm:"group_id"   description:"设备分组id"` // 设备分组id
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`   // 创建时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`   // 删除时间
}
