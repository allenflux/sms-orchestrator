// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntSysDictDataDao is the data access object for table ent_sys_dict_data.
type EntSysDictDataDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns EntSysDictDataColumns // columns contains all the column names of Table for convenient usage.
}

// EntSysDictDataColumns defines and stores column names for table ent_sys_dict_data.
type EntSysDictDataColumns struct {
	DictCode  string // 字典编码
	DictSort  string // 字典排序
	DictLabel string // 字典标签
	DictValue string // 字典键值
	DictType  string // 字典类型
	CssClass  string // 样式属性（其他样式扩展）
	ListClass string // 表格回显样式
	IsDefault string // 是否默认（1是 0否）
	Status    string // 状态（0正常 1停用）
	CreateBy  string // 创建者
	UpdateBy  string // 更新者
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// entSysDictDataColumns holds the columns for table ent_sys_dict_data.
var entSysDictDataColumns = EntSysDictDataColumns{
	DictCode:  "dict_code",
	DictSort:  "dict_sort",
	DictLabel: "dict_label",
	DictValue: "dict_value",
	DictType:  "dict_type",
	CssClass:  "css_class",
	ListClass: "list_class",
	IsDefault: "is_default",
	Status:    "status",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewEntSysDictDataDao creates and returns a new DAO object for table data access.
func NewEntSysDictDataDao() *EntSysDictDataDao {
	return &EntSysDictDataDao{
		group:   "default",
		table:   "ent_sys_dict_data",
		columns: entSysDictDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EntSysDictDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EntSysDictDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EntSysDictDataDao) Columns() EntSysDictDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EntSysDictDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EntSysDictDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EntSysDictDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
