// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TpContentDao is the data access object for table tp_content.
type TpContentDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TpContentColumns // columns contains all the column names of Table for convenient usage.
}

// TpContentColumns defines and stores column names for table tp_content.
type TpContentColumns struct {
	Id        string //
	Str       string //
	Remark    string //
	CreatedAt string // 创建时间
}

// tpContentColumns holds the columns for table tp_content.
var tpContentColumns = TpContentColumns{
	Id:        "id",
	Str:       "str",
	Remark:    "remark",
	CreatedAt: "created_at",
}

// NewTpContentDao creates and returns a new DAO object for table data access.
func NewTpContentDao() *TpContentDao {
	return &TpContentDao{
		group:   "default",
		table:   "tp_content",
		columns: tpContentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TpContentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TpContentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TpContentDao) Columns() TpContentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TpContentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TpContentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TpContentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
