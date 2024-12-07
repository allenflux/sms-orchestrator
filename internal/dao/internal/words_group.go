// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WordsGroupDao is the data access object for table words_group.
type WordsGroupDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns WordsGroupColumns // columns contains all the column names of Table for convenient usage.
}

// WordsGroupColumns defines and stores column names for table words_group.
type WordsGroupColumns struct {
	Id           string //
	Name         string // 分组名称
	PlatformType string // (对应字典platform_type)多个逗号隔开 使用平台1Facebook
	WordsType    string // (对应字典words_type)多个逗号隔开 话术类型1养号2批量私信
	CreatedAt    string // 创建时间
	UpdatedAt    string // 修改时间
	ClientId     string // 客户ID
}

// wordsGroupColumns holds the columns for table words_group.
var wordsGroupColumns = WordsGroupColumns{
	Id:           "id",
	Name:         "name",
	PlatformType: "platform_type",
	WordsType:    "words_type",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ClientId:     "client_id",
}

// NewWordsGroupDao creates and returns a new DAO object for table data access.
func NewWordsGroupDao() *WordsGroupDao {
	return &WordsGroupDao{
		group:   "default",
		table:   "words_group",
		columns: wordsGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WordsGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WordsGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WordsGroupDao) Columns() WordsGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WordsGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WordsGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WordsGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
