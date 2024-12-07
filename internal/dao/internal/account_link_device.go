// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountLinkDeviceDao is the data access object for table account_link_device.
type AccountLinkDeviceDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns AccountLinkDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// AccountLinkDeviceColumns defines and stores column names for table account_link_device.
type AccountLinkDeviceColumns struct {
	Id           string //
	DevId        string // 设备ID
	AccountId    string // 账号ID
	DeviceId     string // 设备表主键ID
	CreatedAt    string //
	UpdatedAt    string //
	PlatformType string // 平台： 1-facebook
}

// accountLinkDeviceColumns holds the columns for table account_link_device.
var accountLinkDeviceColumns = AccountLinkDeviceColumns{
	Id:           "id",
	DevId:        "dev_id",
	AccountId:    "account_id",
	DeviceId:     "device_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	PlatformType: "platform_type",
}

// NewAccountLinkDeviceDao creates and returns a new DAO object for table data access.
func NewAccountLinkDeviceDao() *AccountLinkDeviceDao {
	return &AccountLinkDeviceDao{
		group:   "default",
		table:   "account_link_device",
		columns: accountLinkDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AccountLinkDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccountLinkDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AccountLinkDeviceDao) Columns() AccountLinkDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccountLinkDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccountLinkDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccountLinkDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
