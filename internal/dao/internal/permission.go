// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionDao is the data access object for table permission.
type PermissionDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns PermissionColumns // columns contains all the column names of Table for convenient usage.
}

// PermissionColumns defines and stores column names for table permission.
type PermissionColumns struct {
	Id        string //
	Pid       string // 功能父id
	Name      string // 规则名称
	Title     string // 功能名称
	Icon      string // icon图标
	MenuType  string // 菜单类型，1-目录 2-菜单 3-按钮
	IsHide    string // 显示状态，1-隐藏 2-显示
	Path      string // 路由地址
	Component string // 组件路径
	IsCached  string // 是否缓存，1-不缓存，2-缓存
	CreatedAt string // 创建时间
	UpdateAt  string // 修改时间
	DeleteAt  string // 删除时间
}

// permissionColumns holds the columns for table permission.
var permissionColumns = PermissionColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	Title:     "title",
	Icon:      "icon",
	MenuType:  "menu_type",
	IsHide:    "is_hide",
	Path:      "path",
	Component: "component",
	IsCached:  "is_cached",
	CreatedAt: "created_at",
	UpdateAt:  "update_at",
	DeleteAt:  "delete_at",
}

// NewPermissionDao creates and returns a new DAO object for table data access.
func NewPermissionDao() *PermissionDao {
	return &PermissionDao{
		group:   "default",
		table:   "permission",
		columns: permissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PermissionDao) Columns() PermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
