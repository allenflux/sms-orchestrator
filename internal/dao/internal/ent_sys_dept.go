// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntSysDeptDao is the data access object for table ent_sys_dept.
type EntSysDeptDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns EntSysDeptColumns // columns contains all the column names of Table for convenient usage.
}

// EntSysDeptColumns defines and stores column names for table ent_sys_dept.
type EntSysDeptColumns struct {
	DeptId    string // 部门id
	ParentId  string // 父部门id
	Ancestors string // 祖级列表
	DeptName  string // 部门名称
	OrderNum  string // 显示顺序
	Leader    string // 负责人
	Phone     string // 联系电话
	Email     string // 邮箱
	Status    string // 部门状态（0正常 1停用）
	CreatedBy string // 创建人
	UpdatedBy string // 修改人
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
}

// entSysDeptColumns holds the columns for table ent_sys_dept.
var entSysDeptColumns = EntSysDeptColumns{
	DeptId:    "dept_id",
	ParentId:  "parent_id",
	Ancestors: "ancestors",
	DeptName:  "dept_name",
	OrderNum:  "order_num",
	Leader:    "leader",
	Phone:     "phone",
	Email:     "email",
	Status:    "status",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewEntSysDeptDao creates and returns a new DAO object for table data access.
func NewEntSysDeptDao() *EntSysDeptDao {
	return &EntSysDeptDao{
		group:   "default",
		table:   "ent_sys_dept",
		columns: entSysDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EntSysDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EntSysDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EntSysDeptDao) Columns() EntSysDeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EntSysDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EntSysDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EntSysDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
