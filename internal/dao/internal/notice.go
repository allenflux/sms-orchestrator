// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeDao is the data access object for table notice.
type NoticeDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns NoticeColumns // columns contains all the column names of Table for convenient usage.
}

// NoticeColumns defines and stores column names for table notice.
type NoticeColumns struct {
	Id                  string //
	UserType            string // 推送目标用户类型1管理端用户2企业端用户
	UserId              string // 推送目标用户id
	OrderId             string // 订单id
	Type                string // 通知类型
	Title               string // 标题
	Content             string // 内容
	Read                string // 是否已读0未读1已读
	CreatedAt           string // 推送时间
	DeletedAt           string // 删除时间
	DeviceGrantImportId string // 设备授权导入任务id
	ImportType          string // 1-授权导入  2-终端设备标签导入 3-app标签导入
}

// noticeColumns holds the columns for table notice.
var noticeColumns = NoticeColumns{
	Id:                  "id",
	UserType:            "user_type",
	UserId:              "user_id",
	OrderId:             "order_id",
	Type:                "type",
	Title:               "title",
	Content:             "content",
	Read:                "read",
	CreatedAt:           "created_at",
	DeletedAt:           "deleted_at",
	DeviceGrantImportId: "device_grant_import_id",
	ImportType:          "import_type",
}

// NewNoticeDao creates and returns a new DAO object for table data access.
func NewNoticeDao() *NoticeDao {
	return &NoticeDao{
		group:   "default",
		table:   "notice",
		columns: noticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoticeDao) Columns() NoticeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
