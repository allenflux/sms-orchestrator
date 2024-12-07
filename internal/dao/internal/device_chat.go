// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceChatDao is the data access object for table device_chat.
type DeviceChatDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns DeviceChatColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceChatColumns defines and stores column names for table device_chat.
type DeviceChatColumns struct {
	Id           string //
	DeviceId     string // id
	DeviceDevId  string // id
	LastChatTime string //
}

// deviceChatColumns holds the columns for table device_chat.
var deviceChatColumns = DeviceChatColumns{
	Id:           "id",
	DeviceId:     "device_id",
	DeviceDevId:  "device_dev_id",
	LastChatTime: "last_chat_time",
}

// NewDeviceChatDao creates and returns a new DAO object for table data access.
func NewDeviceChatDao() *DeviceChatDao {
	return &DeviceChatDao{
		group:   "default",
		table:   "device_chat",
		columns: deviceChatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceChatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceChatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceChatDao) Columns() DeviceChatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceChatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceChatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceChatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
