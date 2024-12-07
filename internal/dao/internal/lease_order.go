// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LeaseOrderDao is the data access object for table lease_order.
type LeaseOrderDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns LeaseOrderColumns // columns contains all the column names of Table for convenient usage.
}

// LeaseOrderColumns defines and stores column names for table lease_order.
type LeaseOrderColumns struct {
	Id               string //
	ClientId         string // 客户id
	LeaseDeviceModel string // 租赁设备型号
	LeaseDeviceNum   string // 租赁设备数量
	LeaseDeviceDays  string // 租赁设备天数
	LeaseType        string // 租赁类型1租赁2续费
	State            string // 订单状态1未处理2已完成3已取消4已驳回
	Remark           string // 备注
	UpdatedBy        string // 操作人
	CompletedAt      string // 完成时间
	CreatedAt        string // 下单时间
	UpdatedAt        string // 取消/驳回时间
	DeletedAt        string // 删除时间
	DeviceName       string // 设备名称
	LeaseCount       string // 租赁点数
	TotalLeaseCount  string // 订单总点数
	DevicePriceId    string // 价格表ID
	LeaseEndAt       string // 租赁到期时间
	OrderNo          string // 订单号
	Pid              string // 续费时的id
	CreatedBy        string // 操作人
	DeviceId         string // 续费设备ID
}

// leaseOrderColumns holds the columns for table lease_order.
var leaseOrderColumns = LeaseOrderColumns{
	Id:               "id",
	ClientId:         "client_id",
	LeaseDeviceModel: "lease_device_model",
	LeaseDeviceNum:   "lease_device_num",
	LeaseDeviceDays:  "lease_device_days",
	LeaseType:        "lease_type",
	State:            "state",
	Remark:           "remark",
	UpdatedBy:        "updated_by",
	CompletedAt:      "completed_at",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	DeviceName:       "device_name",
	LeaseCount:       "lease_count",
	TotalLeaseCount:  "total_lease_count",
	DevicePriceId:    "device_price_id",
	LeaseEndAt:       "lease_end_at",
	OrderNo:          "order_no",
	Pid:              "pid",
	CreatedBy:        "created_by",
	DeviceId:         "device_id",
}

// NewLeaseOrderDao creates and returns a new DAO object for table data access.
func NewLeaseOrderDao() *LeaseOrderDao {
	return &LeaseOrderDao{
		group:   "default",
		table:   "lease_order",
		columns: leaseOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LeaseOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LeaseOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LeaseOrderDao) Columns() LeaseOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LeaseOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LeaseOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LeaseOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
