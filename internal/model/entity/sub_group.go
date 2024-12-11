// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SubGroup is the golang structure for table sub_group.
type SubGroup struct {
	Id           int         `json:"id"             orm:"id"             description:""`      //
	SubUserId    int         `json:"sub_user_id"    orm:"sub_user_id"    description:"子账号id"` // 子账号id
	ProjectId    int         `json:"project_id"     orm:"project_id"     description:"项目id"`  // 项目id
	SubGroupName string      `json:"sub_group_name" orm:"sub_group_name" description:"分组名称"`  // 分组名称
	CreatedAt    *gtime.Time `json:"created_at"     orm:"created_at"     description:"创建时间"`  // 创建时间
	UpdateAt     *gtime.Time `json:"update_at"      orm:"update_at"      description:"修改时间"`  // 修改时间
	DeleteAt     *gtime.Time `json:"delete_at"      orm:"delete_at"      description:"删除时间"`  // 删除时间
}
