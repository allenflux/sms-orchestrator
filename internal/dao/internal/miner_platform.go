// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerPlatformDao is the data access object for table miner_platform.
type MinerPlatformDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns MinerPlatformColumns // columns contains all the column names of Table for convenient usage.
}

// MinerPlatformColumns defines and stores column names for table miner_platform.
type MinerPlatformColumns struct {
	Id       string //
	Platform string // 支持的平台
}

// minerPlatformColumns holds the columns for table miner_platform.
var minerPlatformColumns = MinerPlatformColumns{
	Id:       "id",
	Platform: "platform",
}

// NewMinerPlatformDao creates and returns a new DAO object for table data access.
func NewMinerPlatformDao() *MinerPlatformDao {
	return &MinerPlatformDao{
		group:   "default",
		table:   "miner_platform",
		columns: minerPlatformColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerPlatformDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerPlatformDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerPlatformDao) Columns() MinerPlatformColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerPlatformDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerPlatformDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerPlatformDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
