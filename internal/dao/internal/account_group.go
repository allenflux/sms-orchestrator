// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountGroupDao is the data access object for table account_group.
type AccountGroupDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns AccountGroupColumns // columns contains all the column names of Table for convenient usage.
}

// AccountGroupColumns defines and stores column names for table account_group.
type AccountGroupColumns struct {
	Id           string //
	Name         string // 账号分组
	CreatedAt    string // 创建时间
	UpdatedAt    string // 修改时间
	ClientId     string // 客户ID
	PlatformType string // 平台类型 1-facebook
}

// accountGroupColumns holds the columns for table account_group.
var accountGroupColumns = AccountGroupColumns{
	Id:           "id",
	Name:         "name",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ClientId:     "client_id",
	PlatformType: "platform_type",
}

// NewAccountGroupDao creates and returns a new DAO object for table data access.
func NewAccountGroupDao() *AccountGroupDao {
	return &AccountGroupDao{
		group:   "default",
		table:   "account_group",
		columns: accountGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AccountGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccountGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AccountGroupDao) Columns() AccountGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccountGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccountGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccountGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
