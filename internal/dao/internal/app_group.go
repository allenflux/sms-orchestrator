// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppGroupDao is the data access object for table app_group.
type AppGroupDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AppGroupColumns // columns contains all the column names of Table for convenient usage.
}

// AppGroupColumns defines and stores column names for table app_group.
type AppGroupColumns struct {
	Id        string //
	Name      string // 分组名称
	CreatedAt string //
	UpdatedAt string //
	ClientId  string // 客户ID
	RoleId    string // 角色ID
}

// appGroupColumns holds the columns for table app_group.
var appGroupColumns = AppGroupColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	ClientId:  "client_id",
	RoleId:    "role_id",
}

// NewAppGroupDao creates and returns a new DAO object for table data access.
func NewAppGroupDao() *AppGroupDao {
	return &AppGroupDao{
		group:   "default",
		table:   "app_group",
		columns: appGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppGroupDao) Columns() AppGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
