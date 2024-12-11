// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SubGroupDao is the data access object for table sub_group.
type SubGroupDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SubGroupColumns // columns contains all the column names of Table for convenient usage.
}

// SubGroupColumns defines and stores column names for table sub_group.
type SubGroupColumns struct {
	Id           string //
	SubUserId    string // 子账号id
	ProjectId    string // 项目id
	SubGroupName string // 分组名称
	CreatedAt    string // 创建时间
	UpdateAt     string // 修改时间
	DeleteAt     string // 删除时间
}

// subGroupColumns holds the columns for table sub_group.
var subGroupColumns = SubGroupColumns{
	Id:           "id",
	SubUserId:    "sub_user_id",
	ProjectId:    "project_id",
	SubGroupName: "sub_group_name",
	CreatedAt:    "created_at",
	UpdateAt:     "update_at",
	DeleteAt:     "delete_at",
}

// NewSubGroupDao creates and returns a new DAO object for table data access.
func NewSubGroupDao() *SubGroupDao {
	return &SubGroupDao{
		group:   "default",
		table:   "sub_group",
		columns: subGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SubGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SubGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SubGroupDao) Columns() SubGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SubGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SubGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SubGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
