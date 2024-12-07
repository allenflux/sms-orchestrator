// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NicknameTypeDao is the data access object for table nickname_type.
type NicknameTypeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns NicknameTypeColumns // columns contains all the column names of Table for convenient usage.
}

// NicknameTypeColumns defines and stores column names for table nickname_type.
type NicknameTypeColumns struct {
	Id        string //
	Name      string // 昵称类型
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeteledAt string // 删除时间
}

// nicknameTypeColumns holds the columns for table nickname_type.
var nicknameTypeColumns = NicknameTypeColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeteledAt: "deteled_at",
}

// NewNicknameTypeDao creates and returns a new DAO object for table data access.
func NewNicknameTypeDao() *NicknameTypeDao {
	return &NicknameTypeDao{
		group:   "default",
		table:   "nickname_type",
		columns: nicknameTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NicknameTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NicknameTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NicknameTypeDao) Columns() NicknameTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NicknameTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NicknameTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NicknameTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
