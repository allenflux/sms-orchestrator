// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Content is the golang structure of table content for DAO operations like Where/Data.
type Content struct {
	g.Meta    `orm:"table:content, do:true"`
	Id        interface{} //
	Link      interface{} // 跳转链接
	Title     interface{} // 标题
	Content   interface{} // 内容
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
