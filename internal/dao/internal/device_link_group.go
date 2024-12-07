// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceLinkGroupDao is the data access object for table device_link_group.
type DeviceLinkGroupDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns DeviceLinkGroupColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceLinkGroupColumns defines and stores column names for table device_link_group.
type DeviceLinkGroupColumns struct {
	Id        string //
	DeviceId  string // 设备ID
	GroupId   string // 设备分组ID
	CreatedAt string // 创建时间
	ClientId  string // 客户ID
}

// deviceLinkGroupColumns holds the columns for table device_link_group.
var deviceLinkGroupColumns = DeviceLinkGroupColumns{
	Id:        "id",
	DeviceId:  "device_id",
	GroupId:   "group_id",
	CreatedAt: "created_at",
	ClientId:  "client_id",
}

// NewDeviceLinkGroupDao creates and returns a new DAO object for table data access.
func NewDeviceLinkGroupDao() *DeviceLinkGroupDao {
	return &DeviceLinkGroupDao{
		group:   "default",
		table:   "device_link_group",
		columns: deviceLinkGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceLinkGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceLinkGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceLinkGroupDao) Columns() DeviceLinkGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceLinkGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceLinkGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceLinkGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
