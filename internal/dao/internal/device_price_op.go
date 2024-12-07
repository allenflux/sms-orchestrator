// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DevicePriceOpDao is the data access object for table device_price_op.
type DevicePriceOpDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns DevicePriceOpColumns // columns contains all the column names of Table for convenient usage.
}

// DevicePriceOpColumns defines and stores column names for table device_price_op.
type DevicePriceOpColumns struct {
	Id              string //
	DevicePriceId   string // 设备价格表id
	Op              string // 操作1新增2编辑
	Column          string // 列
	ColumnComment   string // 列中文名
	OldValue        string // 原先的值
	OldValueComment string // 原先的值中文名
	Value           string // 值
	ValueComment    string // 值中文名
	CreatedAt       string // 操作时间
	UpdatedBy       string // 操作人
}

// devicePriceOpColumns holds the columns for table device_price_op.
var devicePriceOpColumns = DevicePriceOpColumns{
	Id:              "id",
	DevicePriceId:   "device_price_id",
	Op:              "op",
	Column:          "column",
	ColumnComment:   "column_comment",
	OldValue:        "old_value",
	OldValueComment: "old_value_comment",
	Value:           "value",
	ValueComment:    "value_comment",
	CreatedAt:       "created_at",
	UpdatedBy:       "updated_by",
}

// NewDevicePriceOpDao creates and returns a new DAO object for table data access.
func NewDevicePriceOpDao() *DevicePriceOpDao {
	return &DevicePriceOpDao{
		group:   "default",
		table:   "device_price_op",
		columns: devicePriceOpColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DevicePriceOpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DevicePriceOpDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DevicePriceOpDao) Columns() DevicePriceOpColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DevicePriceOpDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DevicePriceOpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DevicePriceOpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
