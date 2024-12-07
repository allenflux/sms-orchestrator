// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RaiseTask is the golang structure for table raise_task.
type RaiseTask struct {
	Id                 uint        `json:"id"                   orm:"id"                   description:""`                     //
	Name               string      `json:"name"                 orm:"name"                 description:"任务名称"`                 // 任务名称
	BrowseContentType  string      `json:"browse_content_type"  orm:"browse_content_type"  description:"浏览内容类型,用"`             // 浏览内容类型,用
	RunDuration        int         `json:"run_duration"         orm:"run_duration"         description:"运行时间（分钟）"`             // 运行时间（分钟）
	LikePercent        int         `json:"like_percent"         orm:"like_percent"         description:"点赞比例（%）"`              // 点赞比例（%）
	FollowPercent      int         `json:"follow_percent"       orm:"follow_percent"       description:"关注比例（%）"`              // 关注比例（%）
	MinBrowseTime      int         `json:"min_browse_time"      orm:"min_browse_time"      description:"最小浏览时长（秒）"`            // 最小浏览时长（秒）
	MaxBrowseTime      int         `json:"max_browse_time"      orm:"max_browse_time"      description:"最长浏览时长"`               // 最长浏览时长
	CommentLikePercent int         `json:"comment_like_percent" orm:"comment_like_percent" description:"评论区点赞率（%）"`            // 评论区点赞率（%）
	IfComment          int         `json:"if_comment"           orm:"if_comment"           description:"是否评论1是2否"`             // 是否评论1是2否
	CommentMethod      int         `json:"comment_method"       orm:"comment_method"       description:"评论方式1话术2自定义3话术+自定义"`   // 评论方式1话术2自定义3话术+自定义
	CommentWordsId     int         `json:"comment_words_id"     orm:"comment_words_id"     description:"评论使用的话术id"`            // 评论使用的话术id
	CustomComment      string      `json:"custom_comment"       orm:"custom_comment"       description:"自定义评论内容"`              // 自定义评论内容
	CommentPercent     int         `json:"comment_percent"      orm:"comment_percent"      description:"评论比例（%）"`              // 评论比例（%）
	IfSearchKey        int         `json:"if_search_key"        orm:"if_search_key"        description:"是否搜索关键词"`              // 是否搜索关键词
	Keywords           string      `json:"keywords"             orm:"keywords"             description:"关键词"`                  // 关键词
	IfCron             int         `json:"if_cron"              orm:"if_cron"              description:"是否开启定时任务1开启2关闭"`       // 是否开启定时任务1开启2关闭
	CronMode           int         `json:"cron_mode"            orm:"cron_mode"            description:"定时模式1每日定时运行（单次）2循环运行"` // 定时模式1每日定时运行（单次）2循环运行
	CronExecTime       *gtime.Time `json:"cron_exec_time"       orm:"cron_exec_time"       description:"定时任务执行时间(单次）"`         // 定时任务执行时间(单次）
	CronStartTime      *gtime.Time `json:"cron_start_time"      orm:"cron_start_time"      description:"定时任务开始时间"`             // 定时任务开始时间
	CronEndTime        *gtime.Time `json:"cron_end_time"        orm:"cron_end_time"        description:"定时任务结束时间"`             // 定时任务结束时间
	CronGapTime        *gtime.Time `json:"cron_gap_time"        orm:"cron_gap_time"        description:"定时任务间隔时间"`             // 定时任务间隔时间
	CreatedAt          *gtime.Time `json:"created_at"           orm:"created_at"           description:"创建时间"`                 // 创建时间
	UpdatedAt          *gtime.Time `json:"updated_at"           orm:"updated_at"           description:"修改时间"`                 // 修改时间
	DeletedAt          *gtime.Time `json:"deleted_at"           orm:"deleted_at"           description:"删除时间"`                 // 删除时间
	ClientId           int         `json:"client_id"            orm:"client_id"            description:"创建客户id"`               // 创建客户id
}
