// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountDao is the data access object for table account.
type AccountDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns AccountColumns // columns contains all the column names of Table for convenient usage.
}

// AccountColumns defines and stores column names for table account.
type AccountColumns struct {
	Id             string //
	Username       string // 账号
	Password       string // 密码
	LoginType      string // 登录方式1账号密码2验证码
	LoginStatus    string // 登录状态1未分配2已登录3已退出4禁用5异常
	UserInfo       string // 用户信息
	FansNum        string // 粉丝数
	FocusNum       string // 关注数
	LastLoginAt    string // 最后登录时间
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	AccountGroupId string // 账户分组ID
	AccountType    string // 1-账号 2-邮箱 3-手机号
	ClientId       string // 客户ID
	DataImportsId  string // 数据导入ID
	VerifyUrl      string // 验证码地址
	PlatformType   string // 平台 1-facebook
	Sex            string // 性别: 1-女性 2-男性 3-非二次元性别
	Avatar         string // 头像
	Cover          string // 封面图
	Nickname       string // 昵称
	Code           string // 2FA編碼
}

// accountColumns holds the columns for table account.
var accountColumns = AccountColumns{
	Id:             "id",
	Username:       "username",
	Password:       "password",
	LoginType:      "login_type",
	LoginStatus:    "login_status",
	UserInfo:       "user_info",
	FansNum:        "fans_num",
	FocusNum:       "focus_num",
	LastLoginAt:    "last_login_at",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	AccountGroupId: "account_group_id",
	AccountType:    "account_type",
	ClientId:       "client_id",
	DataImportsId:  "data_imports_id",
	VerifyUrl:      "verify_url",
	PlatformType:   "platform_type",
	Sex:            "sex",
	Avatar:         "avatar",
	Cover:          "cover",
	Nickname:       "nickname",
	Code:           "code",
}

// NewAccountDao creates and returns a new DAO object for table data access.
func NewAccountDao() *AccountDao {
	return &AccountDao{
		group:   "default",
		table:   "account",
		columns: accountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AccountDao) Columns() AccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
