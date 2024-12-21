// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure of table permission for DAO operations like Where/Data.
type Permission struct {
	g.Meta    `orm:"table:permission, do:true"`
	Id        interface{} //
	Pid       interface{} // 功能父id
	Name      interface{} // 规则名称
	Title     interface{} // 功能名称
	Icon      interface{} // icon图标
	MenuType  interface{} // 菜单类型，1-目录 2-菜单 3-按钮
	IsHide    interface{} // 显示状态，1-隐藏 2-显示
	Path      interface{} // 路由地址
	Component interface{} // 组件路径
	IsCached  interface{} // 是否缓存，1-不缓存，2-缓存
	CreatedAt *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 修改时间
	DeleteAt  *gtime.Time // 删除时间
}
