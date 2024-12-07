// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientLinkLabel is the golang structure for table client_link_label.
type ClientLinkLabel struct {
	Id            int         `json:"id"              orm:"id"              description:""`       //
	ClientId      int         `json:"client_id"       orm:"client_id"       description:"客户id"`   // 客户id
	ClientLabelId int         `json:"client_label_id" orm:"client_label_id" description:"客户标签id"` // 客户标签id
	DeletedAt     *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:"删除时间"`   // 删除时间
}
