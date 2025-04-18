// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ScriptLinkAppGroupDao is the data access object for table script_link_app_group.
type ScriptLinkAppGroupDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ScriptLinkAppGroupColumns // columns contains all the column names of Table for convenient usage.
}

// ScriptLinkAppGroupColumns defines and stores column names for table script_link_app_group.
type ScriptLinkAppGroupColumns struct {
	Id         string //
	ScriptId   string // 脚本ID
	AppGroupId string // app分组ID
	CreatedAt  string // 创建时间
}

// scriptLinkAppGroupColumns holds the columns for table script_link_app_group.
var scriptLinkAppGroupColumns = ScriptLinkAppGroupColumns{
	Id:         "id",
	ScriptId:   "script_id",
	AppGroupId: "app_group_id",
	CreatedAt:  "created_at",
}

// NewScriptLinkAppGroupDao creates and returns a new DAO object for table data access.
func NewScriptLinkAppGroupDao() *ScriptLinkAppGroupDao {
	return &ScriptLinkAppGroupDao{
		group:   "default",
		table:   "script_link_app_group",
		columns: scriptLinkAppGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ScriptLinkAppGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ScriptLinkAppGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ScriptLinkAppGroupDao) Columns() ScriptLinkAppGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ScriptLinkAppGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ScriptLinkAppGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ScriptLinkAppGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
