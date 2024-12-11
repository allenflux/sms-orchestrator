// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SmsMissionReportDao is the data access object for table sms_mission_report.
type SmsMissionReportDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns SmsMissionReportColumns // columns contains all the column names of Table for convenient usage.
}

// SmsMissionReportColumns defines and stores column names for table sms_mission_report.
type SmsMissionReportColumns struct {
	Id                  string //
	ProjectId           string // 项目id
	TaskName            string // 任务名称
	FileName            string // 文件名
	DeviceQuota         string // 执行设备
	TaskStatus          string // 任务状态，1-待发送 2-发送中 3-已发送 4-已撤销
	SmsQuantity         string // 短信总条数
	SurplusQuantity     string // 剩余数量
	QuantitySent        string // 已发送数量
	AssociatedAccount   string // 所属子账号
	IntervalTime        string // 间隔时间(秒)
	StartTime           string // 任务开始时间
	CreatedAt           string // 创建时间
	UpdateAt            string // 修改时间
	DeleteAt            string // 删除时间
	ProjectName         string // 项目名称
	AssociatedAccountId string // 所属子账号id
}

// smsMissionReportColumns holds the columns for table sms_mission_report.
var smsMissionReportColumns = SmsMissionReportColumns{
	Id:                  "id",
	ProjectId:           "project_id",
	TaskName:            "task_name",
	FileName:            "file_name",
	DeviceQuota:         "device_quota",
	TaskStatus:          "task_status",
	SmsQuantity:         "sms_quantity",
	SurplusQuantity:     "surplus_quantity",
	QuantitySent:        "quantity_sent",
	AssociatedAccount:   "associated_account",
	IntervalTime:        "interval_time",
	StartTime:           "start_time",
	CreatedAt:           "created_at",
	UpdateAt:            "update_at",
	DeleteAt:            "delete_at",
	ProjectName:         "project_name",
	AssociatedAccountId: "associated_account_id",
}

// NewSmsMissionReportDao creates and returns a new DAO object for table data access.
func NewSmsMissionReportDao() *SmsMissionReportDao {
	return &SmsMissionReportDao{
		group:   "default",
		table:   "sms_mission_report",
		columns: smsMissionReportColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SmsMissionReportDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SmsMissionReportDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SmsMissionReportDao) Columns() SmsMissionReportColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SmsMissionReportDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SmsMissionReportDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SmsMissionReportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
