// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PicTypeDao is the data access object for table pic_type.
type PicTypeDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns PicTypeColumns // columns contains all the column names of Table for convenient usage.
}

// PicTypeColumns defines and stores column names for table pic_type.
type PicTypeColumns struct {
	Id        string //
	Name      string // 图片类型
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	Type      string // 1-头像 2-封面图 3-昵称
	ClientId  string // 用户ID
}

// picTypeColumns holds the columns for table pic_type.
var picTypeColumns = PicTypeColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Type:      "type",
	ClientId:  "client_id",
}

// NewPicTypeDao creates and returns a new DAO object for table data access.
func NewPicTypeDao() *PicTypeDao {
	return &PicTypeDao{
		group:   "default",
		table:   "pic_type",
		columns: picTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PicTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PicTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PicTypeDao) Columns() PicTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PicTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PicTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PicTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
