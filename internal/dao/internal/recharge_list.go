// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RechargeListDao is the data access object for table recharge_list.
type RechargeListDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns RechargeListColumns // columns contains all the column names of Table for convenient usage.
}

// RechargeListColumns defines and stores column names for table recharge_list.
type RechargeListColumns struct {
	Id            string //
	FundCardNo    string // 资金卡号
	CardNo        string // 对方卡号
	Amount        string // 金额（hkd）
	Name          string // 对方姓名
	ApplyListId   string // 申请列表id
	AdminUsername string // 操作账号
	CreatedAt     string // 充值时间
	UpdatedAt     string //
	DeletedAt     string //
}

// rechargeListColumns holds the columns for table recharge_list.
var rechargeListColumns = RechargeListColumns{
	Id:            "id",
	FundCardNo:    "fund_card_no",
	CardNo:        "card_no",
	Amount:        "amount",
	Name:          "name",
	ApplyListId:   "apply_list_id",
	AdminUsername: "admin_username",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewRechargeListDao creates and returns a new DAO object for table data access.
func NewRechargeListDao() *RechargeListDao {
	return &RechargeListDao{
		group:   "default",
		table:   "recharge_list",
		columns: rechargeListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RechargeListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RechargeListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RechargeListDao) Columns() RechargeListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RechargeListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RechargeListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RechargeListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
