// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure for table words.
type Words struct {
	Id        uint        `json:"id"         orm:"id"         description:""`     //
	Content   string      `json:"content"    orm:"content"    description:"话术内容"` // 话术内容
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"` // 更新时间
	ClientId  int         `json:"client_id"  orm:"client_id"  description:"客户ID"` // 客户ID
}
