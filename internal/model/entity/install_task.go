// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// InstallTask is the golang structure for table install_task.
type InstallTask struct {
	Id        uint        `json:"id"         orm:"id"         description:""`                                  //
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`                                // 备注
	State     int         `json:"state"      orm:"state"      description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CreatedBy int         `json:"created_by" orm:"created_by" description:"创建者id（企业用户id）"`                     // 创建者id（企业用户id）
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`                              // 创建时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`                              // 删除时间
}
