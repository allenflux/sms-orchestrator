// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectTaskLinkDeviceDao is the data access object for table collect_task_link_device.
type CollectTaskLinkDeviceDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns CollectTaskLinkDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// CollectTaskLinkDeviceColumns defines and stores column names for table collect_task_link_device.
type CollectTaskLinkDeviceColumns struct {
	Id            string //
	CollectTaskId string // 采集任务id
	DeviceId      string // 设备id
	DeviceDevId   string // 设备本机id
	Status        string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	Remark        string // 备注（原因）
	CompletedAt   string // 完成时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 修改时间
	DeletedAt     string // 删除时间
}

// collectTaskLinkDeviceColumns holds the columns for table collect_task_link_device.
var collectTaskLinkDeviceColumns = CollectTaskLinkDeviceColumns{
	Id:            "id",
	CollectTaskId: "collect_task_id",
	DeviceId:      "device_id",
	DeviceDevId:   "device_dev_id",
	Status:        "status",
	Remark:        "remark",
	CompletedAt:   "completed_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewCollectTaskLinkDeviceDao creates and returns a new DAO object for table data access.
func NewCollectTaskLinkDeviceDao() *CollectTaskLinkDeviceDao {
	return &CollectTaskLinkDeviceDao{
		group:   "default",
		table:   "collect_task_link_device",
		columns: collectTaskLinkDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CollectTaskLinkDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CollectTaskLinkDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CollectTaskLinkDeviceDao) Columns() CollectTaskLinkDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CollectTaskLinkDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CollectTaskLinkDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CollectTaskLinkDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
