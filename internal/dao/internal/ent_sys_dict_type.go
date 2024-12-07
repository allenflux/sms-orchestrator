// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntSysDictTypeDao is the data access object for table ent_sys_dict_type.
type EntSysDictTypeDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns EntSysDictTypeColumns // columns contains all the column names of Table for convenient usage.
}

// EntSysDictTypeColumns defines and stores column names for table ent_sys_dict_type.
type EntSysDictTypeColumns struct {
	DictId    string // 字典主键
	DictName  string // 字典名称
	DictType  string // 字典类型
	Status    string // 状态（0正常 1停用）
	CreateBy  string // 创建者
	UpdateBy  string // 更新者
	Remark    string // 备注
	CreatedAt string // 创建日期
	UpdatedAt string // 修改日期
}

// entSysDictTypeColumns holds the columns for table ent_sys_dict_type.
var entSysDictTypeColumns = EntSysDictTypeColumns{
	DictId:    "dict_id",
	DictName:  "dict_name",
	DictType:  "dict_type",
	Status:    "status",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewEntSysDictTypeDao creates and returns a new DAO object for table data access.
func NewEntSysDictTypeDao() *EntSysDictTypeDao {
	return &EntSysDictTypeDao{
		group:   "default",
		table:   "ent_sys_dict_type",
		columns: entSysDictTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EntSysDictTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EntSysDictTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EntSysDictTypeDao) Columns() EntSysDictTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EntSysDictTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EntSysDictTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EntSysDictTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
