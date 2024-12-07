// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectTask is the golang structure for table collect_task.
type CollectTask struct {
	Id                     uint        `json:"id"                       orm:"id"                       description:""`                                  //
	Name                   string      `json:"name"                     orm:"name"                     description:"采集任务名称"`                            // 采集任务名称
	TaskType               int         `json:"task_type"                orm:"task_type"                description:"任务类型1获客采集"`                         // 任务类型1获客采集
	CollectType            int         `json:"collect_type"             orm:"collect_type"             description:"采集类型1Video2群组3帖子4公共主页"`             // 采集类型1Video2群组3帖子4公共主页
	GroupCity              string      `json:"group_city"               orm:"group_city"               description:"采集群组城市"`                            // 采集群组城市
	GroupNearby            int         `json:"group_nearby"             orm:"group_nearby"             description:"采集群组筛选条件-我附近，1true2false"`          // 采集群组筛选条件-我附近，1true2false
	GroupPublic            int         `json:"group_public"             orm:"group_public"             description:"采集群组筛选条件-公开小组，1true2false"`         // 采集群组筛选条件-公开小组，1true2false
	GroupMine              int         `json:"group_mine"               orm:"group_mine"               description:"采集群组筛选条件-我的小组，1true2false"`         // 采集群组筛选条件-我的小组，1true2false
	GroupObjectLocation    string      `json:"group_object_location"    orm:"group_object_location"    description:"采集群组筛选条件-对象所在地"`                    // 采集群组筛选条件-对象所在地
	GroupObjectResidential string      `json:"group_object_residential" orm:"group_object_residential" description:"采集群组筛选条件-对象居住地"`                    // 采集群组筛选条件-对象居住地
	CollectObject          int         `json:"collect_object"           orm:"collect_object"           description:"采集对象1分享用户2点赞用户3成员"`                 // 采集对象1分享用户2点赞用户3成员
	CollectMethod          int         `json:"collect_method"           orm:"collect_method"           description:"采集方式1链接采集2关键词采集"`                   // 采集方式1链接采集2关键词采集
	CollectNum             int         `json:"collect_num"              orm:"collect_num"              description:"采集数量，-1代表全部"`                       // 采集数量，-1代表全部
	CollectKeyword         string      `json:"collect_keyword"          orm:"collect_keyword"          description:"采集关键词"`                             // 采集关键词
	Url                    string      `json:"url"                      orm:"url"                      description:"采集页面链接"`                            // 采集页面链接
	CollectDataGroupId     int         `json:"collect_data_group_id"    orm:"collect_data_group_id"    description:"采集数据分组id"`                          // 采集数据分组id
	Status                 int         `json:"status"                   orm:"status"                   description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	CreatedAt              *gtime.Time `json:"created_at"               orm:"created_at"               description:"创建时间"`                              // 创建时间
	UpdatedAt              *gtime.Time `json:"updated_at"               orm:"updated_at"               description:"修改时间"`                              // 修改时间
	DeletedAt              *gtime.Time `json:"deleted_at"               orm:"deleted_at"               description:"删除时间"`                              // 删除时间
	CompletedAt            *gtime.Time `json:"completed_at"             orm:"completed_at"             description:"完成时间"`                              // 完成时间
	Remark                 string      `json:"remark"                   orm:"remark"                   description:"备注"`                                // 备注
	ClientId               int         `json:"client_id"                orm:"client_id"                description:"用户ID"`                              // 用户ID
	HasEmpty               int         `json:"has_empty"                orm:"has_empty"                description:"1-未清空 2-清空"`                        // 1-未清空 2-清空
}
