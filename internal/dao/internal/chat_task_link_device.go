// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatTaskLinkDeviceDao is the data access object for table chat_task_link_device.
type ChatTaskLinkDeviceDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ChatTaskLinkDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// ChatTaskLinkDeviceColumns defines and stores column names for table chat_task_link_device.
type ChatTaskLinkDeviceColumns struct {
	Id            string //
	ChatTaskId    string // 私信任务id
	DeviceId      string // 设备id
	DeviceDevId   string // id
	DeviceGroupId string // 设备组id
	CreatedAt     string // 创建时间
	CompletedAt   string // 完成时间
	Status        string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	FailedNum     string // 失败数量
	SuccessedNum  string // 成功数量
	TotalNum      string // 总的私信数量
	Remark        string // 备注
}

// chatTaskLinkDeviceColumns holds the columns for table chat_task_link_device.
var chatTaskLinkDeviceColumns = ChatTaskLinkDeviceColumns{
	Id:            "id",
	ChatTaskId:    "chat_task_id",
	DeviceId:      "device_id",
	DeviceDevId:   "device_dev_id",
	DeviceGroupId: "device_group_id",
	CreatedAt:     "created_at",
	CompletedAt:   "completed_at",
	Status:        "status",
	FailedNum:     "failed_num",
	SuccessedNum:  "successed_num",
	TotalNum:      "total_num",
	Remark:        "remark",
}

// NewChatTaskLinkDeviceDao creates and returns a new DAO object for table data access.
func NewChatTaskLinkDeviceDao() *ChatTaskLinkDeviceDao {
	return &ChatTaskLinkDeviceDao{
		group:   "default",
		table:   "chat_task_link_device",
		columns: chatTaskLinkDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChatTaskLinkDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChatTaskLinkDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChatTaskLinkDeviceDao) Columns() ChatTaskLinkDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChatTaskLinkDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChatTaskLinkDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChatTaskLinkDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
