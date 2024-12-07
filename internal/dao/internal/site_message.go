// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SiteMessageDao is the data access object for table site_message.
type SiteMessageDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SiteMessageColumns // columns contains all the column names of Table for convenient usage.
}

// SiteMessageColumns defines and stores column names for table site_message.
type SiteMessageColumns struct {
	Id            string //
	Title         string // 标题
	Content       string // 消息
	AppUserIdStrs string // APP用户ID，逗号隔开
	AdminUsername string // 管理员账号
	HasAll        string // 0-非所有  1-所有
	CreatedAt     string //
	UpdatedAt     string //
}

// siteMessageColumns holds the columns for table site_message.
var siteMessageColumns = SiteMessageColumns{
	Id:            "id",
	Title:         "title",
	Content:       "content",
	AppUserIdStrs: "app_user_id_strs",
	AdminUsername: "admin_username",
	HasAll:        "has_all",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSiteMessageDao creates and returns a new DAO object for table data access.
func NewSiteMessageDao() *SiteMessageDao {
	return &SiteMessageDao{
		group:   "default",
		table:   "site_message",
		columns: siteMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SiteMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SiteMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SiteMessageDao) Columns() SiteMessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SiteMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SiteMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SiteMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
