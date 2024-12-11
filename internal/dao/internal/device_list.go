// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceListDao is the data access object for table device_list.
type DeviceListDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns DeviceListColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceListColumns defines and stores column names for table device_list.
type DeviceListColumns struct {
	Id              string //
	DeviceId        string // 设备id
	Number          string // 设备号码
	ActiveTime      string // 激活时间
	OwnerAccount    string // 子账号
	AssignedItems   string // 所属项目
	SentStatus      string // 发送状态，1-异常 2-占用 3-空闲
	QuantitySent    string // 发送数量
	QuantityAll     string // 全部数量
	DeviceStatus    string // 设备状态，1-异常 2-正常
	GroupName       string // 分组名称
	GroupId         string // 分组id
	CreatedAt       string // 创建时间
	UpdateAt        string // 修改时间
	DeleteAt        string // 删除时间
	DeviceNumber    string // 设备序列号
	StartAt         string // 发送时间
	AssignedItemsId string // 项目id
	OwnerAccountId  string // 子账号id
}

// deviceListColumns holds the columns for table device_list.
var deviceListColumns = DeviceListColumns{
	Id:              "id",
	DeviceId:        "device_id",
	Number:          "number",
	ActiveTime:      "active_time",
	OwnerAccount:    "owner_account",
	AssignedItems:   "assigned_items",
	SentStatus:      "sent_status",
	QuantitySent:    "quantity_sent",
	QuantityAll:     "quantity_all",
	DeviceStatus:    "device_status",
	GroupName:       "group_name",
	GroupId:         "group_id",
	CreatedAt:       "created_at",
	UpdateAt:        "update_at",
	DeleteAt:        "delete_at",
	DeviceNumber:    "device_number",
	StartAt:         "start_at",
	AssignedItemsId: "assigned_items_id",
	OwnerAccountId:  "owner_account_id",
}

// NewDeviceListDao creates and returns a new DAO object for table data access.
func NewDeviceListDao() *DeviceListDao {
	return &DeviceListDao{
		group:   "default",
		table:   "device_list",
		columns: deviceListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceListDao) Columns() DeviceListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
