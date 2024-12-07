// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Pic is the golang structure for table pic.
type Pic struct {
	Id        int         `json:"id"          orm:"id"          description:""`           //
	Name      string      `json:"name"        orm:"name"        description:"名称"`         // 名称
	Url       string      `json:"url"         orm:"url"         description:"图片地址"`       // 图片地址
	PicTypeId int         `json:"pic_type_id" orm:"pic_type_id" description:"类型ID"`       // 类型ID
	Remark    string      `json:"remark"      orm:"remark"      description:"备注"`         // 备注
	CreatedAt *gtime.Time `json:"created_at"  orm:"created_at"  description:""`           //
	UpdatedAt *gtime.Time `json:"updated_at"  orm:"updated_at"  description:""`           //
	Type      int         `json:"type"        orm:"type"        description:"1-头像 2-封面图"` // 1-头像 2-封面图
	ClientId  int         `json:"client_id"   orm:"client_id"   description:"用户ID"`       // 用户ID
}
