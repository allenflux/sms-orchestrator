// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MinerImportPhotoTaskLinkJdGroupDao is the data access object for table miner_import_photo_task_link_jd_group.
type MinerImportPhotoTaskLinkJdGroupDao struct {
	table   string                                 // table is the underlying table name of the DAO.
	group   string                                 // group is the database configuration group name of current DAO.
	columns MinerImportPhotoTaskLinkJdGroupColumns // columns contains all the column names of Table for convenient usage.
}

// MinerImportPhotoTaskLinkJdGroupColumns defines and stores column names for table miner_import_photo_task_link_jd_group.
type MinerImportPhotoTaskLinkJdGroupColumns struct {
	Id                     string //
	MinerImportPhotoTaskId string // 导入相册任务ID
	JdGroupName            string // 金盾分组名称
	Num                    string // 分组数量
	CreatedAt              string //
	MinerTaskType          string // 1-导入相册 2-导入通讯录 3-矿机登录 4-日常任务 5-定时任务计划
}

// minerImportPhotoTaskLinkJdGroupColumns holds the columns for table miner_import_photo_task_link_jd_group.
var minerImportPhotoTaskLinkJdGroupColumns = MinerImportPhotoTaskLinkJdGroupColumns{
	Id:                     "id",
	MinerImportPhotoTaskId: "miner_import_photo_task_id",
	JdGroupName:            "jd_group_name",
	Num:                    "num",
	CreatedAt:              "created_at",
	MinerTaskType:          "miner_task_type",
}

// NewMinerImportPhotoTaskLinkJdGroupDao creates and returns a new DAO object for table data access.
func NewMinerImportPhotoTaskLinkJdGroupDao() *MinerImportPhotoTaskLinkJdGroupDao {
	return &MinerImportPhotoTaskLinkJdGroupDao{
		group:   "default",
		table:   "miner_import_photo_task_link_jd_group",
		columns: minerImportPhotoTaskLinkJdGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MinerImportPhotoTaskLinkJdGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MinerImportPhotoTaskLinkJdGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MinerImportPhotoTaskLinkJdGroupDao) Columns() MinerImportPhotoTaskLinkJdGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MinerImportPhotoTaskLinkJdGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MinerImportPhotoTaskLinkJdGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MinerImportPhotoTaskLinkJdGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
