// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerTaskDao is the data access object for table miner_task.
type MinerTaskDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns MinerTaskColumns // columns contains all the column names of Table for convenient usage.
}

// MinerTaskColumns defines and stores column names for table miner_task.
type MinerTaskColumns struct {
	Id              string //
	MinerCronTaskId string // 矿机定时任务id
	Name            string // 任务名称
	Platform        string // 任务平台
	RunTime         string // 运行时间（分钟）
	DeviceNum       string // 设备数量
	CreatedAt       string // 创建时间
	ExecAt          string // 执行时间
	UpdatedAt       string // 修改时间
	DeletedAt       string // 删除时间
	CreatedBy       string // 创建者id
	CompletedAt     string // 完成时间
	Remark          string // 备注
	Status          string // 任务状态1待执行2执行中3客户取消/已取消/已终止4客服取消5已暂停6异常7已完成
	ExecMode        string // 任务模式1单次2定时
	HasShow         string // 1-展示 2-不展示
}

// minerTaskColumns holds the columns for table miner_task.
var minerTaskColumns = MinerTaskColumns{
	Id:              "id",
	MinerCronTaskId: "miner_cron_task_id",
	Name:            "name",
	Platform:        "platform",
	RunTime:         "run_time",
	DeviceNum:       "device_num",
	CreatedAt:       "created_at",
	ExecAt:          "exec_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	CreatedBy:       "created_by",
	CompletedAt:     "completed_at",
	Remark:          "remark",
	Status:          "status",
	ExecMode:        "exec_mode",
	HasShow:         "has_show",
}

// NewMinerTaskDao creates and returns a new DAO object for table data access.
func NewMinerTaskDao() *MinerTaskDao {
	return &MinerTaskDao{
		group:   "default",
		table:   "miner_task",
		columns: minerTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerTaskDao) Columns() MinerTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
