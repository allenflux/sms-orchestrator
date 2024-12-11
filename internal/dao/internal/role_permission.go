// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RolePermissionDao is the data access object for table role_permission.
type RolePermissionDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns RolePermissionColumns // columns contains all the column names of Table for convenient usage.
}

// RolePermissionColumns defines and stores column names for table role_permission.
type RolePermissionColumns struct {
	Id           string //
	RoleId       string // 角色id
	PermissionId string // 权限id
	CreatedAt    string // 创建时间
	UpdateAt     string // 修改时间
	DeleteAt     string // 删除时间
}

// rolePermissionColumns holds the columns for table role_permission.
var rolePermissionColumns = RolePermissionColumns{
	Id:           "id",
	RoleId:       "role_id",
	PermissionId: "permission_id",
	CreatedAt:    "created_at",
	UpdateAt:     "update_at",
	DeleteAt:     "delete_at",
}

// NewRolePermissionDao creates and returns a new DAO object for table data access.
func NewRolePermissionDao() *RolePermissionDao {
	return &RolePermissionDao{
		group:   "default",
		table:   "role_permission",
		columns: rolePermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RolePermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RolePermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RolePermissionDao) Columns() RolePermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RolePermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RolePermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RolePermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
