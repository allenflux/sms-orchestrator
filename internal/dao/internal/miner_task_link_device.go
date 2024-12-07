// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerTaskLinkDeviceDao is the data access object for table miner_task_link_device.
type MinerTaskLinkDeviceDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns MinerTaskLinkDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// MinerTaskLinkDeviceColumns defines and stores column names for table miner_task_link_device.
type MinerTaskLinkDeviceColumns struct {
	Id          string //
	MinerTaskId string // 矿机任务id
	Platform    string // 平台
	DeviceId    string // 设备id
	DeviceDevId string // 设备本机id
	DeviceName  string // 设备名称
	RunTime     string // 运行时间（分钟）
	Viewed      string // 已刷视频
	Earnings    string // 收益
	Withdrawal  string // 提现
	Balance     string // 余额
	Status      string // 任务状态1待执行2执行中3客户取消/已取消/已终止4客服取消5已暂停6异常7已完成
	ExecAt      string // 执行时间
	CompletedAt string // 完成时间
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
	CreatedBy   string // 创建者id
	Remark      string // 备注
	JdGroupName string // 金盾社区名称
}

// minerTaskLinkDeviceColumns holds the columns for table miner_task_link_device.
var minerTaskLinkDeviceColumns = MinerTaskLinkDeviceColumns{
	Id:          "id",
	MinerTaskId: "miner_task_id",
	Platform:    "platform",
	DeviceId:    "device_id",
	DeviceDevId: "device_dev_id",
	DeviceName:  "device_name",
	RunTime:     "run_time",
	Viewed:      "viewed",
	Earnings:    "earnings",
	Withdrawal:  "withdrawal",
	Balance:     "balance",
	Status:      "status",
	ExecAt:      "exec_at",
	CompletedAt: "completed_at",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreatedBy:   "created_by",
	Remark:      "remark",
	JdGroupName: "jd_group_name",
}

// NewMinerTaskLinkDeviceDao creates and returns a new DAO object for table data access.
func NewMinerTaskLinkDeviceDao() *MinerTaskLinkDeviceDao {
	return &MinerTaskLinkDeviceDao{
		group:   "default",
		table:   "miner_task_link_device",
		columns: minerTaskLinkDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerTaskLinkDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerTaskLinkDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerTaskLinkDeviceDao) Columns() MinerTaskLinkDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerTaskLinkDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerTaskLinkDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerTaskLinkDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
