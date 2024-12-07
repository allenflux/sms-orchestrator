// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Env is the golang structure for table env.
type Env struct {
	Id        int         `json:"id"         orm:"id"         description:""`       //
	Name      string      `json:"name"       orm:"name"       description:"环境名称"`   // 环境名称
	DeviceId  int         `json:"device_id"  orm:"device_id"  description:"设备id"`   // 设备id
	ClientId  int         `json:"client_id"  orm:"client_id"  description:"客户ID"`   // 客户ID
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`     // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"环境保存时间"` // 环境保存时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`   // 删除时间
}
