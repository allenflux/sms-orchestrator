// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppLabelDao is the data access object for table app_label.
type AppLabelDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AppLabelColumns // columns contains all the column names of Table for convenient usage.
}

// AppLabelColumns defines and stores column names for table app_label.
type AppLabelColumns struct {
	Id        string //
	Name      string // 标签名称
	CreatedAt string // 创建时间
	DeletedAt string // 删除时间
}

// appLabelColumns holds the columns for table app_label.
var appLabelColumns = AppLabelColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewAppLabelDao creates and returns a new DAO object for table data access.
func NewAppLabelDao() *AppLabelDao {
	return &AppLabelDao{
		group:   "default",
		table:   "app_label",
		columns: appLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppLabelDao) Columns() AppLabelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
