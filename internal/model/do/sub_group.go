// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SubGroup is the golang structure of table sub_group for DAO operations like Where/Data.
type SubGroup struct {
	g.Meta       `orm:"table:sub_group, do:true"`
	Id           interface{} //
	SubUserId    interface{} // 子账号id
	ProjectId    interface{} // 项目id
	SubGroupName interface{} // 分组名称
	CreatedAt    *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 修改时间
	DeleteAt     *gtime.Time // 删除时间
}
