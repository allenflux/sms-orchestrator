// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskLinkScript is the golang structure for table task_link_script.
type TaskLinkScript struct {
	Id        uint        `json:"id"         orm:"id"         description:""`     //
	TaskId    int         `json:"task_id"    orm:"task_id"    description:"任务id"` // 任务id
	ScriptId  int         `json:"script_id"  orm:"script_id"  description:"脚本id"` // 脚本id
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"` // 删除时间
}
