// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApkSetDao is the data access object for table apk_set.
type ApkSetDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ApkSetColumns // columns contains all the column names of Table for convenient usage.
}

// ApkSetColumns defines and stores column names for table apk_set.
type ApkSetColumns struct {
	Id            string //
	SetId         string // 集合id
	ScriptId      string // 脚本id
	ScriptAppName string // 脚本包名
	Status        string // 1未安装2已安装
	CreatedAt     string // 创建时间
	UpdatedAt     string // 修改时间
}

// apkSetColumns holds the columns for table apk_set.
var apkSetColumns = ApkSetColumns{
	Id:            "id",
	SetId:         "set_id",
	ScriptId:      "script_id",
	ScriptAppName: "script_app_name",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewApkSetDao creates and returns a new DAO object for table data access.
func NewApkSetDao() *ApkSetDao {
	return &ApkSetDao{
		group:   "default",
		table:   "apk_set",
		columns: apkSetColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApkSetDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApkSetDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApkSetDao) Columns() ApkSetColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApkSetDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApkSetDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApkSetDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
