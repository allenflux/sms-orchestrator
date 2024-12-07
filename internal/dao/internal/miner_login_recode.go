// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerLoginRecodeDao is the data access object for table miner_login_recode.
type MinerLoginRecodeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns MinerLoginRecodeColumns // columns contains all the column names of Table for convenient usage.
}

// MinerLoginRecodeColumns defines and stores column names for table miner_login_recode.
type MinerLoginRecodeColumns struct {
	Id        string //
	DeviceId  string // 设备主键ID
	Platform  string // 平台
	Status    string // 1 未登录 2 已登录
	ClientId  string // 客户ID
	Remark    string // 备注
	CreatedAt string //
	DevId     string // 设备ID
}

// minerLoginRecodeColumns holds the columns for table miner_login_recode.
var minerLoginRecodeColumns = MinerLoginRecodeColumns{
	Id:        "id",
	DeviceId:  "device_id",
	Platform:  "platform",
	Status:    "status",
	ClientId:  "client_id",
	Remark:    "remark",
	CreatedAt: "created_at",
	DevId:     "dev_id",
}

// NewMinerLoginRecodeDao creates and returns a new DAO object for table data access.
func NewMinerLoginRecodeDao() *MinerLoginRecodeDao {
	return &MinerLoginRecodeDao{
		group:   "default",
		table:   "miner_login_recode",
		columns: minerLoginRecodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerLoginRecodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerLoginRecodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerLoginRecodeDao) Columns() MinerLoginRecodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerLoginRecodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerLoginRecodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerLoginRecodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
