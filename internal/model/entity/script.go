// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Script is the golang structure for table script.
type Script struct {
	Id          int         `json:"id"           orm:"id"           description:""`                //
	Name        string      `json:"name"         orm:"name"         description:"脚本名称"`            // 脚本名称
	ClientId    int         `json:"client_id"    orm:"client_id"    description:"客户id"`            // 客户id
	DownloadUrl string      `json:"download_url" orm:"download_url" description:"apk名称"`           // apk名称
	AppName     string      `json:"app_name"     orm:"app_name"     description:""`                //
	ObjectName  string      `json:"object_name"  orm:"object_name"  description:"oss的object_name"` // oss的object_name
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:"创建时间"`            // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:"更新时间"`            // 更新时间
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`              // 备注
	DeletedAt   *gtime.Time `json:"deleted_at"   orm:"deleted_at"   description:""`                //
}
