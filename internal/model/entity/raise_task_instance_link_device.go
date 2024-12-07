// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseTaskInstanceLinkDevice is the golang structure for table raise_task_instance_link_device.
type RaiseTaskInstanceLinkDevice struct {
	Id                  uint        `json:"id"                     orm:"id"                     description:""`                                  //
	RaiseTaskId         int         `json:"raise_task_id"          orm:"raise_task_id"          description:"养号定时任务id"`                          // 养号定时任务id
	RaiseTaskInstanceId int         `json:"raise_task_instance_id" orm:"raise_task_instance_id" description:"养号任务实例id"`                          // 养号任务实例id
	DeviceGroupId       int         `json:"device_group_id"        orm:"device_group_id"        description:"设备组id"`                             // 设备组id
	DeviceId            int         `json:"device_id"              orm:"device_id"              description:"设备id（数据库里自增的）"`                     // 设备id（数据库里自增的）
	DeviceDevId         string      `json:"device_dev_id"          orm:"device_dev_id"          description:"设备本机id（手机上的）"`                      // 设备本机id（手机上的）
	CreatedAt           *gtime.Time `json:"created_at"             orm:"created_at"             description:"创建时间"`                              // 创建时间
	ExecAt              *gtime.Time `json:"exec_at"                orm:"exec_at"                description:"执行时间"`                              // 执行时间
	CompletedAt         *gtime.Time `json:"completed_at"           orm:"completed_at"           description:"完成时间"`                              // 完成时间
	Status              int         `json:"status"                 orm:"status"                 description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	DeletedAt           *gtime.Time `json:"deleted_at"             orm:"deleted_at"             description:"删除时间"`                              // 删除时间
	Remark              string      `json:"remark"                 orm:"remark"                 description:"原因"`                                // 原因
}
