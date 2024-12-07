// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CasbinRule20240821Dao is the data access object for table casbin_rule_20240821.
type CasbinRule20240821Dao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns CasbinRule20240821Columns // columns contains all the column names of Table for convenient usage.
}

// CasbinRule20240821Columns defines and stores column names for table casbin_rule_20240821.
type CasbinRule20240821Columns struct {
	Ptype string //
	V0    string //
	V1    string //
	V2    string //
	V3    string //
	V4    string //
	V5    string //
}

// casbinRule20240821Columns holds the columns for table casbin_rule_20240821.
var casbinRule20240821Columns = CasbinRule20240821Columns{
	Ptype: "ptype",
	V0:    "v0",
	V1:    "v1",
	V2:    "v2",
	V3:    "v3",
	V4:    "v4",
	V5:    "v5",
}

// NewCasbinRule20240821Dao creates and returns a new DAO object for table data access.
func NewCasbinRule20240821Dao() *CasbinRule20240821Dao {
	return &CasbinRule20240821Dao{
		group:   "default",
		table:   "casbin_rule_20240821",
		columns: casbinRule20240821Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CasbinRule20240821Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CasbinRule20240821Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CasbinRule20240821Dao) Columns() CasbinRule20240821Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CasbinRule20240821Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CasbinRule20240821Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CasbinRule20240821Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
