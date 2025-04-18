// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SmsChartLogDao is the data access object for table sms_chart_log.
type SmsChartLogDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SmsChartLogColumns // columns contains all the column names of Table for convenient usage.
}

// SmsChartLogColumns defines and stores column names for table sms_chart_log.
type SmsChartLogColumns struct {
	Id                  string //
	ProjectName         string // 项目名称
	ProjectId           string // 项目id
	TargetPhoneNumber   string // 目标手机号
	DeviceNumber        string // 执行设备号
	SmsTopic            string // 短信主题
	SmsContent          string // 短信内容
	SmsStatus           string // 短信发送状态，1-失败 2-成
	AssociatedAccount   string // 所属子账号
	CreatedAt           string // 创建时间
	UpdateAt            string // 修改时间
	DeleteAt            string // 删除时间
	SentOrReceive       string // 1表示此条短信是发送 2表示此条短信是接收
	AssociatedAccountId string // 子账号id
	RowHash             string // log hash 防止重复上报
	TaskId              string // Mission Task的id 任务 id
}

// smsChartLogColumns holds the columns for table sms_chart_log.
var smsChartLogColumns = SmsChartLogColumns{
	Id:                  "id",
	ProjectName:         "project_name",
	ProjectId:           "project_id",
	TargetPhoneNumber:   "target_phone_number",
	DeviceNumber:        "device_number",
	SmsTopic:            "sms_topic",
	SmsContent:          "sms_content",
	SmsStatus:           "sms_status",
	AssociatedAccount:   "associated_account",
	CreatedAt:           "created_at",
	UpdateAt:            "update_at",
	DeleteAt:            "delete_at",
	SentOrReceive:       "sent_or_receive",
	AssociatedAccountId: "associated_account_id",
	RowHash:             "row_hash",
	TaskId:              "task_id",
}

// NewSmsChartLogDao creates and returns a new DAO object for table data access.
func NewSmsChartLogDao() *SmsChartLogDao {
	return &SmsChartLogDao{
		group:   "default",
		table:   "sms_chart_log",
		columns: smsChartLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SmsChartLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SmsChartLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SmsChartLogDao) Columns() SmsChartLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SmsChartLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SmsChartLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SmsChartLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
