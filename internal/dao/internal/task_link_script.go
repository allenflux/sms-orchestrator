// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskLinkScriptDao is the data access object for table task_link_script.
type TaskLinkScriptDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TaskLinkScriptColumns // columns contains all the column names of Table for convenient usage.
}

// TaskLinkScriptColumns defines and stores column names for table task_link_script.
type TaskLinkScriptColumns struct {
	Id        string //
	TaskId    string // 任务id
	ScriptId  string // 脚本id
	CreatedAt string // 创建时间
	DeletedAt string // 删除时间
}

// taskLinkScriptColumns holds the columns for table task_link_script.
var taskLinkScriptColumns = TaskLinkScriptColumns{
	Id:        "id",
	TaskId:    "task_id",
	ScriptId:  "script_id",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewTaskLinkScriptDao creates and returns a new DAO object for table data access.
func NewTaskLinkScriptDao() *TaskLinkScriptDao {
	return &TaskLinkScriptDao{
		group:   "default",
		table:   "task_link_script",
		columns: taskLinkScriptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskLinkScriptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskLinkScriptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskLinkScriptDao) Columns() TaskLinkScriptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskLinkScriptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskLinkScriptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskLinkScriptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
