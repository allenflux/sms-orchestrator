// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Log is the golang structure of table log for DAO operations like Where/Data.
type Log struct {
	g.Meta    `orm:"table:log, do:true"`
	Id        interface{} //
	UserId    interface{} // 操作者的用户id
	UserName  interface{} // 操作者的用户名
	ClientIp  interface{} // 操作ip
	Function  interface{} // 操作功能
	Note      interface{} // 操作内容
	SystemId  interface{} // 后台id
	CreatedAt *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 修改时间
	DeleteAt  *gtime.Time // 删除时间
}
