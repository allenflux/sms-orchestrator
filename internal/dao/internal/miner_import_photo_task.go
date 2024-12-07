// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerImportPhotoTaskDao is the data access object for table miner_import_photo_task.
type MinerImportPhotoTaskDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns MinerImportPhotoTaskColumns // columns contains all the column names of Table for convenient usage.
}

// MinerImportPhotoTaskColumns defines and stores column names for table miner_import_photo_task.
type MinerImportPhotoTaskColumns struct {
	Id        string //
	Name      string // 任务名称
	PhotoNum  string // 图片数量
	Status    string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CreatedAt string //
	ExecAt    string // 执行时间
	Remark    string // 备注
	ClientId  string // 客户ID
	ExecEndAt string // 执行结束时间
	DeviceNum string // 设备数量
	HasShow   string // 1-显示 2-不显示
	Type      string // 1-图片 2-视频
}

// minerImportPhotoTaskColumns holds the columns for table miner_import_photo_task.
var minerImportPhotoTaskColumns = MinerImportPhotoTaskColumns{
	Id:        "id",
	Name:      "name",
	PhotoNum:  "photo_num",
	Status:    "status",
	CreatedAt: "created_at",
	ExecAt:    "exec_at",
	Remark:    "remark",
	ClientId:  "client_id",
	ExecEndAt: "exec_end_at",
	DeviceNum: "device_num",
	HasShow:   "has_show",
	Type:      "type",
}

// NewMinerImportPhotoTaskDao creates and returns a new DAO object for table data access.
func NewMinerImportPhotoTaskDao() *MinerImportPhotoTaskDao {
	return &MinerImportPhotoTaskDao{
		group:   "default",
		table:   "miner_import_photo_task",
		columns: minerImportPhotoTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerImportPhotoTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerImportPhotoTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerImportPhotoTaskDao) Columns() MinerImportPhotoTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerImportPhotoTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerImportPhotoTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerImportPhotoTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
