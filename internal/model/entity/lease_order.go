// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LeaseOrder is the golang structure for table lease_order.
type LeaseOrder struct {
	Id               int         `json:"id"                 orm:"id"                 description:""`                     //
	ClientId         int         `json:"client_id"          orm:"client_id"          description:"客户id"`                 // 客户id
	LeaseDeviceModel string      `json:"lease_device_model" orm:"lease_device_model" description:"租赁设备型号"`               // 租赁设备型号
	LeaseDeviceNum   int         `json:"lease_device_num"   orm:"lease_device_num"   description:"租赁设备数量"`               // 租赁设备数量
	LeaseDeviceDays  int         `json:"lease_device_days"  orm:"lease_device_days"  description:"租赁设备天数"`               // 租赁设备天数
	LeaseType        int         `json:"lease_type"         orm:"lease_type"         description:"租赁类型1租赁2续费"`           // 租赁类型1租赁2续费
	State            int         `json:"state"              orm:"state"              description:"订单状态1未处理2已完成3已取消4已驳回"` // 订单状态1未处理2已完成3已取消4已驳回
	Remark           string      `json:"remark"             orm:"remark"             description:"备注"`                   // 备注
	UpdatedBy        int         `json:"updated_by"         orm:"updated_by"         description:"操作人"`                  // 操作人
	CompletedAt      *gtime.Time `json:"completed_at"       orm:"completed_at"       description:"完成时间"`                 // 完成时间
	CreatedAt        *gtime.Time `json:"created_at"         orm:"created_at"         description:"下单时间"`                 // 下单时间
	UpdatedAt        *gtime.Time `json:"updated_at"         orm:"updated_at"         description:"取消/驳回时间"`              // 取消/驳回时间
	DeletedAt        *gtime.Time `json:"deleted_at"         orm:"deleted_at"         description:"删除时间"`                 // 删除时间
	DeviceName       string      `json:"device_name"        orm:"device_name"        description:"设备名称"`                 // 设备名称
	LeaseCount       int         `json:"lease_count"        orm:"lease_count"        description:"租赁点数"`                 // 租赁点数
	TotalLeaseCount  int         `json:"total_lease_count"  orm:"total_lease_count"  description:"订单总点数"`                // 订单总点数
	DevicePriceId    int         `json:"device_price_id"    orm:"device_price_id"    description:"价格表ID"`                // 价格表ID
	LeaseEndAt       *gtime.Time `json:"lease_end_at"       orm:"lease_end_at"       description:"租赁到期时间"`               // 租赁到期时间
	OrderNo          string      `json:"order_no"           orm:"order_no"           description:"订单号"`                  // 订单号
	Pid              int         `json:"pid"                orm:"pid"                description:"续费时的id"`               // 续费时的id
	CreatedBy        int         `json:"created_by"         orm:"created_by"         description:"操作人"`                  // 操作人
	DeviceId         int         `json:"device_id"          orm:"device_id"          description:"续费设备ID"`               // 续费设备ID
}
