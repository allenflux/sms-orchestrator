// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntSysRoleDeptDao is the data access object for table ent_sys_role_dept.
type EntSysRoleDeptDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns EntSysRoleDeptColumns // columns contains all the column names of Table for convenient usage.
}

// EntSysRoleDeptColumns defines and stores column names for table ent_sys_role_dept.
type EntSysRoleDeptColumns struct {
	RoleId string // 角色ID
	DeptId string // 部门ID
}

// entSysRoleDeptColumns holds the columns for table ent_sys_role_dept.
var entSysRoleDeptColumns = EntSysRoleDeptColumns{
	RoleId: "role_id",
	DeptId: "dept_id",
}

// NewEntSysRoleDeptDao creates and returns a new DAO object for table data access.
func NewEntSysRoleDeptDao() *EntSysRoleDeptDao {
	return &EntSysRoleDeptDao{
		group:   "default",
		table:   "ent_sys_role_dept",
		columns: entSysRoleDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EntSysRoleDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EntSysRoleDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EntSysRoleDeptDao) Columns() EntSysRoleDeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EntSysRoleDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EntSysRoleDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EntSysRoleDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
