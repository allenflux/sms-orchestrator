// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WordsGroupDetailsDao is the data access object for table words_group_details.
type WordsGroupDetailsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns WordsGroupDetailsColumns // columns contains all the column names of Table for convenient usage.
}

// WordsGroupDetailsColumns defines and stores column names for table words_group_details.
type WordsGroupDetailsColumns struct {
	Id           string //
	WordsId      string // 话术ID
	PlatformType string // 平台类型 1-facebook
	WordsType    string // 话术类型 1养号2批量私信
	CreatedAt    string //
	UpdatedAt    string //
	WordsGroupId string // 话术分组ID
}

// wordsGroupDetailsColumns holds the columns for table words_group_details.
var wordsGroupDetailsColumns = WordsGroupDetailsColumns{
	Id:           "id",
	WordsId:      "words_id",
	PlatformType: "platform_type",
	WordsType:    "words_type",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	WordsGroupId: "words_group_id",
}

// NewWordsGroupDetailsDao creates and returns a new DAO object for table data access.
func NewWordsGroupDetailsDao() *WordsGroupDetailsDao {
	return &WordsGroupDetailsDao{
		group:   "default",
		table:   "words_group_details",
		columns: wordsGroupDetailsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WordsGroupDetailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WordsGroupDetailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WordsGroupDetailsDao) Columns() WordsGroupDetailsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WordsGroupDetailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WordsGroupDetailsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WordsGroupDetailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
