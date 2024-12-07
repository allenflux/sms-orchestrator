// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysDept is the golang structure for table ent_sys_dept.
type EntSysDept struct {
	DeptId    int64       `json:"dept_id"    orm:"dept_id"    description:"部门id"`          // 部门id
	ParentId  int64       `json:"parent_id"  orm:"parent_id"  description:"父部门id"`         // 父部门id
	Ancestors string      `json:"ancestors"  orm:"ancestors"  description:"祖级列表"`          // 祖级列表
	DeptName  string      `json:"dept_name"  orm:"dept_name"  description:"部门名称"`          // 部门名称
	OrderNum  int         `json:"order_num"  orm:"order_num"  description:"显示顺序"`          // 显示顺序
	Leader    string      `json:"leader"     orm:"leader"     description:"负责人"`           // 负责人
	Phone     string      `json:"phone"      orm:"phone"      description:"联系电话"`          // 联系电话
	Email     string      `json:"email"      orm:"email"      description:"邮箱"`            // 邮箱
	Status    int         `json:"status"     orm:"status"     description:"部门状态（0正常 1停用）"` // 部门状态（0正常 1停用）
	CreatedBy int64       `json:"created_by" orm:"created_by" description:"创建人"`           // 创建人
	UpdatedBy int64       `json:"updated_by" orm:"updated_by" description:"修改人"`           // 修改人
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`          // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"修改时间"`          // 修改时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`          // 删除时间
}
