// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAuthRule20240821 is the golang structure for table sys_auth_rule_20240821.
type SysAuthRule20240821 struct {
	Id         int         `json:"id"          orm:"id"          description:""`               //
	Pid        int         `json:"pid"         orm:"pid"         description:"父ID"`            // 父ID
	Name       string      `json:"name"        orm:"name"        description:"规则名称"`           // 规则名称
	Title      string      `json:"title"       orm:"title"       description:"规则名称"`           // 规则名称
	Icon       string      `json:"icon"        orm:"icon"        description:"图标"`             // 图标
	Condition  string      `json:"condition"   orm:"condition"   description:"条件"`             // 条件
	Remark     string      `json:"remark"      orm:"remark"      description:"备注"`             // 备注
	MenuType   int         `json:"menu_type"   orm:"menu_type"   description:"类型 0目录 1菜单 2按钮"` // 类型 0目录 1菜单 2按钮
	Weigh      int         `json:"weigh"       orm:"weigh"       description:"权重"`             // 权重
	IsHide     int         `json:"is_hide"     orm:"is_hide"     description:"显示状态"`           // 显示状态
	Path       string      `json:"path"        orm:"path"        description:"路由地址"`           // 路由地址
	Component  string      `json:"component"   orm:"component"   description:"组件路径"`           // 组件路径
	IsLink     int         `json:"is_link"     orm:"is_link"     description:"是否外链 1是 0否"`     // 是否外链 1是 0否
	ModuleType string      `json:"module_type" orm:"module_type" description:"所属模块"`           // 所属模块
	ModelId    int         `json:"model_id"    orm:"model_id"    description:"模型ID"`           // 模型ID
	IsIframe   int         `json:"is_iframe"   orm:"is_iframe"   description:"是否内嵌iframe"`     // 是否内嵌iframe
	IsCached   int         `json:"is_cached"   orm:"is_cached"   description:"是否缓存"`           // 是否缓存
	Redirect   string      `json:"redirect"    orm:"redirect"    description:"路由重定向地址"`        // 路由重定向地址
	IsAffix    int         `json:"is_affix"    orm:"is_affix"    description:"是否固定"`           // 是否固定
	LinkUrl    string      `json:"link_url"    orm:"link_url"    description:"链接地址"`           // 链接地址
	CreatedAt  *gtime.Time `json:"created_at"  orm:"created_at"  description:"创建日期"`           // 创建日期
	UpdatedAt  *gtime.Time `json:"updated_at"  orm:"updated_at"  description:"修改日期"`           // 修改日期
}
