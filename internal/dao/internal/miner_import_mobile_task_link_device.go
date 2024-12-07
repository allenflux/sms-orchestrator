// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerImportMobileTaskLinkDeviceDao is the data access object for table miner_import_mobile_task_link_device.
type MinerImportMobileTaskLinkDeviceDao struct {
	table   string                                 // table is the underlying table name of the DAO.
	group   string                                 // group is the database configuration group name of current DAO.
	columns MinerImportMobileTaskLinkDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// MinerImportMobileTaskLinkDeviceColumns defines and stores column names for table miner_import_mobile_task_link_device.
type MinerImportMobileTaskLinkDeviceColumns struct {
	Id                      string //
	DeviceId                string // 设备主键ID
	DevId                   string // 设备ID
	Status                  string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	MinerImportMobileTaskId string // 矿机导入通讯录任务ID
	CreatedAt               string //
	UpdatedAt               string //
	Identity                string // 唯一标识
	ExecAt                  string // 执行时间
	Remark                  string // 备注
	ClientId                string // 客户ID
	ExecEndAt               string // 执行结束时间
	FailNum                 string // 导入失败通讯录数量
	SuccessNum              string // 导入成功通讯录数量
	JdGroupName             string // 金盾社区名称
}

// minerImportMobileTaskLinkDeviceColumns holds the columns for table miner_import_mobile_task_link_device.
var minerImportMobileTaskLinkDeviceColumns = MinerImportMobileTaskLinkDeviceColumns{
	Id:                      "id",
	DeviceId:                "device_id",
	DevId:                   "dev_id",
	Status:                  "status",
	MinerImportMobileTaskId: "miner_import_mobile_task_id",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	Identity:                "identity",
	ExecAt:                  "exec_at",
	Remark:                  "remark",
	ClientId:                "client_id",
	ExecEndAt:               "exec_end_at",
	FailNum:                 "fail_num",
	SuccessNum:              "success_num",
	JdGroupName:             "jd_group_name",
}

// NewMinerImportMobileTaskLinkDeviceDao creates and returns a new DAO object for table data access.
func NewMinerImportMobileTaskLinkDeviceDao() *MinerImportMobileTaskLinkDeviceDao {
	return &MinerImportMobileTaskLinkDeviceDao{
		group:   "default",
		table:   "miner_import_mobile_task_link_device",
		columns: minerImportMobileTaskLinkDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerImportMobileTaskLinkDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerImportMobileTaskLinkDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerImportMobileTaskLinkDeviceDao) Columns() MinerImportMobileTaskLinkDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerImportMobileTaskLinkDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerImportMobileTaskLinkDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerImportMobileTaskLinkDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
