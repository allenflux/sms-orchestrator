// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SmsTraceTaskDao is the data access object for table sms_trace_task.
type SmsTraceTaskDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SmsTraceTaskColumns // columns contains all the column names of Table for convenient usage.
}

// SmsTraceTaskColumns defines and stores column names for table sms_trace_task.
type SmsTraceTaskColumns struct {
	Id           string //
	TargetNumber string //
	DeviceNumber string //
	TaskId       string //
	CreateAt     string //
	UpdateAt     string //
	DeleteAt     string //
}

// smsTraceTaskColumns holds the columns for table sms_trace_task.
var smsTraceTaskColumns = SmsTraceTaskColumns{
	Id:           "id",
	TargetNumber: "target_number",
	DeviceNumber: "device_number",
	TaskId:       "task_id",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeleteAt:     "delete_at",
}

// NewSmsTraceTaskDao creates and returns a new DAO object for table data access.
func NewSmsTraceTaskDao() *SmsTraceTaskDao {
	return &SmsTraceTaskDao{
		group:   "default",
		table:   "sms_trace_task",
		columns: smsTraceTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SmsTraceTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SmsTraceTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SmsTraceTaskDao) Columns() SmsTraceTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SmsTraceTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SmsTraceTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SmsTraceTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
