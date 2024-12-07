// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatTaskLinkCollectDataGroupDao is the data access object for table chat_task_link_collect_data_group.
type ChatTaskLinkCollectDataGroupDao struct {
	table   string                              // table is the underlying table name of the DAO.
	group   string                              // group is the database configuration group name of current DAO.
	columns ChatTaskLinkCollectDataGroupColumns // columns contains all the column names of Table for convenient usage.
}

// ChatTaskLinkCollectDataGroupColumns defines and stores column names for table chat_task_link_collect_data_group.
type ChatTaskLinkCollectDataGroupColumns struct {
	Id                 string //
	ChatTaskId         string // 私信任务id
	CollectDataGroupId string // 采集数据组id
}

// chatTaskLinkCollectDataGroupColumns holds the columns for table chat_task_link_collect_data_group.
var chatTaskLinkCollectDataGroupColumns = ChatTaskLinkCollectDataGroupColumns{
	Id:                 "id",
	ChatTaskId:         "chat_task_id",
	CollectDataGroupId: "collect_data_group_id",
}

// NewChatTaskLinkCollectDataGroupDao creates and returns a new DAO object for table data access.
func NewChatTaskLinkCollectDataGroupDao() *ChatTaskLinkCollectDataGroupDao {
	return &ChatTaskLinkCollectDataGroupDao{
		group:   "default",
		table:   "chat_task_link_collect_data_group",
		columns: chatTaskLinkCollectDataGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChatTaskLinkCollectDataGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChatTaskLinkCollectDataGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChatTaskLinkCollectDataGroupDao) Columns() ChatTaskLinkCollectDataGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChatTaskLinkCollectDataGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChatTaskLinkCollectDataGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChatTaskLinkCollectDataGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
