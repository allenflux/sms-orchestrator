// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteMessageRead is the golang structure for table site_message_read.
type SiteMessageRead struct {
	Id        int         `json:"id"          orm:"id"          description:""`        //
	AppUserId int         `json:"app_user_id" orm:"app_user_id" description:"APP用户ID"` // APP用户ID
	MessageId int         `json:"message_id"  orm:"message_id"  description:"站内信主键ID"` // 站内信主键ID
	CreatedAt *gtime.Time `json:"created_at"  orm:"created_at"  description:""`        //
	UpdatedAt *gtime.Time `json:"updated_at"  orm:"updated_at"  description:""`        //
}
