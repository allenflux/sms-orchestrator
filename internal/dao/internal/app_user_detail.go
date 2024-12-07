// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppUserDetailDao is the data access object for table app_user_detail.
type AppUserDetailDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AppUserDetailColumns // columns contains all the column names of Table for convenient usage.
}

// AppUserDetailColumns defines and stores column names for table app_user_detail.
type AppUserDetailColumns struct {
	Id                string //
	AppUserId         string // 关联app_user表的id
	FirstnameCn       string // 中文姓
	LastnameCn        string // 中文名
	FirstnamePinyin   string // 拼音姓
	LastnamePinyin    string // 拼音名
	Birthday          string // 出生日期
	AnnualIncome      string // 年收入
	Career            string // 职业
	Position          string // 职位
	Region            string // 地区
	Addr              string // 地址
	IdCardNo          string // 身份证号
	IdCardExpiredAt   string // 证件有效期
	LongTermEffective string // 证件是否长期有效1代表是2代表不是
	IdCardFrontPic    string // 身份证正面照片
	IdCardBackPic     string // 身份证背面照片
	HasCommit         string // 是否提交到万事达,1否2是
	CreatedAt         string // 创建时间
	UpdatedAt         string // 修改时间
	DeletedAt         string // 删除时间
}

// appUserDetailColumns holds the columns for table app_user_detail.
var appUserDetailColumns = AppUserDetailColumns{
	Id:                "id",
	AppUserId:         "app_user_id",
	FirstnameCn:       "firstname_cn",
	LastnameCn:        "lastname_cn",
	FirstnamePinyin:   "firstname_pinyin",
	LastnamePinyin:    "lastname_pinyin",
	Birthday:          "birthday",
	AnnualIncome:      "annual_income",
	Career:            "career",
	Position:          "position",
	Region:            "region",
	Addr:              "addr",
	IdCardNo:          "id_card_no",
	IdCardExpiredAt:   "id_card_expired_at",
	LongTermEffective: "long_term_effective",
	IdCardFrontPic:    "id_card_front_pic",
	IdCardBackPic:     "id_card_back_pic",
	HasCommit:         "has_commit",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
}

// NewAppUserDetailDao creates and returns a new DAO object for table data access.
func NewAppUserDetailDao() *AppUserDetailDao {
	return &AppUserDetailDao{
		group:   "default",
		table:   "app_user_detail",
		columns: appUserDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppUserDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppUserDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppUserDetailDao) Columns() AppUserDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppUserDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppUserDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppUserDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
