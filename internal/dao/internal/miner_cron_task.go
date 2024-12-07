// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerCronTaskDao is the data access object for table miner_cron_task.
type MinerCronTaskDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns MinerCronTaskColumns // columns contains all the column names of Table for convenient usage.
}

// MinerCronTaskColumns defines and stores column names for table miner_cron_task.
type MinerCronTaskColumns struct {
	Id          string //
	Platform    string // 任务平台
	RunTime     string // 运行时间（分钟）
	DeviceNum   string // 设备数量
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
	CreatedBy   string // 创建者id
	Remark      string // 备注
	CronMode    string // 定时模式1每日定时2周期定时3循环
	CronStartAt string // 定时计划开始时间
	CronEndAt   string // 定时计划结束时间
	CronExecNum string // 定时任务执行次数
	CronExecAt  string // 定时计划执行时间
	CronGapTime string // 定时计划间隔时间，单位分钟
	CronStatus  string // 开启状态1开启2关闭
	HadExecNum  string // 已执行次数
	Name        string // 任务名称
	HasShow     string // 1-全部显示 2-企业显示 3-群控显示
}

// minerCronTaskColumns holds the columns for table miner_cron_task.
var minerCronTaskColumns = MinerCronTaskColumns{
	Id:          "id",
	Platform:    "platform",
	RunTime:     "run_time",
	DeviceNum:   "device_num",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreatedBy:   "created_by",
	Remark:      "remark",
	CronMode:    "cron_mode",
	CronStartAt: "cron_start_at",
	CronEndAt:   "cron_end_at",
	CronExecNum: "cron_exec_num",
	CronExecAt:  "cron_exec_at",
	CronGapTime: "cron_gap_time",
	CronStatus:  "cron_status",
	HadExecNum:  "had_exec_num",
	Name:        "name",
	HasShow:     "has_show",
}

// NewMinerCronTaskDao creates and returns a new DAO object for table data access.
func NewMinerCronTaskDao() *MinerCronTaskDao {
	return &MinerCronTaskDao{
		group:   "default",
		table:   "miner_cron_task",
		columns: minerCronTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerCronTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerCronTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerCronTaskDao) Columns() MinerCronTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerCronTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerCronTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerCronTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
