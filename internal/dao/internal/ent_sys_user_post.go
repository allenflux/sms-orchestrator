// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntSysUserPostDao is the data access object for table ent_sys_user_post.
type EntSysUserPostDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns EntSysUserPostColumns // columns contains all the column names of Table for convenient usage.
}

// EntSysUserPostColumns defines and stores column names for table ent_sys_user_post.
type EntSysUserPostColumns struct {
	UserId string // 用户ID
	PostId string // 岗位ID
}

// entSysUserPostColumns holds the columns for table ent_sys_user_post.
var entSysUserPostColumns = EntSysUserPostColumns{
	UserId: "user_id",
	PostId: "post_id",
}

// NewEntSysUserPostDao creates and returns a new DAO object for table data access.
func NewEntSysUserPostDao() *EntSysUserPostDao {
	return &EntSysUserPostDao{
		group:   "default",
		table:   "ent_sys_user_post",
		columns: entSysUserPostColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EntSysUserPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EntSysUserPostDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EntSysUserPostDao) Columns() EntSysUserPostColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EntSysUserPostDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EntSysUserPostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EntSysUserPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
