// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EntSysPost is the golang structure for table ent_sys_post.
type EntSysPost struct {
	PostId    int64       `json:"post_id"    orm:"post_id"    description:"岗位ID"`        // 岗位ID
	PostCode  string      `json:"post_code"  orm:"post_code"  description:"岗位编码"`        // 岗位编码
	PostName  string      `json:"post_name"  orm:"post_name"  description:"岗位名称"`        // 岗位名称
	PostSort  int         `json:"post_sort"  orm:"post_sort"  description:"显示顺序"`        // 显示顺序
	Status    int         `json:"status"     orm:"status"     description:"状态（0正常 1停用）"` // 状态（0正常 1停用）
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`          // 备注
	CreatedBy int64       `json:"created_by" orm:"created_by" description:"创建人"`         // 创建人
	UpdatedBy int64       `json:"updated_by" orm:"updated_by" description:"修改人"`         // 修改人
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`        // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"修改时间"`        // 修改时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`        // 删除时间
}
