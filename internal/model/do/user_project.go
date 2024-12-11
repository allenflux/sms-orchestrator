// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProject is the golang structure of table user_project for DAO operations like Where/Data.
type UserProject struct {
	g.Meta    `orm:"table:user_project, do:true"`
	Id        interface{} //
	UserId    interface{} // 用户id
	ProjectId interface{} // 项目id
	CreatedAt *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 修改时间
	DeleteAt  *gtime.Time // 删除时间
}
