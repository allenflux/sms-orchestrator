// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysDictData is the golang structure for table ent_sys_dict_data.
type EntSysDictData struct {
	DictCode  int64       `json:"dict_code"  orm:"dict_code"  description:"字典编码"`         // 字典编码
	DictSort  int         `json:"dict_sort"  orm:"dict_sort"  description:"字典排序"`         // 字典排序
	DictLabel string      `json:"dict_label" orm:"dict_label" description:"字典标签"`         // 字典标签
	DictValue string      `json:"dict_value" orm:"dict_value" description:"字典键值"`         // 字典键值
	DictType  string      `json:"dict_type"  orm:"dict_type"  description:"字典类型"`         // 字典类型
	CssClass  string      `json:"css_class"  orm:"css_class"  description:"样式属性（其他样式扩展）"` // 样式属性（其他样式扩展）
	ListClass string      `json:"list_class" orm:"list_class" description:"表格回显样式"`       // 表格回显样式
	IsDefault int         `json:"is_default" orm:"is_default" description:"是否默认（1是 0否）"`  // 是否默认（1是 0否）
	Status    int         `json:"status"     orm:"status"     description:"状态（0正常 1停用）"`  // 状态（0正常 1停用）
	CreateBy  int64       `json:"create_by"  orm:"create_by"  description:"创建者"`          // 创建者
	UpdateBy  int64       `json:"update_by"  orm:"update_by"  description:"更新者"`          // 更新者
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`           // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"修改时间"`         // 修改时间
}
