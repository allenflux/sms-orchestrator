// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProjectListDao is the data access object for table project_list.
type ProjectListDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ProjectListColumns // columns contains all the column names of Table for convenient usage.
}

// ProjectListColumns defines and stores column names for table project_list.
type ProjectListColumns struct {
	Id                  string //
	ProjectName         string // 项目名称
	QuantityDevice      string // 设备数量
	AssociatedAccount   string // 关联账号
	Note                string // 备注
	CreatedAt           string // 创建时间
	UpdateAt            string // 修改时间
	DeleteAt            string // 删除时间
	AssociatedAccountId string //
}

// projectListColumns holds the columns for table project_list.
var projectListColumns = ProjectListColumns{
	Id:                  "id",
	ProjectName:         "project_name",
	QuantityDevice:      "quantity_device",
	AssociatedAccount:   "associated_account",
	Note:                "note",
	CreatedAt:           "created_at",
	UpdateAt:            "update_at",
	DeleteAt:            "delete_at",
	AssociatedAccountId: "associated_account_id",
}

// NewProjectListDao creates and returns a new DAO object for table data access.
func NewProjectListDao() *ProjectListDao {
	return &ProjectListDao{
		group:   "default",
		table:   "project_list",
		columns: projectListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProjectListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProjectListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProjectListDao) Columns() ProjectListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProjectListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProjectListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProjectListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
