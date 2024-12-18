package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SubUser struct {
	g.Meta   `orm:"table:user"`
	Id       int         `json:"id"         orm:"id"         description:""`               //
	Name     string      `json:"name"       orm:"name"       description:"用户名"`            // 用户名
	Status   int         `json:"status"     orm:"status"     description:"账号状态，1-停用 2-启用"` // 账号状态，1-停用 2-启用
	Note     string      `json:"note"       orm:"note"       description:"备注"`             // 备注
	UpdateAt *gtime.Time `json:"update_at"  orm:"update_at"  description:"更新时间"`           // 更新时间
	Project  []Project   `json:"project" orm:"with:associated_account_id=id"`
}

type Project struct {
	g.Meta              `orm:"table:project_list"`
	ProjectName         string `json:"project_name"          orm:"project_name"          description:"项目名称"`
	AssociatedAccountId int    `json:"associated_account_id" orm:"associated_account_id" description:"关联子账号id"`
}
