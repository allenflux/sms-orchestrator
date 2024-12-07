// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskLinkDeviceGroupDao is the data access object for table task_link_device_group.
type TaskLinkDeviceGroupDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns TaskLinkDeviceGroupColumns // columns contains all the column names of Table for convenient usage.
}

// TaskLinkDeviceGroupColumns defines and stores column names for table task_link_device_group.
type TaskLinkDeviceGroupColumns struct {
	Id        string //
	TaskId    string // 任务id
	GroupId   string // 设备分组id
	CreatedAt string // 创建时间
	DeletedAt string // 删除时间
}

// taskLinkDeviceGroupColumns holds the columns for table task_link_device_group.
var taskLinkDeviceGroupColumns = TaskLinkDeviceGroupColumns{
	Id:        "id",
	TaskId:    "task_id",
	GroupId:   "group_id",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewTaskLinkDeviceGroupDao creates and returns a new DAO object for table data access.
func NewTaskLinkDeviceGroupDao() *TaskLinkDeviceGroupDao {
	return &TaskLinkDeviceGroupDao{
		group:   "default",
		table:   "task_link_device_group",
		columns: taskLinkDeviceGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskLinkDeviceGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskLinkDeviceGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskLinkDeviceGroupDao) Columns() TaskLinkDeviceGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskLinkDeviceGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskLinkDeviceGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskLinkDeviceGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
