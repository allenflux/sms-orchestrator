// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} //
	Name      interface{} // 用户名
	Password  interface{} // 用户密码
	Salt      interface{} // 密码盐
	Status    interface{} // 账号状态，1-停用 2-启用
	RoleId    interface{} // 角色权限id
	SystemId  interface{} // 所属后台id
	Note      interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeleteAt  *gtime.Time // 删除时间
}
