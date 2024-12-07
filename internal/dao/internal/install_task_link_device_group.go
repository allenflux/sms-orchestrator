// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InstallTaskLinkDeviceGroupDao is the data access object for table install_task_link_device_group.
type InstallTaskLinkDeviceGroupDao struct {
	table   string                            // table is the underlying table name of the DAO.
	group   string                            // group is the database configuration group name of current DAO.
	columns InstallTaskLinkDeviceGroupColumns // columns contains all the column names of Table for convenient usage.
}

// InstallTaskLinkDeviceGroupColumns defines and stores column names for table install_task_link_device_group.
type InstallTaskLinkDeviceGroupColumns struct {
	Id            string //
	InstallTaskId string // 安装任务id
	DeviceGroupId string // 设备组id
}

// installTaskLinkDeviceGroupColumns holds the columns for table install_task_link_device_group.
var installTaskLinkDeviceGroupColumns = InstallTaskLinkDeviceGroupColumns{
	Id:            "id",
	InstallTaskId: "install_task_id",
	DeviceGroupId: "device_group_id",
}

// NewInstallTaskLinkDeviceGroupDao creates and returns a new DAO object for table data access.
func NewInstallTaskLinkDeviceGroupDao() *InstallTaskLinkDeviceGroupDao {
	return &InstallTaskLinkDeviceGroupDao{
		group:   "default",
		table:   "install_task_link_device_group",
		columns: installTaskLinkDeviceGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InstallTaskLinkDeviceGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InstallTaskLinkDeviceGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InstallTaskLinkDeviceGroupDao) Columns() InstallTaskLinkDeviceGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InstallTaskLinkDeviceGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InstallTaskLinkDeviceGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InstallTaskLinkDeviceGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
