// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientLabel is the golang structure for table client_label.
type ClientLabel struct {
	Id        int         `json:"id"         orm:"id"         description:""`     //
	Name      string      `json:"name"       orm:"name"       description:"名称"`   // 名称
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"` // 删除时间
}
