// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AllocationTaskDao is the data access object for table allocation_task.
type AllocationTaskDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns AllocationTaskColumns // columns contains all the column names of Table for convenient usage.
}

// AllocationTaskColumns defines and stores column names for table allocation_task.
type AllocationTaskColumns struct {
	Id           string //
	Status       string // 1-待执行 2-执行中 3-已完成 4-已终止
	ClientId     string // 客户ID
	PlatformType string // 平台 1-facebook
	CreatedAt    string //
	UpdatedAt    string //
	Remark       string // 备注
}

// allocationTaskColumns holds the columns for table allocation_task.
var allocationTaskColumns = AllocationTaskColumns{
	Id:           "id",
	Status:       "status",
	ClientId:     "client_id",
	PlatformType: "platform_type",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	Remark:       "remark",
}

// NewAllocationTaskDao creates and returns a new DAO object for table data access.
func NewAllocationTaskDao() *AllocationTaskDao {
	return &AllocationTaskDao{
		group:   "default",
		table:   "allocation_task",
		columns: allocationTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AllocationTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AllocationTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AllocationTaskDao) Columns() AllocationTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AllocationTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AllocationTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AllocationTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
