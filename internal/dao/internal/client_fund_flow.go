// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientFundFlowDao is the data access object for table client_fund_flow.
type ClientFundFlowDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns ClientFundFlowColumns // columns contains all the column names of Table for convenient usage.
}

// ClientFundFlowColumns defines and stores column names for table client_fund_flow.
type ClientFundFlowColumns struct {
	Id              string //
	FundType        string // 1-租赁 2-续费 3-退款 4-余额 5-授信  6-产生订单  7-取消订单 8-驳回订单 9-绑定 10-延期
	OrderId         string // 订单ID
	ClientId        string // 客户ID
	BalanceBefore   string // 操作前余额
	BalanceAfter    string // 操作后余额
	CreditBefore    string // 操作前授信金额
	CreditAfter     string // 操作后授信金额
	AdvanceBefore   string // 操作前预扣金额
	AdvanceAfter    string // 操作后预扣金额
	UseCreditBefore string // 操作前已使用预扣金额
	UseCreditAfter  string // 操作后已使用预扣金额
	Amount          string // 操作数量
	CreatedBy       string // 操作人ID
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 删除时间
	TransactionType string // 交易类型1-存入 2-支出
	Remark          string // 备注
	OrderNo         string // 订单号
}

// clientFundFlowColumns holds the columns for table client_fund_flow.
var clientFundFlowColumns = ClientFundFlowColumns{
	Id:              "id",
	FundType:        "fund_type",
	OrderId:         "order_id",
	ClientId:        "client_id",
	BalanceBefore:   "balance_before",
	BalanceAfter:    "balance_after",
	CreditBefore:    "credit_before",
	CreditAfter:     "credit_after",
	AdvanceBefore:   "advance_before",
	AdvanceAfter:    "advance_after",
	UseCreditBefore: "use_credit_before",
	UseCreditAfter:  "use_credit_after",
	Amount:          "amount",
	CreatedBy:       "created_by",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	TransactionType: "transaction_type",
	Remark:          "remark",
	OrderNo:         "order_no",
}

// NewClientFundFlowDao creates and returns a new DAO object for table data access.
func NewClientFundFlowDao() *ClientFundFlowDao {
	return &ClientFundFlowDao{
		group:   "default",
		table:   "client_fund_flow",
		columns: clientFundFlowColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientFundFlowDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientFundFlowDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientFundFlowDao) Columns() ClientFundFlowColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientFundFlowDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientFundFlowDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientFundFlowDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
