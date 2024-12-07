// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ScriptDao is the data access object for table script.
type ScriptDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ScriptColumns // columns contains all the column names of Table for convenient usage.
}

// ScriptColumns defines and stores column names for table script.
type ScriptColumns struct {
	Id          string //
	Name        string // 脚本名称
	ClientId    string // 客户id
	DownloadUrl string // apk名称
	AppName     string //
	ObjectName  string // oss的object_name
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	Remark      string // 备注
	DeletedAt   string //
}

// scriptColumns holds the columns for table script.
var scriptColumns = ScriptColumns{
	Id:          "id",
	Name:        "name",
	ClientId:    "client_id",
	DownloadUrl: "download_url",
	AppName:     "app_name",
	ObjectName:  "object_name",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Remark:      "remark",
	DeletedAt:   "deleted_at",
}

// NewScriptDao creates and returns a new DAO object for table data access.
func NewScriptDao() *ScriptDao {
	return &ScriptDao{
		group:   "default",
		table:   "script",
		columns: scriptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ScriptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ScriptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ScriptDao) Columns() ScriptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ScriptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ScriptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ScriptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
