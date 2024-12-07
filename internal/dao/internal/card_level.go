// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CardLevelDao is the data access object for table card_level.
type CardLevelDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns CardLevelColumns // columns contains all the column names of Table for convenient usage.
}

// CardLevelColumns defines and stores column names for table card_level.
type CardLevelColumns struct {
	Id        string //
	Name      string // 卡片等级名称
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	File      string // 文件base64值
}

// cardLevelColumns holds the columns for table card_level.
var cardLevelColumns = CardLevelColumns{
	Id:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	File:      "file",
}

// NewCardLevelDao creates and returns a new DAO object for table data access.
func NewCardLevelDao() *CardLevelDao {
	return &CardLevelDao{
		group:   "default",
		table:   "card_level",
		columns: cardLevelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CardLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CardLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CardLevelDao) Columns() CardLevelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CardLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CardLevelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CardLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
