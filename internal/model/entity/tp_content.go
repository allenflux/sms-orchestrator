// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TpContent is the golang structure for table tp_content.
type TpContent struct {
	Id        uint        `json:"id"         orm:"id"         description:""`     //
	Str       string      `json:"str"        orm:"str"        description:""`     //
	Remark    string      `json:"remark"     orm:"remark"     description:""`     //
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
}
