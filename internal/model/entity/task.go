// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure for table task.
type Task struct {
	Id              int         `json:"id"               orm:"id"               description:""`                                          //
	Name            string      `json:"name"             orm:"name"             description:"任务名称"`                                      // 任务名称
	ClientId        int         `json:"client_id"        orm:"client_id"        description:"客户id"`                                      // 客户id
	ScriptName      string      `json:"script_name"      orm:"script_name"      description:"脚本名称"`                                      // 脚本名称
	DeviceNum       int         `json:"device_num"       orm:"device_num"       description:"设备数量"`                                      // 设备数量
	AppId           int         `json:"app_id"           orm:"app_id"           description:"平台id"`                                      // 平台id
	TaskType        int         `json:"task_type"        orm:"task_type"        description:"任务类型 1养号2分配账号3刷视频 4 批量解绑账号  5-修改账号信息"`      // 任务类型 1养号2分配账号3刷视频 4 批量解绑账号  5-修改账号信息
	Num             int         `json:"num"              orm:"num"              description:"执行次数"`                                      // 执行次数
	ImplementAt     *gtime.Time `json:"implement_at"     orm:"implement_at"     description:"执行时间"`                                      // 执行时间
	ImplementType   int         `json:"implement_type"   orm:"implement_type"   description:"1-立即执行 2-定时执行"`                             // 1-立即执行 2-定时执行
	TimingAt        *gtime.Time `json:"timing_at"        orm:"timing_at"        description:"定时执行时间"`                                    // 定时执行时间
	Remark          string      `json:"remark"           orm:"remark"           description:"备注"`                                        // 备注
	ImplementStatus int         `json:"implement_status" orm:"implement_status" description:"1-待执行 2-执行中 3-客户取消 4-客服取消 5-暂停 6-异常 7-已完成"` // 1-待执行 2-执行中 3-客户取消 4-客服取消 5-暂停 6-异常 7-已完成
	OperationStatus int         `json:"operation_status" orm:"operation_status" description:"1-正常（编辑） 2-取消 3-暂停 4-开始 5-明细"`              // 1-正常（编辑） 2-取消 3-暂停 4-开始 5-明细
	CompleteAt      *gtime.Time `json:"complete_at"      orm:"complete_at"      description:"任务完成时间"`                                    // 任务完成时间
	CancelAt        *gtime.Time `json:"cancel_at"        orm:"cancel_at"        description:"任务取消时间"`                                    // 任务取消时间
	CreatedBy       int         `json:"created_by"       orm:"created_by"       description:"创建者id"`                                     // 创建者id
	CreatedAt       *gtime.Time `json:"created_at"       orm:"created_at"       description:"创建时间"`                                      // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"       orm:"updated_at"       description:"修改时间"`                                      // 修改时间
	PlatformType    int         `json:"platform_type"    orm:"platform_type"    description:"平台类型 1-facebook"`                           // 平台类型 1-facebook
	HasEmpty        int         `json:"has_empty"        orm:"has_empty"        description:"1-未清空 2-清空"`                                // 1-未清空 2-清空
}
