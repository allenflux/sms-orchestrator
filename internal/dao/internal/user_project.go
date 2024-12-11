// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserProjectDao is the data access object for table user_project.
type UserProjectDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns UserProjectColumns // columns contains all the column names of Table for convenient usage.
}

// UserProjectColumns defines and stores column names for table user_project.
type UserProjectColumns struct {
	Id        string //
	UserId    string // 用户id
	ProjectId string // 项目id
	CreatedAt string // 创建时间
	UpdateAt  string // 修改时间
	DeleteAt  string // 删除时间
}

// userProjectColumns holds the columns for table user_project.
var userProjectColumns = UserProjectColumns{
	Id:        "id",
	UserId:    "user_id",
	ProjectId: "project_id",
	CreatedAt: "created_at",
	UpdateAt:  "update_at",
	DeleteAt:  "delete_at",
}

// NewUserProjectDao creates and returns a new DAO object for table data access.
func NewUserProjectDao() *UserProjectDao {
	return &UserProjectDao{
		group:   "default",
		table:   "user_project",
		columns: userProjectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserProjectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserProjectDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserProjectDao) Columns() UserProjectColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserProjectDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserProjectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserProjectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
