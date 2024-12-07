// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceDao is the data access object for table device.
type DeviceDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns DeviceColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceColumns defines and stores column names for table device.
type DeviceColumns struct {
	Id             string //
	Name           string //
	Model          string //
	State          string // 1234
	GroupId        string // ID
	LeasedAt       string //
	LeaseDay       string //
	LeaseEndAt     string //
	IdleState      string // 12
	ClientId       string // id
	ConnId         string // websocketid
	FacebookConnId string // facebook的websocket连接id
	MinerConnId    string // 矿机的websocket连接id
	DevId          string // id
	LeaseOrderId   string // id
	Ip             string // ip
	Remark         string //
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
	IsNew          string // 12
	JdUserNickname string // 金盾用户昵称
}

// deviceColumns holds the columns for table device.
var deviceColumns = DeviceColumns{
	Id:             "id",
	Name:           "name",
	Model:          "model",
	State:          "state",
	GroupId:        "group_id",
	LeasedAt:       "leased_at",
	LeaseDay:       "lease_day",
	LeaseEndAt:     "lease_end_at",
	IdleState:      "idle_state",
	ClientId:       "client_id",
	ConnId:         "conn_id",
	FacebookConnId: "facebook_conn_id",
	MinerConnId:    "miner_conn_id",
	DevId:          "dev_id",
	LeaseOrderId:   "lease_order_id",
	Ip:             "ip",
	Remark:         "remark",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	IsNew:          "is_new",
	JdUserNickname: "jd_user_nickname",
}

// NewDeviceDao creates and returns a new DAO object for table data access.
func NewDeviceDao() *DeviceDao {
	return &DeviceDao{
		group:   "default",
		table:   "device",
		columns: deviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceDao) Columns() DeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
