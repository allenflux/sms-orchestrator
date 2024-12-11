// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProject is the golang structure for table user_project.
type UserProject struct {
	Id        int         `json:"id"         orm:"id"         description:""`     //
	UserId    int         `json:"user_id"    orm:"user_id"    description:"用户id"` // 用户id
	ProjectId int         `json:"project_id" orm:"project_id" description:"项目id"` // 项目id
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	UpdateAt  *gtime.Time `json:"update_at"  orm:"update_at"  description:"修改时间"` // 修改时间
	DeleteAt  *gtime.Time `json:"delete_at"  orm:"delete_at"  description:"删除时间"` // 删除时间
}
