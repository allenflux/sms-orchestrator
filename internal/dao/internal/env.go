// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EnvDao is the data access object for table env.
type EnvDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns EnvColumns // columns contains all the column names of Table for convenient usage.
}

// EnvColumns defines and stores column names for table env.
type EnvColumns struct {
	Id        string //
	Name      string // 环境名称
	DeviceId  string // 设备id
	ClientId  string // 客户ID
	Remark    string // 备注
	CreatedAt string // 环境保存时间
	DeletedAt string // 删除时间
}

// envColumns holds the columns for table env.
var envColumns = EnvColumns{
	Id:        "id",
	Name:      "name",
	DeviceId:  "device_id",
	ClientId:  "client_id",
	Remark:    "remark",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewEnvDao creates and returns a new DAO object for table data access.
func NewEnvDao() *EnvDao {
	return &EnvDao{
		group:   "default",
		table:   "env",
		columns: envColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EnvDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EnvDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EnvDao) Columns() EnvColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EnvDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EnvDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EnvDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
