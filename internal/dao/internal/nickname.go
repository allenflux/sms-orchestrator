// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NicknameDao is the data access object for table nickname.
type NicknameDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns NicknameColumns // columns contains all the column names of Table for convenient usage.
}

// NicknameColumns defines and stores column names for table nickname.
type NicknameColumns struct {
	Id        string //
	Name      string // 昵称
	PicTypeId string // 昵称类型
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
	ClientId  string // 客户ID
}

// nicknameColumns holds the columns for table nickname.
var nicknameColumns = NicknameColumns{
	Id:        "id",
	Name:      "name",
	PicTypeId: "pic_type_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	ClientId:  "client_id",
}

// NewNicknameDao creates and returns a new DAO object for table data access.
func NewNicknameDao() *NicknameDao {
	return &NicknameDao{
		group:   "default",
		table:   "nickname",
		columns: nicknameColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NicknameDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NicknameDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NicknameDao) Columns() NicknameColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NicknameDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NicknameDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NicknameDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
