// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogDao is the data access object for table log.
type LogDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns LogColumns // columns contains all the column names of Table for convenient usage.
}

// LogColumns defines and stores column names for table log.
type LogColumns struct {
	Id        string //
	UserId    string // 操作者的用户id
	UserName  string // 操作者的用户名
	ClientIp  string // 操作ip
	Function  string // 操作功能
	Note      string // 操作内容
	CreatedAt string // 创建时间
	UpdateAt  string // 修改时间
	DeleteAt  string // 删除时间
}

// logColumns holds the columns for table log.
var logColumns = LogColumns{
	Id:        "id",
	UserId:    "user_id",
	UserName:  "user_name",
	ClientIp:  "client_ip",
	Function:  "function",
	Note:      "note",
	CreatedAt: "created_at",
	UpdateAt:  "update_at",
	DeleteAt:  "delete_at",
}

// NewLogDao creates and returns a new DAO object for table data access.
func NewLogDao() *LogDao {
	return &LogDao{
		group:   "default",
		table:   "log",
		columns: logColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LogDao) Columns() LogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
