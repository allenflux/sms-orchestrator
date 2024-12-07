// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id        int         `json:"id"        orm:"id"         ` //
	Status    int         `json:"status"    orm:"status"     ` // 状态
	ListOrder int         `json:"listOrder" orm:"list_order" ` // 排序
	Name      string      `json:"name"      orm:"name"       ` // 角色名称
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	DataScope int         `json:"dataScope" orm:"data_scope" ` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	RoleType  int         `json:"roleType"  orm:"role_type"  ` // 1-群控 2-企业
}
