// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RaiseTaskDao is the data access object for table raise_task.
type RaiseTaskDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns RaiseTaskColumns // columns contains all the column names of Table for convenient usage.
}

// RaiseTaskColumns defines and stores column names for table raise_task.
type RaiseTaskColumns struct {
	Id                 string //
	Name               string // 任务名称
	BrowseContentType  string // 浏览内容类型,用
	RunDuration        string // 运行时间（分钟）
	LikePercent        string // 点赞比例（%）
	FollowPercent      string // 关注比例（%）
	MinBrowseTime      string // 最小浏览时长（秒）
	MaxBrowseTime      string // 最长浏览时长
	CommentLikePercent string // 评论区点赞率（%）
	IfComment          string // 是否评论1是2否
	CommentMethod      string // 评论方式1话术2自定义3话术+自定义
	CommentWordsId     string // 评论使用的话术id
	CustomComment      string // 自定义评论内容
	CommentPercent     string // 评论比例（%）
	IfSearchKey        string // 是否搜索关键词
	Keywords           string // 关键词
	IfCron             string // 是否开启定时任务1开启2关闭
	CronMode           string // 定时模式1每日定时运行（单次）2循环运行
	CronExecTime       string // 定时任务执行时间(单次）
	CronStartTime      string // 定时任务开始时间
	CronEndTime        string // 定时任务结束时间
	CronGapTime        string // 定时任务间隔时间
	CreatedAt          string // 创建时间
	UpdatedAt          string // 修改时间
	DeletedAt          string // 删除时间
	ClientId           string // 创建客户id
}

// raiseTaskColumns holds the columns for table raise_task.
var raiseTaskColumns = RaiseTaskColumns{
	Id:                 "id",
	Name:               "name",
	BrowseContentType:  "browse_content_type",
	RunDuration:        "run_duration",
	LikePercent:        "like_percent",
	FollowPercent:      "follow_percent",
	MinBrowseTime:      "min_browse_time",
	MaxBrowseTime:      "max_browse_time",
	CommentLikePercent: "comment_like_percent",
	IfComment:          "if_comment",
	CommentMethod:      "comment_method",
	CommentWordsId:     "comment_words_id",
	CustomComment:      "custom_comment",
	CommentPercent:     "comment_percent",
	IfSearchKey:        "if_search_key",
	Keywords:           "keywords",
	IfCron:             "if_cron",
	CronMode:           "cron_mode",
	CronExecTime:       "cron_exec_time",
	CronStartTime:      "cron_start_time",
	CronEndTime:        "cron_end_time",
	CronGapTime:        "cron_gap_time",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
	ClientId:           "client_id",
}

// NewRaiseTaskDao creates and returns a new DAO object for table data access.
func NewRaiseTaskDao() *RaiseTaskDao {
	return &RaiseTaskDao{
		group:   "default",
		table:   "raise_task",
		columns: raiseTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RaiseTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RaiseTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RaiseTaskDao) Columns() RaiseTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RaiseTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RaiseTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RaiseTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
