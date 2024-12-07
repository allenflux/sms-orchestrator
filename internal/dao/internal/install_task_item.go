// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InstallTaskItemDao is the data access object for table install_task_item.
type InstallTaskItemDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns InstallTaskItemColumns // columns contains all the column names of Table for convenient usage.
}

// InstallTaskItemColumns defines and stores column names for table install_task_item.
type InstallTaskItemColumns struct {
	Id            string //
	InstallTaskId string // 安装任务id
	DeviceId      string // 设备id
	ScriptId      string // 脚本id
	State         string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CreatedAt     string // 创建时间
}

// installTaskItemColumns holds the columns for table install_task_item.
var installTaskItemColumns = InstallTaskItemColumns{
	Id:            "id",
	InstallTaskId: "install_task_id",
	DeviceId:      "device_id",
	ScriptId:      "script_id",
	State:         "state",
	CreatedAt:     "created_at",
}

// NewInstallTaskItemDao creates and returns a new DAO object for table data access.
func NewInstallTaskItemDao() *InstallTaskItemDao {
	return &InstallTaskItemDao{
		group:   "default",
		table:   "install_task_item",
		columns: installTaskItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InstallTaskItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InstallTaskItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InstallTaskItemDao) Columns() InstallTaskItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InstallTaskItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InstallTaskItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InstallTaskItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
