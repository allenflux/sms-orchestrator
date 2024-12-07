// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ScriptOp is the golang structure for table script_op.
type ScriptOp struct {
	Id              uint        `json:"id"                orm:"id"                description:""`         //
	ScriptId        int         `json:"script_id"         orm:"script_id"         description:"脚本表id"`    // 脚本表id
	Op              int         `json:"op"                orm:"op"                description:"操作1新增2编辑"` // 操作1新增2编辑
	Column          string      `json:"column"            orm:"column"            description:"列"`        // 列
	ColumnComment   string      `json:"column_comment"    orm:"column_comment"    description:"列中文名"`     // 列中文名
	OldValue        string      `json:"old_value"         orm:"old_value"         description:"原先的值"`     // 原先的值
	OldValueComment string      `json:"old_value_comment" orm:"old_value_comment" description:"原先的值中文名"`  // 原先的值中文名
	Value           string      `json:"value"             orm:"value"             description:"值"`        // 值
	ValueComment    string      `json:"value_comment"     orm:"value_comment"     description:"值中文名"`     // 值中文名
	CreatedAt       *gtime.Time `json:"created_at"        orm:"created_at"        description:"操作时间"`     // 操作时间
	UpdatedBy       int         `json:"updated_by"        orm:"updated_by"        description:"操作人"`      // 操作人
}
