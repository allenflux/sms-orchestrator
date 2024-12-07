// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// InstallTaskItem is the golang structure for table install_task_item.
type InstallTaskItem struct {
	Id            uint        `json:"id"              orm:"id"              description:""`                                  //
	InstallTaskId int         `json:"install_task_id" orm:"install_task_id" description:"安装任务id"`                            // 安装任务id
	DeviceId      int         `json:"device_id"       orm:"device_id"       description:"设备id"`                              // 设备id
	ScriptId      int         `json:"script_id"       orm:"script_id"       description:"脚本id"`                              // 脚本id
	State         int         `json:"state"           orm:"state"           description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`                              // 创建时间
}
