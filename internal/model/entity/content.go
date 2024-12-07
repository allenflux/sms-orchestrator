// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Content is the golang structure for table content.
type Content struct {
	Id        int         `json:"id"         orm:"id"         description:""`     //
	Link      string      `json:"link"       orm:"link"       description:"跳转链接"` // 跳转链接
	Title     string      `json:"title"      orm:"title"      description:"标题"`   // 标题
	Content   string      `json:"content"    orm:"content"    description:"内容"`   // 内容
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`     //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`     //
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`     //
}
