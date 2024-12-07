// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CronsDao is the data access object for table crons.
type CronsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns CronsColumns // columns contains all the column names of Table for convenient usage.
}

// CronsColumns defines and stores column names for table crons.
type CronsColumns struct {
	Id          string //
	Type        string // 1-解绑过期设备
	CreatedAt   string //
	ExecTime    string // 耗时
	CreatedDate string // 执行日期
}

// cronsColumns holds the columns for table crons.
var cronsColumns = CronsColumns{
	Id:          "id",
	Type:        "type",
	CreatedAt:   "created_at",
	ExecTime:    "exec_time",
	CreatedDate: "created_date",
}

// NewCronsDao creates and returns a new DAO object for table data access.
func NewCronsDao() *CronsDao {
	return &CronsDao{
		group:   "default",
		table:   "crons",
		columns: cronsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CronsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CronsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CronsDao) Columns() CronsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CronsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CronsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CronsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
