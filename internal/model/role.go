package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Role struct {
	g.Meta         `orm:"table:role"`
	Id             int               `json:"id"         orm:"id"         description:""`     //
	Name           string            `json:"name"       orm:"name"       description:"角色名称"` // 角色名称
	Note           string            `json:"note"       orm:"note"       description:"备注"`   // 备注
	UpdateAt       *gtime.Time       `json:"update_at"  orm:"update_at"  description:"修改时间"` // 修改时间
	RolePermission []*RolePermission `json:"role_permission" orm:"with:role_id=id" description:"权限信息"`
}
type RolePermission struct {
	g.Meta       `orm:"table:role_permission"`
	RoleId       int           `json:"-"       orm:"role_id"       description:"角色id"` // 角色id
	PermissionId int           `json:"-" orm:"permission_id" description:"权限id"`       // 权限id
	Permission   []*Permission `json:"permission" orm:"with:id=permission_id" description:"权限信息"`
}

type Permission struct {
	g.Meta   `orm:"table:permission"`
	Id       int    `json:"id"         orm:"id"         description:""`        //
	Function string `json:"function"   orm:"function"   description:"功能名称"`    // 功能名称
	Redirect string `json:"redirect"   orm:"redirect"   description:"路由重定向地址"` // 路由重定向地址
}
