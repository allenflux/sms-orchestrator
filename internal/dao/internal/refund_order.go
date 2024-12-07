// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RefundOrderDao is the data access object for table refund_order.
type RefundOrderDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns RefundOrderColumns // columns contains all the column names of Table for convenient usage.
}

// RefundOrderColumns defines and stores column names for table refund_order.
type RefundOrderColumns struct {
	Id           string //
	RefundAmount string // 退款金额
	ClientId     string // 客户ID
	OrderNo      string // 订单号
	CreatedBy    string // 操作人ID
	CreatedAt    string // 创建时间
}

// refundOrderColumns holds the columns for table refund_order.
var refundOrderColumns = RefundOrderColumns{
	Id:           "id",
	RefundAmount: "refund_amount",
	ClientId:     "client_id",
	OrderNo:      "order_no",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
}

// NewRefundOrderDao creates and returns a new DAO object for table data access.
func NewRefundOrderDao() *RefundOrderDao {
	return &RefundOrderDao{
		group:   "default",
		table:   "refund_order",
		columns: refundOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RefundOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RefundOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RefundOrderDao) Columns() RefundOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RefundOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RefundOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RefundOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
