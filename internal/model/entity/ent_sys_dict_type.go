// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysDictType is the golang structure for table ent_sys_dict_type.
type EntSysDictType struct {
	DictId    int64       `json:"dict_id"    orm:"dict_id"    description:"字典主键"`        // 字典主键
	DictName  string      `json:"dict_name"  orm:"dict_name"  description:"字典名称"`        // 字典名称
	DictType  string      `json:"dict_type"  orm:"dict_type"  description:"字典类型"`        // 字典类型
	Status    int         `json:"status"     orm:"status"     description:"状态（0正常 1停用）"` // 状态（0正常 1停用）
	CreateBy  int         `json:"create_by"  orm:"create_by"  description:"创建者"`         // 创建者
	UpdateBy  int         `json:"update_by"  orm:"update_by"  description:"更新者"`         // 更新者
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`          // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建日期"`        // 创建日期
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"修改日期"`        // 修改日期
}
