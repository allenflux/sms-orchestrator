// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectDataDao is the data access object for table collect_data.
type CollectDataDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CollectDataColumns // columns contains all the column names of Table for convenient usage.
}

// CollectDataColumns defines and stores column names for table collect_data.
type CollectDataColumns struct {
	Id                 string //
	UserAvatar         string // 用户头像
	UserNickname       string // 用户昵称
	UserUid            string // 用户uid
	UserCountry        string // 用户国家
	UserGender         string // 用户性别1男2女3未知
	UserAge            string // 用户年龄
	UserHomePage       string // 用户个人主页
	CollectUrl         string // 采集地址
	CollectKeyword     string // 采集关键词
	CollectDataGroupId string // 采集数据分组
	CreatedAt          string // 采集时间
	UseStatus          string // 使用状态1待使用2已使用
	UseNum             string // 使用次数
	LastUsedAt         string // 上次使用时间
	DataStatus         string // 数据状态1正常2异常
	CollectDevId       string // id
	CollectTaskId      string // 采集任务id
	UpdatedAt          string // 修改时间
	DeletedAt          string // 删除时间
	CreatedBy          string // 创建用户id
}

// collectDataColumns holds the columns for table collect_data.
var collectDataColumns = CollectDataColumns{
	Id:                 "id",
	UserAvatar:         "user_avatar",
	UserNickname:       "user_nickname",
	UserUid:            "user_uid",
	UserCountry:        "user_country",
	UserGender:         "user_gender",
	UserAge:            "user_age",
	UserHomePage:       "user_home_page",
	CollectUrl:         "collect_url",
	CollectKeyword:     "collect_keyword",
	CollectDataGroupId: "collect_data_group_id",
	CreatedAt:          "created_at",
	UseStatus:          "use_status",
	UseNum:             "use_num",
	LastUsedAt:         "last_used_at",
	DataStatus:         "data_status",
	CollectDevId:       "collect_dev_id",
	CollectTaskId:      "collect_task_id",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
	CreatedBy:          "created_by",
}

// NewCollectDataDao creates and returns a new DAO object for table data access.
func NewCollectDataDao() *CollectDataDao {
	return &CollectDataDao{
		group:   "default",
		table:   "collect_data",
		columns: collectDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CollectDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CollectDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CollectDataDao) Columns() CollectDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CollectDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CollectDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CollectDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
