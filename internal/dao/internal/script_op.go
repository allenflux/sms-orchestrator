// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ScriptOpDao is the data access object for table script_op.
type ScriptOpDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ScriptOpColumns // columns contains all the column names of Table for convenient usage.
}

// ScriptOpColumns defines and stores column names for table script_op.
type ScriptOpColumns struct {
	Id              string //
	ScriptId        string // 脚本表id
	Op              string // 操作1新增2编辑
	Column          string // 列
	ColumnComment   string // 列中文名
	OldValue        string // 原先的值
	OldValueComment string // 原先的值中文名
	Value           string // 值
	ValueComment    string // 值中文名
	CreatedAt       string // 操作时间
	UpdatedBy       string // 操作人
}

// scriptOpColumns holds the columns for table script_op.
var scriptOpColumns = ScriptOpColumns{
	Id:              "id",
	ScriptId:        "script_id",
	Op:              "op",
	Column:          "column",
	ColumnComment:   "column_comment",
	OldValue:        "old_value",
	OldValueComment: "old_value_comment",
	Value:           "value",
	ValueComment:    "value_comment",
	CreatedAt:       "created_at",
	UpdatedBy:       "updated_by",
}

// NewScriptOpDao creates and returns a new DAO object for table data access.
func NewScriptOpDao() *ScriptOpDao {
	return &ScriptOpDao{
		group:   "default",
		table:   "script_op",
		columns: scriptOpColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ScriptOpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ScriptOpDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ScriptOpDao) Columns() ScriptOpColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ScriptOpDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ScriptOpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ScriptOpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
