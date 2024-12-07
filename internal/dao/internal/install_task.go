// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InstallTaskDao is the data access object for table install_task.
type InstallTaskDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns InstallTaskColumns // columns contains all the column names of Table for convenient usage.
}

// InstallTaskColumns defines and stores column names for table install_task.
type InstallTaskColumns struct {
	Id        string //
	Remark    string // 备注
	State     string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CreatedBy string // 创建者id（企业用户id）
	CreatedAt string // 创建时间
	DeletedAt string // 删除时间
}

// installTaskColumns holds the columns for table install_task.
var installTaskColumns = InstallTaskColumns{
	Id:        "id",
	Remark:    "remark",
	State:     "state",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewInstallTaskDao creates and returns a new DAO object for table data access.
func NewInstallTaskDao() *InstallTaskDao {
	return &InstallTaskDao{
		group:   "default",
		table:   "install_task",
		columns: installTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InstallTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InstallTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InstallTaskDao) Columns() InstallTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InstallTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InstallTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InstallTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
