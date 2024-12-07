// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientLinkLabelDao is the data access object for table client_link_label.
type ClientLinkLabelDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns ClientLinkLabelColumns // columns contains all the column names of Table for convenient usage.
}

// ClientLinkLabelColumns defines and stores column names for table client_link_label.
type ClientLinkLabelColumns struct {
	Id            string //
	ClientId      string // 客户id
	ClientLabelId string // 客户标签id
	DeletedAt     string // 删除时间
}

// clientLinkLabelColumns holds the columns for table client_link_label.
var clientLinkLabelColumns = ClientLinkLabelColumns{
	Id:            "id",
	ClientId:      "client_id",
	ClientLabelId: "client_label_id",
	DeletedAt:     "deleted_at",
}

// NewClientLinkLabelDao creates and returns a new DAO object for table data access.
func NewClientLinkLabelDao() *ClientLinkLabelDao {
	return &ClientLinkLabelDao{
		group:   "default",
		table:   "client_link_label",
		columns: clientLinkLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientLinkLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientLinkLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientLinkLabelDao) Columns() ClientLinkLabelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientLinkLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientLinkLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientLinkLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
