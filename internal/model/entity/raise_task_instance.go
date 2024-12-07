// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseTaskInstance is the golang structure for table raise_task_instance.
type RaiseTaskInstance struct {
	Id          uint        `json:"id"            orm:"id"            description:""`                                  //
	RaiseTaskId int         `json:"raise_task_id" orm:"raise_task_id" description:"养号任务id"`                            // 养号任务id
	CreatedAt   *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间（开始执行时间）"`                      // 创建时间（开始执行时间）
	UpdatedAt   *gtime.Time `json:"updated_at"    orm:"updated_at"    description:"修改时间"`                              // 修改时间
	DeletedAt   *gtime.Time `json:"deleted_at"    orm:"deleted_at"    description:"删除时间"`                              // 删除时间
	ExecAt      *gtime.Time `json:"exec_at"       orm:"exec_at"       description:"执行时间"`                              // 执行时间
	CompletedAt *gtime.Time `json:"completed_at"  orm:"completed_at"  description:"完成时间"`                              // 完成时间
	Status      int         `json:"status"        orm:"status"        description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	Remark      string      `json:"remark"        orm:"remark"        description:"备注"`                                // 备注
	HasEmpty    int         `json:"has_empty"     orm:"has_empty"     description:"1-未清空 2-清空"`                        // 1-未清空 2-清空
	ClientId    int         `json:"client_id"     orm:"client_id"     description:"创建用户id"`                            // 创建用户id
}
