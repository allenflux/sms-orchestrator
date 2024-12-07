// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatTaskLinkDevice is the golang structure for table chat_task_link_device.
type ChatTaskLinkDevice struct {
	Id            uint        `json:"id"              orm:"id"              description:""`                                  //
	ChatTaskId    int         `json:"chat_task_id"    orm:"chat_task_id"    description:"私信任务id"`                            // 私信任务id
	DeviceId      int         `json:"device_id"       orm:"device_id"       description:"设备id"`                              // 设备id
	DeviceDevId   string      `json:"device_dev_id"   orm:"device_dev_id"   description:"id"`                                // id
	DeviceGroupId int         `json:"device_group_id" orm:"device_group_id" description:"设备组id"`                             // 设备组id
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`                              // 创建时间
	CompletedAt   *gtime.Time `json:"completed_at"    orm:"completed_at"    description:"完成时间"`                              // 完成时间
	Status        int         `json:"status"          orm:"status"          description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	FailedNum     int         `json:"failed_num"      orm:"failed_num"      description:"失败数量"`                              // 失败数量
	SuccessedNum  int         `json:"successed_num"   orm:"successed_num"   description:"成功数量"`                              // 成功数量
	TotalNum      int         `json:"total_num"       orm:"total_num"       description:"总的私信数量"`                            // 总的私信数量
	Remark        string      `json:"remark"          orm:"remark"          description:"备注"`                                // 备注
}
