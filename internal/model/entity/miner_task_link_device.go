// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerTaskLinkDevice is the golang structure for table miner_task_link_device.
type MinerTaskLinkDevice struct {
	Id          int         `json:"id"            orm:"id"            description:""`                                          //
	MinerTaskId int         `json:"miner_task_id" orm:"miner_task_id" description:"矿机任务id"`                                    // 矿机任务id
	Platform    string      `json:"platform"      orm:"platform"      description:"平台"`                                        // 平台
	DeviceId    int         `json:"device_id"     orm:"device_id"     description:"设备id"`                                      // 设备id
	DeviceDevId string      `json:"device_dev_id" orm:"device_dev_id" description:"设备本机id"`                                    // 设备本机id
	DeviceName  string      `json:"device_name"   orm:"device_name"   description:"设备名称"`                                      // 设备名称
	RunTime     int         `json:"run_time"      orm:"run_time"      description:"运行时间（分钟）"`                                  // 运行时间（分钟）
	Viewed      int         `json:"viewed"        orm:"viewed"        description:"已刷视频"`                                      // 已刷视频
	Earnings    float64     `json:"earnings"      orm:"earnings"      description:"收益"`                                        // 收益
	Withdrawal  float64     `json:"withdrawal"    orm:"withdrawal"    description:"提现"`                                        // 提现
	Balance     float64     `json:"balance"       orm:"balance"       description:"余额"`                                        // 余额
	Status      int         `json:"status"        orm:"status"        description:"任务状态1待执行2执行中3客户取消/已取消/已终止4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消/已取消/已终止4客服取消5已暂停6异常7已完成
	ExecAt      *gtime.Time `json:"exec_at"       orm:"exec_at"       description:"执行时间"`                                      // 执行时间
	CompletedAt *gtime.Time `json:"completed_at"  orm:"completed_at"  description:"完成时间"`                                      // 完成时间
	CreatedAt   *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`                                      // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"    orm:"updated_at"    description:"修改时间"`                                      // 修改时间
	DeletedAt   *gtime.Time `json:"deleted_at"    orm:"deleted_at"    description:"删除时间"`                                      // 删除时间
	CreatedBy   int         `json:"created_by"    orm:"created_by"    description:"创建者id"`                                     // 创建者id
	Remark      string      `json:"remark"        orm:"remark"        description:"备注"`                                        // 备注
	JdGroupName string      `json:"jd_group_name" orm:"jd_group_name" description:"金盾社区名称"`                                    // 金盾社区名称
}
