// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerLoginTaskLinkDeviceDao is the data access object for table miner_login_task_link_device.
type MinerLoginTaskLinkDeviceDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns MinerLoginTaskLinkDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// MinerLoginTaskLinkDeviceColumns defines and stores column names for table miner_login_task_link_device.
type MinerLoginTaskLinkDeviceColumns struct {
	Id               string //
	DevId            string // 设备ID
	Status           string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	MinerLoginTaskId string // 矿机登录任务ID
	CreatedAt        string //
	UpdatedAt        string //
	Identity         string // 唯一标识
	ExecAt           string // 开始执行时间
	Remark           string // 备注
	DeviceId         string // 设备主键ID
	Platform         string // 平台
	ClientId         string // 客户ID
	ExecEndAt        string // 执行结束时间
	JdGroupName      string // 金盾社区名称
}

// minerLoginTaskLinkDeviceColumns holds the columns for table miner_login_task_link_device.
var minerLoginTaskLinkDeviceColumns = MinerLoginTaskLinkDeviceColumns{
	Id:               "id",
	DevId:            "dev_id",
	Status:           "status",
	MinerLoginTaskId: "miner_login_task_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	Identity:         "identity",
	ExecAt:           "exec_at",
	Remark:           "remark",
	DeviceId:         "device_id",
	Platform:         "platform",
	ClientId:         "client_id",
	ExecEndAt:        "exec_end_at",
	JdGroupName:      "jd_group_name",
}

// NewMinerLoginTaskLinkDeviceDao creates and returns a new DAO object for table data access.
func NewMinerLoginTaskLinkDeviceDao() *MinerLoginTaskLinkDeviceDao {
	return &MinerLoginTaskLinkDeviceDao{
		group:   "default",
		table:   "miner_login_task_link_device",
		columns: minerLoginTaskLinkDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerLoginTaskLinkDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerLoginTaskLinkDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerLoginTaskLinkDeviceDao) Columns() MinerLoginTaskLinkDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerLoginTaskLinkDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerLoginTaskLinkDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerLoginTaskLinkDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
