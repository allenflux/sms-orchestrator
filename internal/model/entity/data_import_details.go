// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DataImportDetails is the golang structure for table data_import_details.
type DataImportDetails struct {
	Id           int         `json:"id"             orm:"id"             description:""`          //
	DataImportId int         `json:"data_import_id" orm:"data_import_id" description:"导入ID"`      // 导入ID
	DevId        string      `json:"dev_id"         orm:"dev_id"         description:"设备ID"`      // 设备ID
	Remark1      string      `json:"remark_1"       orm:"remark1"        description:"备注1"`       // 备注1
	Remark2      string      `json:"remark_2"       orm:"remark2"        description:"备注2"`       // 备注2
	CreatedAt    *gtime.Time `json:"created_at"     orm:"created_at"     description:""`          //
	Status       int         `json:"status"         orm:"status"         description:"1-成功 2-失败"` // 1-成功 2-失败
}
