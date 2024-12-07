// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerLoginTaskLinkDevice is the golang structure for table miner_login_task_link_device.
type MinerLoginTaskLinkDevice struct {
	Id               int         `json:"id"                  orm:"id"                  description:""`                                  //
	DevId            string      `json:"dev_id"              orm:"dev_id"              description:"设备ID"`                              // 设备ID
	Status           int         `json:"status"              orm:"status"              description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	MinerLoginTaskId int         `json:"miner_login_task_id" orm:"miner_login_task_id" description:"矿机登录任务ID"`                          // 矿机登录任务ID
	CreatedAt        *gtime.Time `json:"created_at"          orm:"created_at"          description:""`                                  //
	UpdatedAt        *gtime.Time `json:"updated_at"          orm:"updated_at"          description:""`                                  //
	Identity         string      `json:"identity"            orm:"identity"            description:"唯一标识"`                              // 唯一标识
	ExecAt           *gtime.Time `json:"exec_at"             orm:"exec_at"             description:"开始执行时间"`                            // 开始执行时间
	Remark           string      `json:"remark"              orm:"remark"              description:"备注"`                                // 备注
	DeviceId         int         `json:"device_id"           orm:"device_id"           description:"设备主键ID"`                            // 设备主键ID
	Platform         string      `json:"platform"            orm:"platform"            description:"平台"`                                // 平台
	ClientId         int         `json:"client_id"           orm:"client_id"           description:"客户ID"`                              // 客户ID
	ExecEndAt        *gtime.Time `json:"exec_end_at"         orm:"exec_end_at"         description:"执行结束时间"`                            // 执行结束时间
	JdGroupName      string      `json:"jd_group_name"       orm:"jd_group_name"       description:"金盾社区名称"`                            // 金盾社区名称
}
