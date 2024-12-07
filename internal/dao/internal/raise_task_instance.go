// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RaiseTaskInstanceDao is the data access object for table raise_task_instance.
type RaiseTaskInstanceDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns RaiseTaskInstanceColumns // columns contains all the column names of Table for convenient usage.
}

// RaiseTaskInstanceColumns defines and stores column names for table raise_task_instance.
type RaiseTaskInstanceColumns struct {
	Id          string //
	RaiseTaskId string // 养号任务id
	CreatedAt   string // 创建时间（开始执行时间）
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
	ExecAt      string // 执行时间
	CompletedAt string // 完成时间
	Status      string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	Remark      string // 备注
	HasEmpty    string // 1-未清空 2-清空
	ClientId    string // 创建用户id
}

// raiseTaskInstanceColumns holds the columns for table raise_task_instance.
var raiseTaskInstanceColumns = RaiseTaskInstanceColumns{
	Id:          "id",
	RaiseTaskId: "raise_task_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	ExecAt:      "exec_at",
	CompletedAt: "completed_at",
	Status:      "status",
	Remark:      "remark",
	HasEmpty:    "has_empty",
	ClientId:    "client_id",
}

// NewRaiseTaskInstanceDao creates and returns a new DAO object for table data access.
func NewRaiseTaskInstanceDao() *RaiseTaskInstanceDao {
	return &RaiseTaskInstanceDao{
		group:   "default",
		table:   "raise_task_instance",
		columns: raiseTaskInstanceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RaiseTaskInstanceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RaiseTaskInstanceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RaiseTaskInstanceDao) Columns() RaiseTaskInstanceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RaiseTaskInstanceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RaiseTaskInstanceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RaiseTaskInstanceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
