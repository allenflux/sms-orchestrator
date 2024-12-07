// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntPlatformDao is the data access object for table ent_platform.
type EntPlatformDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns EntPlatformColumns // columns contains all the column names of Table for convenient usage.
}

// EntPlatformColumns defines and stores column names for table ent_platform.
type EntPlatformColumns struct {
	Id           string //
	ParentId     string //
	CatalogLevel string // 目录级别
	Name         string // 名称
	Nickname     string // 别名
	ComponentUrl string // 组件路径
}

// entPlatformColumns holds the columns for table ent_platform.
var entPlatformColumns = EntPlatformColumns{
	Id:           "id",
	ParentId:     "parent_id",
	CatalogLevel: "catalog_level",
	Name:         "name",
	Nickname:     "nickname",
	ComponentUrl: "component_url",
}

// NewEntPlatformDao creates and returns a new DAO object for table data access.
func NewEntPlatformDao() *EntPlatformDao {
	return &EntPlatformDao{
		group:   "default",
		table:   "ent_platform",
		columns: entPlatformColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EntPlatformDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EntPlatformDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EntPlatformDao) Columns() EntPlatformColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EntPlatformDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EntPlatformDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EntPlatformDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
