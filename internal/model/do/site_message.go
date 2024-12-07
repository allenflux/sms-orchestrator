// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteMessage is the golang structure of table site_message for DAO operations like Where/Data.
type SiteMessage struct {
	g.Meta        `orm:"table:site_message, do:true"`
	Id            interface{} //
	Title         interface{} // 标题
	Content       interface{} // 消息
	AppUserIdStrs interface{} // APP用户ID，逗号隔开
	AdminUsername interface{} // 管理员账号
	HasAll        interface{} // 0-非所有  1-所有
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
