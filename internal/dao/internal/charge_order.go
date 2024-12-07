// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChargeOrderDao is the data access object for table charge_order.
type ChargeOrderDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ChargeOrderColumns // columns contains all the column names of Table for convenient usage.
}

// ChargeOrderColumns defines and stores column names for table charge_order.
type ChargeOrderColumns struct {
	Id                 string //
	ClientId           string // 客户id
	ChargeType         string // 加款类型1余额2授信
	ChargeAmount       string // 添加金额
	BalanceAfterCharge string // 加款后余额
	CreditAfterCharge  string // 加款后授信
	ChargeBank         string // 加款银行
	Remark             string // 备注
	CreatedBy          string // 操作人
	CreatedAt          string // 加款时间
	DeletedAt          string // 删除时间
	OrderNo            string // 订单号
}

// chargeOrderColumns holds the columns for table charge_order.
var chargeOrderColumns = ChargeOrderColumns{
	Id:                 "id",
	ClientId:           "client_id",
	ChargeType:         "charge_type",
	ChargeAmount:       "charge_amount",
	BalanceAfterCharge: "balance_after_charge",
	CreditAfterCharge:  "credit_after_charge",
	ChargeBank:         "charge_bank",
	Remark:             "remark",
	CreatedBy:          "created_by",
	CreatedAt:          "created_at",
	DeletedAt:          "deleted_at",
	OrderNo:            "order_no",
}

// NewChargeOrderDao creates and returns a new DAO object for table data access.
func NewChargeOrderDao() *ChargeOrderDao {
	return &ChargeOrderDao{
		group:   "default",
		table:   "charge_order",
		columns: chargeOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChargeOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChargeOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChargeOrderDao) Columns() ChargeOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChargeOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChargeOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChargeOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
