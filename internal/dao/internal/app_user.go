// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppUserDao is the data access object for table app_user.
type AppUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns AppUserColumns // columns contains all the column names of Table for convenient usage.
}

// AppUserColumns defines and stores column names for table app_user.
type AppUserColumns struct {
	Id                  string //
	Password            string // 密码
	Salt                string // 盐
	TransactionPassword string // 交易密码
	TransactionSalt     string // 交易密码的盐
	CreatedAt           string // 创建时间
	UpdatedAt           string // 修改时间
	DeletedAt           string // 删除时间
	Uid                 string // 银行uid
	Mobile              string // 手机号
	Status              string // 1-待KYC 2-待上传 3-待配卡 4-待申请 5-已申请
	KycStatus           string // 1-待kyc 2-已kyc
	UserCardStatus      string // 1-未上传 2-已上传
	BankCardStatus      string // 1-未配卡 2-已配卡
	BankImage           string // 万事达照片
	CardLevelId         string // 卡片等级表ID
	UserCardImage       string // 用户卡照片
	Remark              string // 备注
	Email               string // 邮箱
}

// appUserColumns holds the columns for table app_user.
var appUserColumns = AppUserColumns{
	Id:                  "id",
	Password:            "password",
	Salt:                "salt",
	TransactionPassword: "transaction_password",
	TransactionSalt:     "transaction_salt",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
	Uid:                 "uid",
	Mobile:              "mobile",
	Status:              "status",
	KycStatus:           "kyc_status",
	UserCardStatus:      "user_card_status",
	BankCardStatus:      "bank_card_status",
	BankImage:           "bank_image",
	CardLevelId:         "card_level_id",
	UserCardImage:       "user_card_image",
	Remark:              "remark",
	Email:               "email",
}

// NewAppUserDao creates and returns a new DAO object for table data access.
func NewAppUserDao() *AppUserDao {
	return &AppUserDao{
		group:   "default",
		table:   "app_user",
		columns: appUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppUserDao) Columns() AppUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
