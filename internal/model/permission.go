package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type UserPermission struct {
	Id        int         `json:"id"         orm:"id"         description:""`                    //
	Pid       int         `json:"pid"        orm:"pid"        description:"功能父id"`               // 功能父id
	Name      string      `json:"name"       orm:"name"       description:"规则名称"`                // 规则名称
	MenuType  int         `json:"menu_type"  orm:"menu_type"  description:"菜单类型，1-目录 2-菜单 3-按钮"` // 菜单类型，1-目录 2-菜单 3-按钮
	Path      string      `json:"path"       orm:"path"       description:"路由地址"`                // 路由地址
	Component string      `json:"component"  orm:"component"  description:"组件路径"`                // 组件路径
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`                // 创建时间
	UpdateAt  *gtime.Time `json:"update_at"  orm:"update_at"  description:"修改时间"`                // 修改时间
	DeleteAt  *gtime.Time `json:"delete_at"  orm:"delete_at"  description:"删除时间"`                // 删除时间
	Meta      Meta        `json:"meta"`
}

type Meta struct {
	Title    string `json:"title"     description:"功能名称"`
	Icon     string `json:"icon"      description:"icon图标"`
	IsHide   int    `json:"is_hide"    description:"显示状态，1-隐藏 2-显示"`
	IsCached int    `json:"is_cached"  description:"是否缓存，1-不缓存，2-缓存"`
}

type UserPermissionTree struct {
	*UserPermission
	Children []*UserPermissionTree `json:"children"`
}

type UserInfo struct {
	Id       int    `json:"id"       orm:"id"       description:""`
	Name     string `json:"name"       orm:"name"       description:"用户名"`    // 用户名
	SystemId int    `json:"system_id"  orm:"system_id"  description:"所属后台id"` // 所属后台id
}
