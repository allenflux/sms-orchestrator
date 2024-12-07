// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DataImportDetailsDao is the data access object for table data_import_details.
type DataImportDetailsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns DataImportDetailsColumns // columns contains all the column names of Table for convenient usage.
}

// DataImportDetailsColumns defines and stores column names for table data_import_details.
type DataImportDetailsColumns struct {
	Id           string //
	DataImportId string // 导入ID
	DevId        string // 设备ID
	Remark1      string // 备注1
	Remark2      string // 备注2
	CreatedAt    string //
	Status       string // 1-成功 2-失败
}

// dataImportDetailsColumns holds the columns for table data_import_details.
var dataImportDetailsColumns = DataImportDetailsColumns{
	Id:           "id",
	DataImportId: "data_import_id",
	DevId:        "dev_id",
	Remark1:      "remark1",
	Remark2:      "remark2",
	CreatedAt:    "created_at",
	Status:       "status",
}

// NewDataImportDetailsDao creates and returns a new DAO object for table data access.
func NewDataImportDetailsDao() *DataImportDetailsDao {
	return &DataImportDetailsDao{
		group:   "default",
		table:   "data_import_details",
		columns: dataImportDetailsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DataImportDetailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DataImportDetailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DataImportDetailsDao) Columns() DataImportDetailsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DataImportDetailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DataImportDetailsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DataImportDetailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
