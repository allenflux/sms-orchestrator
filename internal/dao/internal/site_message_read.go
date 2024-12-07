// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SiteMessageReadDao is the data access object for table site_message_read.
type SiteMessageReadDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SiteMessageReadColumns // columns contains all the column names of Table for convenient usage.
}

// SiteMessageReadColumns defines and stores column names for table site_message_read.
type SiteMessageReadColumns struct {
	Id        string //
	AppUserId string // APP用户ID
	MessageId string // 站内信主键ID
	CreatedAt string //
	UpdatedAt string //
}

// siteMessageReadColumns holds the columns for table site_message_read.
var siteMessageReadColumns = SiteMessageReadColumns{
	Id:        "id",
	AppUserId: "app_user_id",
	MessageId: "message_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSiteMessageReadDao creates and returns a new DAO object for table data access.
func NewSiteMessageReadDao() *SiteMessageReadDao {
	return &SiteMessageReadDao{
		group:   "default",
		table:   "site_message_read",
		columns: siteMessageReadColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SiteMessageReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SiteMessageReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SiteMessageReadDao) Columns() SiteMessageReadColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SiteMessageReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SiteMessageReadDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SiteMessageReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
