// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RaiseTaskLinkWordsGroupDao is the data access object for table raise_task_link_words_group.
type RaiseTaskLinkWordsGroupDao struct {
	table   string                         // table is the underlying table name of the DAO.
	group   string                         // group is the database configuration group name of current DAO.
	columns RaiseTaskLinkWordsGroupColumns // columns contains all the column names of Table for convenient usage.
}

// RaiseTaskLinkWordsGroupColumns defines and stores column names for table raise_task_link_words_group.
type RaiseTaskLinkWordsGroupColumns struct {
	Id           string //
	RaiseTaskId  string // 养号任务id
	WordsGroupId string // 话术分组id
}

// raiseTaskLinkWordsGroupColumns holds the columns for table raise_task_link_words_group.
var raiseTaskLinkWordsGroupColumns = RaiseTaskLinkWordsGroupColumns{
	Id:           "id",
	RaiseTaskId:  "raise_task_id",
	WordsGroupId: "words_group_id",
}

// NewRaiseTaskLinkWordsGroupDao creates and returns a new DAO object for table data access.
func NewRaiseTaskLinkWordsGroupDao() *RaiseTaskLinkWordsGroupDao {
	return &RaiseTaskLinkWordsGroupDao{
		group:   "default",
		table:   "raise_task_link_words_group",
		columns: raiseTaskLinkWordsGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RaiseTaskLinkWordsGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RaiseTaskLinkWordsGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RaiseTaskLinkWordsGroupDao) Columns() RaiseTaskLinkWordsGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RaiseTaskLinkWordsGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RaiseTaskLinkWordsGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RaiseTaskLinkWordsGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
