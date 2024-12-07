// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientDao is the data access object for table client.
type ClientDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ClientColumns // columns contains all the column names of Table for convenient usage.
}

// ClientColumns defines and stores column names for table client.
type ClientColumns struct {
	Id           string //
	LabelId      string // 标签id
	UserName     string // 账户
	UserPassword string // 密码
	UserSalt     string // 密码加密的盐
	UserNickname string // 昵称
	Mobile       string // 手机号
	UserEmail    string // 邮箱
	Avatar       string // 头像
	Balance      string // 余额
	Credit       string // 授信余额
	Advance      string // 预扣金额
	UseCredit    string // 已使用的授信金额
	UserStatus   string // 客户状态0停用1启用2未验证
	CreatedAt    string // 创建时间
	UpdatedAt    string // 修改时间
	DeletedAt    string // 删除时间
	RoleId       string // 角色ID
}

// clientColumns holds the columns for table client.
var clientColumns = ClientColumns{
	Id:           "id",
	LabelId:      "label_id",
	UserName:     "user_name",
	UserPassword: "user_password",
	UserSalt:     "user_salt",
	UserNickname: "user_nickname",
	Mobile:       "mobile",
	UserEmail:    "user_email",
	Avatar:       "avatar",
	Balance:      "balance",
	Credit:       "credit",
	Advance:      "advance",
	UseCredit:    "use_credit",
	UserStatus:   "user_status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	RoleId:       "role_id",
}

// NewClientDao creates and returns a new DAO object for table data access.
func NewClientDao() *ClientDao {
	return &ClientDao{
		group:   "default",
		table:   "client",
		columns: clientColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientDao) Columns() ClientColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
