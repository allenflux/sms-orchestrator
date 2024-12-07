// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceBindHistoryDao is the data access object for table device_bind_history.
type DeviceBindHistoryDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns DeviceBindHistoryColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceBindHistoryColumns defines and stores column names for table device_bind_history.
type DeviceBindHistoryColumns struct {
	Id         string //
	DeviceId   string // 设备id
	ClientId   string // 客户id
	LeasedAt   string // 租赁时间
	LeaseEndAt string // 租赁到期时间
	UnbindAt   string // 解绑时间
	Remark     string // 备注
}

// deviceBindHistoryColumns holds the columns for table device_bind_history.
var deviceBindHistoryColumns = DeviceBindHistoryColumns{
	Id:         "id",
	DeviceId:   "device_id",
	ClientId:   "client_id",
	LeasedAt:   "leased_at",
	LeaseEndAt: "lease_end_at",
	UnbindAt:   "unbind_at",
	Remark:     "remark",
}

// NewDeviceBindHistoryDao creates and returns a new DAO object for table data access.
func NewDeviceBindHistoryDao() *DeviceBindHistoryDao {
	return &DeviceBindHistoryDao{
		group:   "default",
		table:   "device_bind_history",
		columns: deviceBindHistoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceBindHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceBindHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceBindHistoryDao) Columns() DeviceBindHistoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceBindHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceBindHistoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceBindHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
