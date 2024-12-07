// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EndpointDao is the data access object for table endpoint.
type EndpointDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns EndpointColumns // columns contains all the column names of Table for convenient usage.
}

// EndpointColumns defines and stores column names for table endpoint.
type EndpointColumns struct {
	Id       string //
	ClientId string // 客户id
	DeviceId string // 设备id
}

// endpointColumns holds the columns for table endpoint.
var endpointColumns = EndpointColumns{
	Id:       "id",
	ClientId: "client_id",
	DeviceId: "device_id",
}

// NewEndpointDao creates and returns a new DAO object for table data access.
func NewEndpointDao() *EndpointDao {
	return &EndpointDao{
		group:   "default",
		table:   "endpoint",
		columns: endpointColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EndpointDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EndpointDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EndpointDao) Columns() EndpointColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EndpointDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EndpointDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EndpointDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
