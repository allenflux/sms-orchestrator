// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KeywordDao is the data access object for table keyword.
type KeywordDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns KeywordColumns // columns contains all the column names of Table for convenient usage.
}

// KeywordColumns defines and stores column names for table keyword.
type KeywordColumns struct {
	Id   string //
	Name string // 关键词
}

// keywordColumns holds the columns for table keyword.
var keywordColumns = KeywordColumns{
	Id:   "id",
	Name: "name",
}

// NewKeywordDao creates and returns a new DAO object for table data access.
func NewKeywordDao() *KeywordDao {
	return &KeywordDao{
		group:   "default",
		table:   "keyword",
		columns: keywordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KeywordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KeywordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KeywordDao) Columns() KeywordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KeywordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KeywordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KeywordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
