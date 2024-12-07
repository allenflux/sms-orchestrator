// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerSessionDao is the data access object for table miner_session.
type MinerSessionDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns MinerSessionColumns // columns contains all the column names of Table for convenient usage.
}

// MinerSessionColumns defines and stores column names for table miner_session.
type MinerSessionColumns struct {
	Id       string //
	DevId    string // 设备dev id
	LoginAt  string // 登录时间
	LogoffAt string // 登出时间
}

// minerSessionColumns holds the columns for table miner_session.
var minerSessionColumns = MinerSessionColumns{
	Id:       "id",
	DevId:    "dev_id",
	LoginAt:  "login_at",
	LogoffAt: "logoff_at",
}

// NewMinerSessionDao creates and returns a new DAO object for table data access.
func NewMinerSessionDao() *MinerSessionDao {
	return &MinerSessionDao{
		group:   "default",
		table:   "miner_session",
		columns: minerSessionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerSessionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerSessionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerSessionDao) Columns() MinerSessionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerSessionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerSessionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerSessionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
