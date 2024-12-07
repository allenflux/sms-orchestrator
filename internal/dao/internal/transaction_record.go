// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TransactionRecordDao is the data access object for table transaction_record.
type TransactionRecordDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns TransactionRecordColumns // columns contains all the column names of Table for convenient usage.
}

// TransactionRecordColumns defines and stores column names for table transaction_record.
type TransactionRecordColumns struct {
	Id               string //
	Xid              string // 用户xid
	RefId            string //
	CardAccountId    string //
	CardLast4        string //
	CardEmbossedName string //
	CreatedAt        string // 创建时间
	Status           string // 状态 pending 待处理  posted 已完成 declined 已拒绝  vodi 已取消
	PostedAt         string //
	Intent           string // 交易目的/方式，charge收费 refund退款 topup充值 repay偿还 cashback现金返回 interest利息 fee费用 other其他
	Amount           string //
	EntryType        string //
	Currency         string //
	Description      string // 响应消息（例如：响应被接受）
}

// transactionRecordColumns holds the columns for table transaction_record.
var transactionRecordColumns = TransactionRecordColumns{
	Id:               "id",
	Xid:              "xid",
	RefId:            "ref_id",
	CardAccountId:    "card_account_id",
	CardLast4:        "card_last_4",
	CardEmbossedName: "card_embossed_name",
	CreatedAt:        "created_at",
	Status:           "status",
	PostedAt:         "posted_at",
	Intent:           "intent",
	Amount:           "amount",
	EntryType:        "entry_type",
	Currency:         "currency",
	Description:      "description",
}

// NewTransactionRecordDao creates and returns a new DAO object for table data access.
func NewTransactionRecordDao() *TransactionRecordDao {
	return &TransactionRecordDao{
		group:   "default",
		table:   "transaction_record",
		columns: transactionRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TransactionRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TransactionRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TransactionRecordDao) Columns() TransactionRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TransactionRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TransactionRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TransactionRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
