// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DataImportsDao is the data access object for table data_imports.
type DataImportsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns DataImportsColumns // columns contains all the column names of Table for convenient usage.
}

// DataImportsColumns defines and stores column names for table data_imports.
type DataImportsColumns struct {
	Id           string //
	Type         string // 1-账户数据导入 2-昵称导入 3-授权导入 4-设备标签导入 5-app标签导入
	ClientId     string // 客户ID
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
	SuccessNum   string // 成功数量
	FailNum      string // 失败数量
	FileName     string // 文件名
	TotalNum     string // 总数量
	PlatformType string // 平台 1-facebook
	UserType     string // 1-群控 2-企业
}

// dataImportsColumns holds the columns for table data_imports.
var dataImportsColumns = DataImportsColumns{
	Id:           "id",
	Type:         "type",
	ClientId:     "client_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	SuccessNum:   "success_num",
	FailNum:      "fail_num",
	FileName:     "file_name",
	TotalNum:     "total_num",
	PlatformType: "platform_type",
	UserType:     "user_type",
}

// NewDataImportsDao creates and returns a new DAO object for table data access.
func NewDataImportsDao() *DataImportsDao {
	return &DataImportsDao{
		group:   "default",
		table:   "data_imports",
		columns: dataImportsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DataImportsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DataImportsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DataImportsDao) Columns() DataImportsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DataImportsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DataImportsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DataImportsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
