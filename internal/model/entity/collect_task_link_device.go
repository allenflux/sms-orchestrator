// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectTaskLinkDevice is the golang structure for table collect_task_link_device.
type CollectTaskLinkDevice struct {
	Id            uint        `json:"id"              orm:"id"              description:""`                                  //
	CollectTaskId int         `json:"collect_task_id" orm:"collect_task_id" description:"采集任务id"`                            // 采集任务id
	DeviceId      int         `json:"device_id"       orm:"device_id"       description:"设备id"`                              // 设备id
	DeviceDevId   string      `json:"device_dev_id"   orm:"device_dev_id"   description:"设备本机id"`                            // 设备本机id
	Status        int         `json:"status"          orm:"status"          description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	Remark        string      `json:"remark"          orm:"remark"          description:"备注（原因）"`                            // 备注（原因）
	CompletedAt   *gtime.Time `json:"completed_at"    orm:"completed_at"    description:"完成时间"`                              // 完成时间
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`                              // 创建时间
	UpdatedAt     *gtime.Time `json:"updated_at"      orm:"updated_at"      description:"修改时间"`                              // 修改时间
	DeletedAt     *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:"删除时间"`                              // 删除时间
}
