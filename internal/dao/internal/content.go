// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContentDao is the data access object for table content.
type ContentDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ContentColumns // columns contains all the column names of Table for convenient usage.
}

// ContentColumns defines and stores column names for table content.
type ContentColumns struct {
	Id        string //
	Link      string // 跳转链接
	Title     string // 标题
	Content   string // 内容
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// contentColumns holds the columns for table content.
var contentColumns = ContentColumns{
	Id:        "id",
	Link:      "link",
	Title:     "title",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewContentDao creates and returns a new DAO object for table data access.
func NewContentDao() *ContentDao {
	return &ContentDao{
		group:   "default",
		table:   "content",
		columns: contentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContentDao) Columns() ContentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
