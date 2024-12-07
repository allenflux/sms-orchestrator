// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntDeviceLinkLabelDao is the data access object for table ent_device_link_label.
type EntDeviceLinkLabelDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns EntDeviceLinkLabelColumns // columns contains all the column names of Table for convenient usage.
}

// EntDeviceLinkLabelColumns defines and stores column names for table ent_device_link_label.
type EntDeviceLinkLabelColumns struct {
	Id            string //
	DeviceLabelId string // 标签ID
	DevId         string // 设备ID
	CreatedAt     string //
}

// entDeviceLinkLabelColumns holds the columns for table ent_device_link_label.
var entDeviceLinkLabelColumns = EntDeviceLinkLabelColumns{
	Id:            "id",
	DeviceLabelId: "device_label_id",
	DevId:         "dev_id",
	CreatedAt:     "created_at",
}

// NewEntDeviceLinkLabelDao creates and returns a new DAO object for table data access.
func NewEntDeviceLinkLabelDao() *EntDeviceLinkLabelDao {
	return &EntDeviceLinkLabelDao{
		group:   "default",
		table:   "ent_device_link_label",
		columns: entDeviceLinkLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EntDeviceLinkLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EntDeviceLinkLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EntDeviceLinkLabelDao) Columns() EntDeviceLinkLabelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EntDeviceLinkLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EntDeviceLinkLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EntDeviceLinkLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
