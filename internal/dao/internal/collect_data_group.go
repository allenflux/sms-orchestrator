// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectDataGroupDao is the data access object for table collect_data_group.
type CollectDataGroupDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns CollectDataGroupColumns // columns contains all the column names of Table for convenient usage.
}

// CollectDataGroupColumns defines and stores column names for table collect_data_group.
type CollectDataGroupColumns struct {
	Id        string //
	Name      string // 名称
	ClientId  string // 创建客户id
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
}

// collectDataGroupColumns holds the columns for table collect_data_group.
var collectDataGroupColumns = CollectDataGroupColumns{
	Id:        "id",
	Name:      "name",
	ClientId:  "client_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCollectDataGroupDao creates and returns a new DAO object for table data access.
func NewCollectDataGroupDao() *CollectDataGroupDao {
	return &CollectDataGroupDao{
		group:   "default",
		table:   "collect_data_group",
		columns: collectDataGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CollectDataGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CollectDataGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CollectDataGroupDao) Columns() CollectDataGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CollectDataGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CollectDataGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CollectDataGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
