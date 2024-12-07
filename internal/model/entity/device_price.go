// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DevicePrice is the golang structure for table device_price.
type DevicePrice struct {
	Id          int         `json:"id"           orm:"id"           description:""`           //
	Name        string      `json:"name"         orm:"name"         description:"价格表名称"`      // 价格表名称
	DeviceModel string      `json:"device_model" orm:"device_model" description:"设备型号"`       // 设备型号
	LeaseCount  int         `json:"lease_count"  orm:"lease_count"  description:"租赁点数/天"`     // 租赁点数/天
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`         // 备注
	Enabled     int         `json:"enabled"      orm:"enabled"      description:"是否启用1启用2停用"` // 是否启用1启用2停用
	UpdatedBy   int         `json:"updated_by"   orm:"updated_by"   description:"操作人"`        // 操作人
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:"创建时间"`       // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:"修改时间"`       // 修改时间
	DeletedAt   *gtime.Time `json:"deleted_at"   orm:"deleted_at"   description:"删除时间"`       // 删除时间
}
