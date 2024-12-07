// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatTaskDao is the data access object for table chat_task.
type ChatTaskDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ChatTaskColumns // columns contains all the column names of Table for convenient usage.
}

// ChatTaskColumns defines and stores column names for table chat_task.
type ChatTaskColumns struct {
	Id              string //
	Name            string // 任务名称
	ChatNum         string // 私信数量
	ChatContentType string // 私信内容类型1话术+自定义2话术3自定义
	WordsGroupId    string // 话术组id
	CustomContent   string // 自定义内容
	SendGap         string // 发生时间间隔秒数
	CreatedAt       string // 创建时间
	CompletedAt     string // 完成时间
	Status          string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	Remark          string // 备注
	ClientId        string // 用户ID
	HasEmpty        string // 1-未清空 2-已清空
}

// chatTaskColumns holds the columns for table chat_task.
var chatTaskColumns = ChatTaskColumns{
	Id:              "id",
	Name:            "name",
	ChatNum:         "chat_num",
	ChatContentType: "chat_content_type",
	WordsGroupId:    "words_group_id",
	CustomContent:   "custom_content",
	SendGap:         "send_gap",
	CreatedAt:       "created_at",
	CompletedAt:     "completed_at",
	Status:          "status",
	Remark:          "remark",
	ClientId:        "client_id",
	HasEmpty:        "has_empty",
}

// NewChatTaskDao creates and returns a new DAO object for table data access.
func NewChatTaskDao() *ChatTaskDao {
	return &ChatTaskDao{
		group:   "default",
		table:   "chat_task",
		columns: chatTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChatTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChatTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChatTaskDao) Columns() ChatTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChatTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChatTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChatTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
