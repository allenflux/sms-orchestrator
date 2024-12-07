// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CardLevel is the golang structure for table card_level.
type CardLevel struct {
	Id        int         `json:"id"         orm:"id"         description:""`          //
	Name      string      `json:"name"       orm:"name"       description:"卡片等级名称"`    // 卡片等级名称
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`      // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"`      // 更新时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`      // 删除时间
	File      string      `json:"file"       orm:"file"       description:"文件base64值"` // 文件base64值
}
