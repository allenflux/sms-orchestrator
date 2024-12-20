// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SmsMissionRecordDao is the data access object for table sms_mission_record.
type SmsMissionRecordDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns SmsMissionRecordColumns // columns contains all the column names of Table for convenient usage.
}

// SmsMissionRecordColumns defines and stores column names for table sms_mission_record.
type SmsMissionRecordColumns struct {
	Id                  string //
	TaskName            string // 任务名称
	SubTaskId           string // 子任务id
	TargetPhoneNumber   string // 目标手机号
	DeviceNumber        string // 执行设备号
	SmsTopic            string // 短信主题
	SmsContent          string // 短信内容
	SmsStatus           string // 短信发送状态，2-失败 1-成功
	AssociatedAccount   string // 所属子账号
	ProjectName         string // 所属项目名称
	ProjectId           string // 所属项目id
	StartTime           string // 开始时间
	CreatedAt           string // 创建时间
	UpdateAt            string // 修改时间
	DeleteAt            string // 删除时间
	AssociatedAccountId string // 子账户id
	RowHash             string // 每行内容的hash串 为了防止同样的记录重复提交
	Reason              string // 失败原因
}

// smsMissionRecordColumns holds the columns for table sms_mission_record.
var smsMissionRecordColumns = SmsMissionRecordColumns{
	Id:                  "id",
	TaskName:            "task_name",
	SubTaskId:           "sub_task_id",
	TargetPhoneNumber:   "target_phone_number",
	DeviceNumber:        "device_number",
	SmsTopic:            "sms_topic",
	SmsContent:          "sms_content",
	SmsStatus:           "sms_status",
	AssociatedAccount:   "associated_account",
	ProjectName:         "project_name",
	ProjectId:           "project_id",
	StartTime:           "start_time",
	CreatedAt:           "created_at",
	UpdateAt:            "update_at",
	DeleteAt:            "delete_at",
	AssociatedAccountId: "associated_account_id",
	RowHash:             "row_hash",
	Reason:              "reason",
}

// NewSmsMissionRecordDao creates and returns a new DAO object for table data access.
func NewSmsMissionRecordDao() *SmsMissionRecordDao {
	return &SmsMissionRecordDao{
		group:   "default",
		table:   "sms_mission_record",
		columns: smsMissionRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SmsMissionRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SmsMissionRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SmsMissionRecordDao) Columns() SmsMissionRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SmsMissionRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SmsMissionRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SmsMissionRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
