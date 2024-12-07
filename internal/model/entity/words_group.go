// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WordsGroup is the golang structure for table words_group.
type WordsGroup struct {
	Id           uint        `json:"id"            orm:"id"            description:""`                                        //
	Name         string      `json:"name"          orm:"name"          description:"分组名称"`                                    // 分组名称
	PlatformType string      `json:"platform_type" orm:"platform_type" description:"(对应字典platform_type)多个逗号隔开 使用平台1Facebook"` // (对应字典platform_type)多个逗号隔开 使用平台1Facebook
	WordsType    string      `json:"words_type"    orm:"words_type"    description:"(对应字典words_type)多个逗号隔开 话术类型1养号2批量私信"`     // (对应字典words_type)多个逗号隔开 话术类型1养号2批量私信
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`                                    // 创建时间
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:"修改时间"`                                    // 修改时间
	ClientId     int         `json:"client_id"     orm:"client_id"     description:"客户ID"`                                    // 客户ID
}
