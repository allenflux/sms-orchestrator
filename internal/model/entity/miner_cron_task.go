// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerCronTask is the golang structure for table miner_cron_task.
type MinerCronTask struct {
	Id          int         `json:"id"            orm:"id"            description:""`                     //
	Platform    string      `json:"platform"      orm:"platform"      description:"任务平台"`                 // 任务平台
	RunTime     int         `json:"run_time"      orm:"run_time"      description:"运行时间（分钟）"`             // 运行时间（分钟）
	DeviceNum   int         `json:"device_num"    orm:"device_num"    description:"设备数量"`                 // 设备数量
	CreatedAt   *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`                 // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"    orm:"updated_at"    description:"修改时间"`                 // 修改时间
	DeletedAt   *gtime.Time `json:"deleted_at"    orm:"deleted_at"    description:"删除时间"`                 // 删除时间
	CreatedBy   int         `json:"created_by"    orm:"created_by"    description:"创建者id"`                // 创建者id
	Remark      string      `json:"remark"        orm:"remark"        description:"备注"`                   // 备注
	CronMode    int         `json:"cron_mode"     orm:"cron_mode"     description:"定时模式1每日定时2周期定时3循环"`    // 定时模式1每日定时2周期定时3循环
	CronStartAt *gtime.Time `json:"cron_start_at" orm:"cron_start_at" description:"定时计划开始时间"`             // 定时计划开始时间
	CronEndAt   *gtime.Time `json:"cron_end_at"   orm:"cron_end_at"   description:"定时计划结束时间"`             // 定时计划结束时间
	CronExecNum int         `json:"cron_exec_num" orm:"cron_exec_num" description:"定时任务执行次数"`             // 定时任务执行次数
	CronExecAt  *gtime.Time `json:"cron_exec_at"  orm:"cron_exec_at"  description:"定时计划执行时间"`             // 定时计划执行时间
	CronGapTime int         `json:"cron_gap_time" orm:"cron_gap_time" description:"定时计划间隔时间，单位分钟"`        // 定时计划间隔时间，单位分钟
	CronStatus  int         `json:"cron_status"   orm:"cron_status"   description:"开启状态1开启2关闭"`           // 开启状态1开启2关闭
	HadExecNum  int         `json:"had_exec_num"  orm:"had_exec_num"  description:"已执行次数"`                // 已执行次数
	Name        string      `json:"name"          orm:"name"          description:"任务名称"`                 // 任务名称
	HasShow     int         `json:"has_show"      orm:"has_show"      description:"1-全部显示 2-企业显示 3-群控显示"` // 1-全部显示 2-企业显示 3-群控显示
}
