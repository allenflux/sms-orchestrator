// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceGroupDao is the data access object for table device_group.
type DeviceGroupDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns DeviceGroupColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceGroupColumns defines and stores column names for table device_group.
type DeviceGroupColumns struct {
	Id        string //
	Name      string // 分组名称
	DeviceNum string // 设备数量
	Remark    string // 备注
	ClientId  string // 客户id
	CreatedAt string // 创建时间
	UpdatedAt string //
	DeletedAt string //
}

// deviceGroupColumns holds the columns for table device_group.
var deviceGroupColumns = DeviceGroupColumns{
	Id:        "id",
	Name:      "name",
	DeviceNum: "device_num",
	Remark:    "remark",
	ClientId:  "client_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewDeviceGroupDao creates and returns a new DAO object for table data access.
func NewDeviceGroupDao() *DeviceGroupDao {
	return &DeviceGroupDao{
		group:   "default",
		table:   "device_group",
		columns: deviceGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceGroupDao) Columns() DeviceGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
