// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure for table permission.
type Permission struct {
	Id        int         `json:"id"         orm:"id"         description:""`     //
	Name      string      `json:"name"       orm:"name"       description:"功能名称"` // 功能名称
	CreatedAd *gtime.Time `json:"created_ad" orm:"created_ad" description:"创建时间"` // 创建时间
	UpdateAt  *gtime.Time `json:"update_at"  orm:"update_at"  description:"修改时间"` // 修改时间
	DeleteAt  *gtime.Time `json:"delete_at"  orm:"delete_at"  description:"删除时间"` // 删除时间
}
