// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NicknameImportFailsDao is the data access object for table nickname_import_fails.
type NicknameImportFailsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns NicknameImportFailsColumns // columns contains all the column names of Table for convenient usage.
}

// NicknameImportFailsColumns defines and stores column names for table nickname_import_fails.
type NicknameImportFailsColumns struct {
	Id            string //
	Nickname      string // 昵称
	TypeName      string // 昵称类型
	FailReason    string // 失败原因
	DataImportsId string // 数据导入ID
	CreatedAt     string //
	ClientId      string // 客户ID
}

// nicknameImportFailsColumns holds the columns for table nickname_import_fails.
var nicknameImportFailsColumns = NicknameImportFailsColumns{
	Id:            "id",
	Nickname:      "nickname",
	TypeName:      "type_name",
	FailReason:    "fail_reason",
	DataImportsId: "data_imports_id",
	CreatedAt:     "created_at",
	ClientId:      "client_id",
}

// NewNicknameImportFailsDao creates and returns a new DAO object for table data access.
func NewNicknameImportFailsDao() *NicknameImportFailsDao {
	return &NicknameImportFailsDao{
		group:   "default",
		table:   "nickname_import_fails",
		columns: nicknameImportFailsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NicknameImportFailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NicknameImportFailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NicknameImportFailsDao) Columns() NicknameImportFailsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NicknameImportFailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NicknameImportFailsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NicknameImportFailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
