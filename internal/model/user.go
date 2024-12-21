package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type User struct {
	g.Meta   `orm:"table:user"`
	Id       int         `json:"id"         orm:"id"         description:""`               //
	Name     string      `json:"name"       orm:"name"       description:"用户名"`            // 用户名
	Status   int         `json:"status"     orm:"status"     description:"账号状态，1-停用 2-启用"` // 账号状态，1-停用 2-启用
	RoleId   int         `json:"role_id"    orm:"role_id"    description:"角色权限id"`         // 角色权限id
	Note     string      `json:"note"       orm:"note"       description:"备注"`             // 备注
	UpdateAt *gtime.Time `json:"update_at"  orm:"update_at"  description:"更新时间"`           // 更新时间
	RoleName `json:"role_name" orm:"with:id=role_id" description:"角色名称"`
}

type RoleName struct {
	g.Meta `orm:"table:role"`
	Id     int    `json:"-"     orm:"id"     description:""`
	Name   string `json:"name"   orm:"name"   description:"角色名称"`
}
