// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FundingCardDao is the data access object for table funding_card.
type FundingCardDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns FundingCardColumns // columns contains all the column names of Table for convenient usage.
}

// FundingCardColumns defines and stores column names for table funding_card.
type FundingCardColumns struct {
	Id            string //
	FundingCardNo string // 资金卡号
	BankUid       string // 银行uid
	Balance       string // 余额（HKD）
	CreatedAt     string // 创建时间
	UpdatedAt     string // 修改时间
	DeletedAt     string // 删除时间
}

// fundingCardColumns holds the columns for table funding_card.
var fundingCardColumns = FundingCardColumns{
	Id:            "id",
	FundingCardNo: "funding_card_no",
	BankUid:       "bank_uid",
	Balance:       "balance",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewFundingCardDao creates and returns a new DAO object for table data access.
func NewFundingCardDao() *FundingCardDao {
	return &FundingCardDao{
		group:   "default",
		table:   "funding_card",
		columns: fundingCardColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FundingCardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FundingCardDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FundingCardDao) Columns() FundingCardColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FundingCardDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FundingCardDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FundingCardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
