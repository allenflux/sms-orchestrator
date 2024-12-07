// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectTaskLinkDeviceGroup is the golang structure for table collect_task_link_device_group.
type CollectTaskLinkDeviceGroup struct {
	Id            uint        `json:"id"              orm:"id"              description:""`       //
	CollectTaskId int         `json:"collect_task_id" orm:"collect_task_id" description:"采集任务id"` // 采集任务id
	DeviceGroupId int         `json:"device_group_id" orm:"device_group_id" description:"设备组id"`  // 设备组id
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`   // 创建时间
	DeletedAt     *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:"删除时间"`   // 删除时间
}
