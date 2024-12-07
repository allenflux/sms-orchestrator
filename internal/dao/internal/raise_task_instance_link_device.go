// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RaiseTaskInstanceLinkDeviceDao is the data access object for table raise_task_instance_link_device.
type RaiseTaskInstanceLinkDeviceDao struct {
	table   string                             // table is the underlying table name of the DAO.
	group   string                             // group is the database configuration group name of current DAO.
	columns RaiseTaskInstanceLinkDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// RaiseTaskInstanceLinkDeviceColumns defines and stores column names for table raise_task_instance_link_device.
type RaiseTaskInstanceLinkDeviceColumns struct {
	Id                  string //
	RaiseTaskId         string // 养号定时任务id
	RaiseTaskInstanceId string // 养号任务实例id
	DeviceGroupId       string // 设备组id
	DeviceId            string // 设备id（数据库里自增的）
	DeviceDevId         string // 设备本机id（手机上的）
	CreatedAt           string // 创建时间
	ExecAt              string // 执行时间
	CompletedAt         string // 完成时间
	Status              string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	DeletedAt           string // 删除时间
	Remark              string // 原因
}

// raiseTaskInstanceLinkDeviceColumns holds the columns for table raise_task_instance_link_device.
var raiseTaskInstanceLinkDeviceColumns = RaiseTaskInstanceLinkDeviceColumns{
	Id:                  "id",
	RaiseTaskId:         "raise_task_id",
	RaiseTaskInstanceId: "raise_task_instance_id",
	DeviceGroupId:       "device_group_id",
	DeviceId:            "device_id",
	DeviceDevId:         "device_dev_id",
	CreatedAt:           "created_at",
	ExecAt:              "exec_at",
	CompletedAt:         "completed_at",
	Status:              "status",
	DeletedAt:           "deleted_at",
	Remark:              "remark",
}

// NewRaiseTaskInstanceLinkDeviceDao creates and returns a new DAO object for table data access.
func NewRaiseTaskInstanceLinkDeviceDao() *RaiseTaskInstanceLinkDeviceDao {
	return &RaiseTaskInstanceLinkDeviceDao{
		group:   "default",
		table:   "raise_task_instance_link_device",
		columns: raiseTaskInstanceLinkDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RaiseTaskInstanceLinkDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RaiseTaskInstanceLinkDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RaiseTaskInstanceLinkDeviceDao) Columns() RaiseTaskInstanceLinkDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RaiseTaskInstanceLinkDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RaiseTaskInstanceLinkDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RaiseTaskInstanceLinkDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
