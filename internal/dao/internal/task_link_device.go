// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskLinkDeviceDao is the data access object for table task_link_device.
type TaskLinkDeviceDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TaskLinkDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// TaskLinkDeviceColumns defines and stores column names for table task_link_device.
type TaskLinkDeviceColumns struct {
	Id         string //
	TaskId     string // 任务ID
	DeviceId   string // 任务ID
	Num        string // 执行次数
	Status     string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CompleteAt string // 完成时间
	CreatedAt  string //
	UpdatedAt  string //
}

// taskLinkDeviceColumns holds the columns for table task_link_device.
var taskLinkDeviceColumns = TaskLinkDeviceColumns{
	Id:         "id",
	TaskId:     "task_id",
	DeviceId:   "device_id",
	Num:        "num",
	Status:     "status",
	CompleteAt: "complete_at",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewTaskLinkDeviceDao creates and returns a new DAO object for table data access.
func NewTaskLinkDeviceDao() *TaskLinkDeviceDao {
	return &TaskLinkDeviceDao{
		group:   "default",
		table:   "task_link_device",
		columns: taskLinkDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskLinkDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskLinkDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskLinkDeviceDao) Columns() TaskLinkDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskLinkDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskLinkDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskLinkDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
