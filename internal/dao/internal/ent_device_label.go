// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntDeviceLabelDao is the data access object for table ent_device_label.
type EntDeviceLabelDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns EntDeviceLabelColumns // columns contains all the column names of Table for convenient usage.
}

// EntDeviceLabelColumns defines and stores column names for table ent_device_label.
type EntDeviceLabelColumns struct {
	Id        string //
	Name      string // 标签名称
	CreatedAt string //
	CreatedBy string // 创建者id
	DeletedAt string // 删除时间
}

// entDeviceLabelColumns holds the columns for table ent_device_label.
var entDeviceLabelColumns = EntDeviceLabelColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	DeletedAt: "deleted_at",
}

// NewEntDeviceLabelDao creates and returns a new DAO object for table data access.
func NewEntDeviceLabelDao() *EntDeviceLabelDao {
	return &EntDeviceLabelDao{
		group:   "default",
		table:   "ent_device_label",
		columns: entDeviceLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EntDeviceLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EntDeviceLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EntDeviceLabelDao) Columns() EntDeviceLabelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EntDeviceLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EntDeviceLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EntDeviceLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
