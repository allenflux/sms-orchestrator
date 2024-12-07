// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceRemarksDao is the data access object for table device_remarks.
type DeviceRemarksDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns DeviceRemarksColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceRemarksColumns defines and stores column names for table device_remarks.
type DeviceRemarksColumns struct {
	Id        string //
	DeviceId  string // 设备ID
	Remark    string // 备注
	CreatedBy string // 创建者
	CreatedAt string // 创建时间
}

// deviceRemarksColumns holds the columns for table device_remarks.
var deviceRemarksColumns = DeviceRemarksColumns{
	Id:        "id",
	DeviceId:  "device_id",
	Remark:    "remark",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
}

// NewDeviceRemarksDao creates and returns a new DAO object for table data access.
func NewDeviceRemarksDao() *DeviceRemarksDao {
	return &DeviceRemarksDao{
		group:   "default",
		table:   "device_remarks",
		columns: deviceRemarksColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceRemarksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceRemarksDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceRemarksDao) Columns() DeviceRemarksColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceRemarksDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceRemarksDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceRemarksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
