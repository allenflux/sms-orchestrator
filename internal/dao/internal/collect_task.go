// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectTaskDao is the data access object for table collect_task.
type CollectTaskDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CollectTaskColumns // columns contains all the column names of Table for convenient usage.
}

// CollectTaskColumns defines and stores column names for table collect_task.
type CollectTaskColumns struct {
	Id                     string //
	Name                   string // 采集任务名称
	TaskType               string // 任务类型1获客采集
	CollectType            string // 采集类型1Video2群组3帖子4公共主页
	GroupCity              string // 采集群组城市
	GroupNearby            string // 采集群组筛选条件-我附近，1true2false
	GroupPublic            string // 采集群组筛选条件-公开小组，1true2false
	GroupMine              string // 采集群组筛选条件-我的小组，1true2false
	GroupObjectLocation    string // 采集群组筛选条件-对象所在地
	GroupObjectResidential string // 采集群组筛选条件-对象居住地
	CollectObject          string // 采集对象1分享用户2点赞用户3成员
	CollectMethod          string // 采集方式1链接采集2关键词采集
	CollectNum             string // 采集数量，-1代表全部
	CollectKeyword         string // 采集关键词
	Url                    string // 采集页面链接
	CollectDataGroupId     string // 采集数据分组id
	Status                 string // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CreatedAt              string // 创建时间
	UpdatedAt              string // 修改时间
	DeletedAt              string // 删除时间
	CompletedAt            string // 完成时间
	Remark                 string // 备注
	ClientId               string // 用户ID
	HasEmpty               string // 1-未清空 2-清空
}

// collectTaskColumns holds the columns for table collect_task.
var collectTaskColumns = CollectTaskColumns{
	Id:                     "id",
	Name:                   "name",
	TaskType:               "task_type",
	CollectType:            "collect_type",
	GroupCity:              "group_city",
	GroupNearby:            "group_nearby",
	GroupPublic:            "group_public",
	GroupMine:              "group_mine",
	GroupObjectLocation:    "group_object_location",
	GroupObjectResidential: "group_object_residential",
	CollectObject:          "collect_object",
	CollectMethod:          "collect_method",
	CollectNum:             "collect_num",
	CollectKeyword:         "collect_keyword",
	Url:                    "url",
	CollectDataGroupId:     "collect_data_group_id",
	Status:                 "status",
	CreatedAt:              "created_at",
	UpdatedAt:              "updated_at",
	DeletedAt:              "deleted_at",
	CompletedAt:            "completed_at",
	Remark:                 "remark",
	ClientId:               "client_id",
	HasEmpty:               "has_empty",
}

// NewCollectTaskDao creates and returns a new DAO object for table data access.
func NewCollectTaskDao() *CollectTaskDao {
	return &CollectTaskDao{
		group:   "default",
		table:   "collect_task",
		columns: collectTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CollectTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CollectTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CollectTaskDao) Columns() CollectTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CollectTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CollectTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CollectTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
