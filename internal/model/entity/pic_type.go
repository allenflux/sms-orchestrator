// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PicType is the golang structure for table pic_type.
type PicType struct {
	Id        uint        `json:"id"         orm:"id"         description:""`                //
	Name      string      `json:"name"       orm:"name"       description:"图片类型"`            // 图片类型
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`            // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"修改时间"`            // 修改时间
	Type      int         `json:"type"       orm:"type"       description:"1-头像 2-封面图 3-昵称"` // 1-头像 2-封面图 3-昵称
	ClientId  int         `json:"client_id"  orm:"client_id"  description:"用户ID"`            // 用户ID
}
