// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PicDao is the data access object for table pic.
type PicDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns PicColumns // columns contains all the column names of Table for convenient usage.
}

// PicColumns defines and stores column names for table pic.
type PicColumns struct {
	Id        string //
	Name      string // 名称
	Url       string // 图片地址
	PicTypeId string // 类型ID
	Remark    string // 备注
	CreatedAt string //
	UpdatedAt string //
	Type      string // 1-头像 2-封面图
	ClientId  string // 用户ID
}

// picColumns holds the columns for table pic.
var picColumns = PicColumns{
	Id:        "id",
	Name:      "name",
	Url:       "url",
	PicTypeId: "pic_type_id",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Type:      "type",
	ClientId:  "client_id",
}

// NewPicDao creates and returns a new DAO object for table data access.
func NewPicDao() *PicDao {
	return &PicDao{
		group:   "default",
		table:   "pic",
		columns: picColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PicDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PicDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PicDao) Columns() PicColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PicDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PicDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PicDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
