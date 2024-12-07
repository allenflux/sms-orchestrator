// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskDao is the data access object for table task.
type TaskDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TaskColumns // columns contains all the column names of Table for convenient usage.
}

// TaskColumns defines and stores column names for table task.
type TaskColumns struct {
	Id              string //
	Name            string // 任务名称
	ClientId        string // 客户id
	ScriptName      string // 脚本名称
	DeviceNum       string // 设备数量
	AppId           string // 平台id
	TaskType        string // 任务类型 1养号2分配账号3刷视频 4 批量解绑账号  5-修改账号信息
	Num             string // 执行次数
	ImplementAt     string // 执行时间
	ImplementType   string // 1-立即执行 2-定时执行
	TimingAt        string // 定时执行时间
	Remark          string // 备注
	ImplementStatus string // 1-待执行 2-执行中 3-客户取消 4-客服取消 5-暂停 6-异常 7-已完成
	OperationStatus string // 1-正常（编辑） 2-取消 3-暂停 4-开始 5-明细
	CompleteAt      string // 任务完成时间
	CancelAt        string // 任务取消时间
	CreatedBy       string // 创建者id
	CreatedAt       string // 创建时间
	UpdatedAt       string // 修改时间
	PlatformType    string // 平台类型 1-facebook
	HasEmpty        string // 1-未清空 2-清空
}

// taskColumns holds the columns for table task.
var taskColumns = TaskColumns{
	Id:              "id",
	Name:            "name",
	ClientId:        "client_id",
	ScriptName:      "script_name",
	DeviceNum:       "device_num",
	AppId:           "app_id",
	TaskType:        "task_type",
	Num:             "num",
	ImplementAt:     "implement_at",
	ImplementType:   "implement_type",
	TimingAt:        "timing_at",
	Remark:          "remark",
	ImplementStatus: "implement_status",
	OperationStatus: "operation_status",
	CompleteAt:      "complete_at",
	CancelAt:        "cancel_at",
	CreatedBy:       "created_by",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	PlatformType:    "platform_type",
	HasEmpty:        "has_empty",
}

// NewTaskDao creates and returns a new DAO object for table data access.
func NewTaskDao() *TaskDao {
	return &TaskDao{
		group:   "default",
		table:   "task",
		columns: taskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskDao) Columns() TaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
