// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectTaskLinkDeviceGroupDao is the data access object for table collect_task_link_device_group.
type CollectTaskLinkDeviceGroupDao struct {
	table   string                            // table is the underlying table name of the DAO.
	group   string                            // group is the database configuration group name of current DAO.
	columns CollectTaskLinkDeviceGroupColumns // columns contains all the column names of Table for convenient usage.
}

// CollectTaskLinkDeviceGroupColumns defines and stores column names for table collect_task_link_device_group.
type CollectTaskLinkDeviceGroupColumns struct {
	Id            string //
	CollectTaskId string // 采集任务id
	DeviceGroupId string // 设备组id
	CreatedAt     string // 创建时间
	DeletedAt     string // 删除时间
}

// collectTaskLinkDeviceGroupColumns holds the columns for table collect_task_link_device_group.
var collectTaskLinkDeviceGroupColumns = CollectTaskLinkDeviceGroupColumns{
	Id:            "id",
	CollectTaskId: "collect_task_id",
	DeviceGroupId: "device_group_id",
	CreatedAt:     "created_at",
	DeletedAt:     "deleted_at",
}

// NewCollectTaskLinkDeviceGroupDao creates and returns a new DAO object for table data access.
func NewCollectTaskLinkDeviceGroupDao() *CollectTaskLinkDeviceGroupDao {
	return &CollectTaskLinkDeviceGroupDao{
		group:   "default",
		table:   "collect_task_link_device_group",
		columns: collectTaskLinkDeviceGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CollectTaskLinkDeviceGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CollectTaskLinkDeviceGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CollectTaskLinkDeviceGroupDao) Columns() CollectTaskLinkDeviceGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CollectTaskLinkDeviceGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CollectTaskLinkDeviceGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CollectTaskLinkDeviceGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
