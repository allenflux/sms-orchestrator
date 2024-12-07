// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpecialOperateDao is the data access object for table special_operate.
type SpecialOperateDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SpecialOperateColumns // columns contains all the column names of Table for convenient usage.
}

// SpecialOperateColumns defines and stores column names for table special_operate.
type SpecialOperateColumns struct {
	Id         string //
	Type       string // 1-绑定 2-解绑 3-延期
	CreatedBy  string // 操作人
	Days       string // 延期操作前的租赁天数
	LeaseEndAt string // 延期前的租赁到期时间
	CreatedAt  string //
	DeletedAt  string //
}

// specialOperateColumns holds the columns for table special_operate.
var specialOperateColumns = SpecialOperateColumns{
	Id:         "id",
	Type:       "type",
	CreatedBy:  "created_by",
	Days:       "days",
	LeaseEndAt: "lease_end_at",
	CreatedAt:  "created_at",
	DeletedAt:  "deleted_at",
}

// NewSpecialOperateDao creates and returns a new DAO object for table data access.
func NewSpecialOperateDao() *SpecialOperateDao {
	return &SpecialOperateDao{
		group:   "default",
		table:   "special_operate",
		columns: specialOperateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SpecialOperateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SpecialOperateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SpecialOperateDao) Columns() SpecialOperateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SpecialOperateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SpecialOperateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SpecialOperateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
