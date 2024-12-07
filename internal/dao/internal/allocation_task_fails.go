// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AllocationTaskFailsDao is the data access object for table allocation_task_fails.
type AllocationTaskFailsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns AllocationTaskFailsColumns // columns contains all the column names of Table for convenient usage.
}

// AllocationTaskFailsColumns defines and stores column names for table allocation_task_fails.
type AllocationTaskFailsColumns struct {
	Id        string //
	AccountId string // 账号ID
	DevId     string // 设备ID
	TaskId    string // 任务表ID
	Remark    string // 备注
	CreatedAt string //
	Username  string // 账号
}

// allocationTaskFailsColumns holds the columns for table allocation_task_fails.
var allocationTaskFailsColumns = AllocationTaskFailsColumns{
	Id:        "id",
	AccountId: "account_id",
	DevId:     "dev_id",
	TaskId:    "task_id",
	Remark:    "remark",
	CreatedAt: "created_at",
	Username:  "username",
}

// NewAllocationTaskFailsDao creates and returns a new DAO object for table data access.
func NewAllocationTaskFailsDao() *AllocationTaskFailsDao {
	return &AllocationTaskFailsDao{
		group:   "default",
		table:   "allocation_task_fails",
		columns: allocationTaskFailsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AllocationTaskFailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AllocationTaskFailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AllocationTaskFailsDao) Columns() AllocationTaskFailsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AllocationTaskFailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AllocationTaskFailsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AllocationTaskFailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
