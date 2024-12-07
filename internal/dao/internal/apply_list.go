// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyListDao is the data access object for table apply_list.
type ApplyListDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ApplyListColumns // columns contains all the column names of Table for convenient usage.
}

// ApplyListColumns defines and stores column names for table apply_list.
type ApplyListColumns struct {
	Id            string //
	ApplyId       string // 申请单ID
	CardNo        string // 卡号
	AppUserId     string // 平台用户ID
	Status        string // 1-申请中 2-审核通过 3-已发卡 4-已激活 5-已实名 6-已驳回 7-已取消 8-供应商异常
	TotalQuota    string // 总额度
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	UserCardImage string // 后台上传的卡片照片
	UserName      string //
	PhoneNumber   string // User Phone Number
	BankUid       string // bank_uid
	Balance       string // 余额信息
	Note          string // 备注信息
}

// applyListColumns holds the columns for table apply_list.
var applyListColumns = ApplyListColumns{
	Id:            "id",
	ApplyId:       "apply_id",
	CardNo:        "card_no",
	AppUserId:     "app_user_id",
	Status:        "status",
	TotalQuota:    "total_quota",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	UserCardImage: "user_card_image",
	UserName:      "user_name",
	PhoneNumber:   "phone_number",
	BankUid:       "bank_uid",
	Balance:       "balance",
	Note:          "note",
}

// NewApplyListDao creates and returns a new DAO object for table data access.
func NewApplyListDao() *ApplyListDao {
	return &ApplyListDao{
		group:   "default",
		table:   "apply_list",
		columns: applyListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApplyListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApplyListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApplyListDao) Columns() ApplyListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApplyListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApplyListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApplyListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
