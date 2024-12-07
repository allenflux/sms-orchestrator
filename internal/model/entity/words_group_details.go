// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WordsGroupDetails is the golang structure for table words_group_details.
type WordsGroupDetails struct {
	Id           int         `json:"id"             orm:"id"             description:""`                //
	WordsId      int         `json:"words_id"       orm:"words_id"       description:"话术ID"`            // 话术ID
	PlatformType string      `json:"platform_type"  orm:"platform_type"  description:"平台类型 1-facebook"` // 平台类型 1-facebook
	WordsType    string      `json:"words_type"     orm:"words_type"     description:"话术类型 1养号2批量私信"`   // 话术类型 1养号2批量私信
	CreatedAt    *gtime.Time `json:"created_at"     orm:"created_at"     description:""`                //
	UpdatedAt    *gtime.Time `json:"updated_at"     orm:"updated_at"     description:""`                //
	WordsGroupId int         `json:"words_group_id" orm:"words_group_id" description:"话术分组ID"`          // 话术分组ID
}
