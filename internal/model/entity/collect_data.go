// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectData is the golang structure for table collect_data.
type CollectData struct {
	Id                 uint        `json:"id"                    orm:"id"                    description:""`             //
	UserAvatar         string      `json:"user_avatar"           orm:"user_avatar"           description:"用户头像"`         // 用户头像
	UserNickname       string      `json:"user_nickname"         orm:"user_nickname"         description:"用户昵称"`         // 用户昵称
	UserUid            string      `json:"user_uid"              orm:"user_uid"              description:"用户uid"`        // 用户uid
	UserCountry        string      `json:"user_country"          orm:"user_country"          description:"用户国家"`         // 用户国家
	UserGender         int         `json:"user_gender"           orm:"user_gender"           description:"用户性别1男2女3未知"`  // 用户性别1男2女3未知
	UserAge            int         `json:"user_age"              orm:"user_age"              description:"用户年龄"`         // 用户年龄
	UserHomePage       string      `json:"user_home_page"        orm:"user_home_page"        description:"用户个人主页"`       // 用户个人主页
	CollectUrl         string      `json:"collect_url"           orm:"collect_url"           description:"采集地址"`         // 采集地址
	CollectKeyword     string      `json:"collect_keyword"       orm:"collect_keyword"       description:"采集关键词"`        // 采集关键词
	CollectDataGroupId int         `json:"collect_data_group_id" orm:"collect_data_group_id" description:"采集数据分组"`       // 采集数据分组
	CreatedAt          *gtime.Time `json:"created_at"            orm:"created_at"            description:"采集时间"`         // 采集时间
	UseStatus          int         `json:"use_status"            orm:"use_status"            description:"使用状态1待使用2已使用"` // 使用状态1待使用2已使用
	UseNum             int         `json:"use_num"               orm:"use_num"               description:"使用次数"`         // 使用次数
	LastUsedAt         *gtime.Time `json:"last_used_at"          orm:"last_used_at"          description:"上次使用时间"`       // 上次使用时间
	DataStatus         int         `json:"data_status"           orm:"data_status"           description:"数据状态1正常2异常"`   // 数据状态1正常2异常
	CollectDevId       string      `json:"collect_dev_id"        orm:"collect_dev_id"        description:"id"`           // id
	CollectTaskId      int         `json:"collect_task_id"       orm:"collect_task_id"       description:"采集任务id"`       // 采集任务id
	UpdatedAt          *gtime.Time `json:"updated_at"            orm:"updated_at"            description:"修改时间"`         // 修改时间
	DeletedAt          *gtime.Time `json:"deleted_at"            orm:"deleted_at"            description:"删除时间"`         // 删除时间
	CreatedBy          int         `json:"created_by"            orm:"created_by"            description:"创建用户id"`       // 创建用户id
}
