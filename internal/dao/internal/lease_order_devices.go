// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LeaseOrderDevicesDao is the data access object for table lease_order_devices.
type LeaseOrderDevicesDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns LeaseOrderDevicesColumns // columns contains all the column names of Table for convenient usage.
}

// LeaseOrderDevicesColumns defines and stores column names for table lease_order_devices.
type LeaseOrderDevicesColumns struct {
	Id                     string //
	LeaseOrderId           string // （租赁/续费）订单ID
	DeviceId               string // 设备主键ID
	ClientId               string // 客户ID
	OrderNo                string // 订单号
	CreatedAt              string // 创建时间
	UpdatedAt              string // 更新时间
	LeaseAt                string // 租赁时间
	LeaseEndAt             string // 到期时间
	UnbindAt               string // 解绑时间
	Remark                 string // 备注
	LeaseDays              string // 租赁天数
	BindSpecialOperateId   string // 特殊操作表ID（绑定）
	UnbindSpecialOperateId string // 特殊操作表ID（解绑）
	DelaySpecialOperateId  string // 特殊操作表ID（延期）
	ProduceWay             string // 1-租赁  2-续费 3-批量绑定
}

// leaseOrderDevicesColumns holds the columns for table lease_order_devices.
var leaseOrderDevicesColumns = LeaseOrderDevicesColumns{
	Id:                     "id",
	LeaseOrderId:           "lease_order_id",
	DeviceId:               "device_id",
	ClientId:               "client_id",
	OrderNo:                "order_no",
	CreatedAt:              "created_at",
	UpdatedAt:              "updated_at",
	LeaseAt:                "lease_at",
	LeaseEndAt:             "lease_end_at",
	UnbindAt:               "unbind_at",
	Remark:                 "remark",
	LeaseDays:              "lease_days",
	BindSpecialOperateId:   "bind_special_operate_id",
	UnbindSpecialOperateId: "unbind_special_operate_id",
	DelaySpecialOperateId:  "delay_special_operate_id",
	ProduceWay:             "produce_way",
}

// NewLeaseOrderDevicesDao creates and returns a new DAO object for table data access.
func NewLeaseOrderDevicesDao() *LeaseOrderDevicesDao {
	return &LeaseOrderDevicesDao{
		group:   "default",
		table:   "lease_order_devices",
		columns: leaseOrderDevicesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LeaseOrderDevicesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LeaseOrderDevicesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LeaseOrderDevicesDao) Columns() LeaseOrderDevicesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LeaseOrderDevicesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LeaseOrderDevicesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LeaseOrderDevicesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
