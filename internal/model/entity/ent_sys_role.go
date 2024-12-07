// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysRole is the golang structure for table ent_sys_role.
type EntSysRole struct {
	Id        int         `json:"id"         orm:"id"         description:""`                                               //
	Status    int         `json:"status"     orm:"status"     description:"状态;0:禁用;1:正常"`                                   // 状态;0:禁用;1:正常
	ListOrder int         `json:"list_order" orm:"list_order" description:"排序"`                                             // 排序
	Name      string      `json:"name"       orm:"name"       description:"角色名称"`                                           // 角色名称
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`                                             // 备注
	DataScope int         `json:"data_scope" orm:"data_scope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`                                           // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"`                                           // 更新时间
}
