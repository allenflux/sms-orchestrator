// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerTask is the golang structure for table miner_task.
type MinerTask struct {
	Id              int         `json:"id"                 orm:"id"                 description:""`                                          //
	MinerCronTaskId int         `json:"miner_cron_task_id" orm:"miner_cron_task_id" description:"矿机定时任务id"`                                  // 矿机定时任务id
	Name            string      `json:"name"               orm:"name"               description:"任务名称"`                                      // 任务名称
	Platform        string      `json:"platform"           orm:"platform"           description:"任务平台"`                                      // 任务平台
	RunTime         int         `json:"run_time"           orm:"run_time"           description:"运行时间（分钟）"`                                  // 运行时间（分钟）
	DeviceNum       int         `json:"device_num"         orm:"device_num"         description:"设备数量"`                                      // 设备数量
	CreatedAt       *gtime.Time `json:"created_at"         orm:"created_at"         description:"创建时间"`                                      // 创建时间
	ExecAt          *gtime.Time `json:"exec_at"            orm:"exec_at"            description:"执行时间"`                                      // 执行时间
	UpdatedAt       *gtime.Time `json:"updated_at"         orm:"updated_at"         description:"修改时间"`                                      // 修改时间
	DeletedAt       *gtime.Time `json:"deleted_at"         orm:"deleted_at"         description:"删除时间"`                                      // 删除时间
	CreatedBy       int         `json:"created_by"         orm:"created_by"         description:"创建者id"`                                     // 创建者id
	CompletedAt     *gtime.Time `json:"completed_at"       orm:"completed_at"       description:"完成时间"`                                      // 完成时间
	Remark          string      `json:"remark"             orm:"remark"             description:"备注"`                                        // 备注
	Status          int         `json:"status"             orm:"status"             description:"任务状态1待执行2执行中3客户取消/已取消/已终止4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消/已取消/已终止4客服取消5已暂停6异常7已完成
	ExecMode        int         `json:"exec_mode"          orm:"exec_mode"          description:"任务模式1单次2定时"`                                // 任务模式1单次2定时
	HasShow         int         `json:"has_show"           orm:"has_show"           description:"1-展示 2-不展示"`                                // 1-展示 2-不展示
}
