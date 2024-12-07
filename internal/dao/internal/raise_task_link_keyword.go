// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RaiseTaskLinkKeywordDao is the data access object for table raise_task_link_keyword.
type RaiseTaskLinkKeywordDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns RaiseTaskLinkKeywordColumns // columns contains all the column names of Table for convenient usage.
}

// RaiseTaskLinkKeywordColumns defines and stores column names for table raise_task_link_keyword.
type RaiseTaskLinkKeywordColumns struct {
	Id          string //
	RaiseTaskId string // 养号任务id
	KeywordId   string // 关键词id
}

// raiseTaskLinkKeywordColumns holds the columns for table raise_task_link_keyword.
var raiseTaskLinkKeywordColumns = RaiseTaskLinkKeywordColumns{
	Id:          "id",
	RaiseTaskId: "raise_task_id",
	KeywordId:   "keyword_id",
}

// NewRaiseTaskLinkKeywordDao creates and returns a new DAO object for table data access.
func NewRaiseTaskLinkKeywordDao() *RaiseTaskLinkKeywordDao {
	return &RaiseTaskLinkKeywordDao{
		group:   "default",
		table:   "raise_task_link_keyword",
		columns: raiseTaskLinkKeywordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RaiseTaskLinkKeywordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RaiseTaskLinkKeywordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RaiseTaskLinkKeywordDao) Columns() RaiseTaskLinkKeywordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RaiseTaskLinkKeywordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RaiseTaskLinkKeywordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RaiseTaskLinkKeywordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
