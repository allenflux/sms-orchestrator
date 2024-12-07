// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceGrantImportFailsDao is the data access object for table device_grant_import_fails.
type DeviceGrantImportFailsDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns DeviceGrantImportFailsColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceGrantImportFailsColumns defines and stores column names for table device_grant_import_fails.
type DeviceGrantImportFailsColumns struct {
	Id            string //
	DataImportsId string // 数据导入ID
	DeviceDevId   string // 设备id
	GrantDb       string // 授权数据库
	DeviceLabels  string // 设备标签
	Remark        string // 备注
	FailReason    string // 失败原因
	CreatedAt     string // 创建时间
	ClientId      string // 客户ID
}

// deviceGrantImportFailsColumns holds the columns for table device_grant_import_fails.
var deviceGrantImportFailsColumns = DeviceGrantImportFailsColumns{
	Id:            "id",
	DataImportsId: "data_imports_id",
	DeviceDevId:   "device_dev_id",
	GrantDb:       "grant_db",
	DeviceLabels:  "device_labels",
	Remark:        "remark",
	FailReason:    "fail_reason",
	CreatedAt:     "created_at",
	ClientId:      "client_id",
}

// NewDeviceGrantImportFailsDao creates and returns a new DAO object for table data access.
func NewDeviceGrantImportFailsDao() *DeviceGrantImportFailsDao {
	return &DeviceGrantImportFailsDao{
		group:   "default",
		table:   "device_grant_import_fails",
		columns: deviceGrantImportFailsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceGrantImportFailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceGrantImportFailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceGrantImportFailsDao) Columns() DeviceGrantImportFailsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceGrantImportFailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceGrantImportFailsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceGrantImportFailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
