// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceGrantImportFails is the golang structure for table device_grant_import_fails.
type DeviceGrantImportFails struct {
	Id            int         `json:"id"              orm:"id"              description:""`       //
	DataImportsId int         `json:"data_imports_id" orm:"data_imports_id" description:"数据导入ID"` // 数据导入ID
	DeviceDevId   string      `json:"device_dev_id"   orm:"device_dev_id"   description:"设备id"`   // 设备id
	GrantDb       string      `json:"grant_db"        orm:"grant_db"        description:"授权数据库"`  // 授权数据库
	DeviceLabels  string      `json:"device_labels"   orm:"device_labels"   description:"设备标签"`   // 设备标签
	Remark        string      `json:"remark"          orm:"remark"          description:"备注"`     // 备注
	FailReason    string      `json:"fail_reason"     orm:"fail_reason"     description:"失败原因"`   // 失败原因
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`   // 创建时间
	ClientId      int         `json:"client_id"       orm:"client_id"       description:"客户ID"`   // 客户ID
}
