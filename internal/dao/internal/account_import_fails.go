// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountImportFailsDao is the data access object for table account_import_fails.
type AccountImportFailsDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns AccountImportFailsColumns // columns contains all the column names of Table for convenient usage.
}

// AccountImportFailsColumns defines and stores column names for table account_import_fails.
type AccountImportFailsColumns struct {
	Id            string //
	LoginType     string //
	AccountType   string //
	Username      string //
	Password      string //
	VerifyUrl     string //
	FailReason    string //
	DataImportsId string //
	CreatedAt     string //
	UpdatedAt     string //
	DeletedAt     string //
	Code          string //
}

// accountImportFailsColumns holds the columns for table account_import_fails.
var accountImportFailsColumns = AccountImportFailsColumns{
	Id:            "id",
	LoginType:     "login_type",
	AccountType:   "account_type",
	Username:      "username",
	Password:      "password",
	VerifyUrl:     "verify_url",
	FailReason:    "fail_reason",
	DataImportsId: "data_imports_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	Code:          "code",
}

// NewAccountImportFailsDao creates and returns a new DAO object for table data access.
func NewAccountImportFailsDao() *AccountImportFailsDao {
	return &AccountImportFailsDao{
		group:   "default",
		table:   "account_import_fails",
		columns: accountImportFailsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AccountImportFailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccountImportFailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AccountImportFailsDao) Columns() AccountImportFailsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccountImportFailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccountImportFailsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccountImportFailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
