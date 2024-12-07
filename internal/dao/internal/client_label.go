// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientLabelDao is the data access object for table client_label.
type ClientLabelDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ClientLabelColumns // columns contains all the column names of Table for convenient usage.
}

// ClientLabelColumns defines and stores column names for table client_label.
type ClientLabelColumns struct {
	Id        string //
	Name      string // 名称
	DeletedAt string // 删除时间
}

// clientLabelColumns holds the columns for table client_label.
var clientLabelColumns = ClientLabelColumns{
	Id:        "id",
	Name:      "name",
	DeletedAt: "deleted_at",
}

// NewClientLabelDao creates and returns a new DAO object for table data access.
func NewClientLabelDao() *ClientLabelDao {
	return &ClientLabelDao{
		group:   "default",
		table:   "client_label",
		columns: clientLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientLabelDao) Columns() ClientLabelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
