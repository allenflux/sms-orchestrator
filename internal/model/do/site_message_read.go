// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteMessageRead is the golang structure of table site_message_read for DAO operations like Where/Data.
type SiteMessageRead struct {
	g.Meta    `orm:"table:site_message_read, do:true"`
	Id        interface{} //
	AppUserId interface{} // APP用户ID
	MessageId interface{} // 站内信主键ID
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
