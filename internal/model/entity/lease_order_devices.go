// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LeaseOrderDevices is the golang structure for table lease_order_devices.
type LeaseOrderDevices struct {
	Id                     int         `json:"id"                        orm:"id"                        description:""`                  //
	LeaseOrderId           int         `json:"lease_order_id"            orm:"lease_order_id"            description:"（租赁/续费）订单ID"`       // （租赁/续费）订单ID
	DeviceId               int         `json:"device_id"                 orm:"device_id"                 description:"设备主键ID"`            // 设备主键ID
	ClientId               int         `json:"client_id"                 orm:"client_id"                 description:"客户ID"`              // 客户ID
	OrderNo                string      `json:"order_no"                  orm:"order_no"                  description:"订单号"`               // 订单号
	CreatedAt              *gtime.Time `json:"created_at"                orm:"created_at"                description:"创建时间"`              // 创建时间
	UpdatedAt              *gtime.Time `json:"updated_at"                orm:"updated_at"                description:"更新时间"`              // 更新时间
	LeaseAt                *gtime.Time `json:"lease_at"                  orm:"lease_at"                  description:"租赁时间"`              // 租赁时间
	LeaseEndAt             *gtime.Time `json:"lease_end_at"              orm:"lease_end_at"              description:"到期时间"`              // 到期时间
	UnbindAt               *gtime.Time `json:"unbind_at"                 orm:"unbind_at"                 description:"解绑时间"`              // 解绑时间
	Remark                 string      `json:"remark"                    orm:"remark"                    description:"备注"`                // 备注
	LeaseDays              int         `json:"lease_days"                orm:"lease_days"                description:"租赁天数"`              // 租赁天数
	BindSpecialOperateId   int         `json:"bind_special_operate_id"   orm:"bind_special_operate_id"   description:"特殊操作表ID（绑定）"`       // 特殊操作表ID（绑定）
	UnbindSpecialOperateId int         `json:"unbind_special_operate_id" orm:"unbind_special_operate_id" description:"特殊操作表ID（解绑）"`       // 特殊操作表ID（解绑）
	DelaySpecialOperateId  int         `json:"delay_special_operate_id"  orm:"delay_special_operate_id"  description:"特殊操作表ID（延期）"`       // 特殊操作表ID（延期）
	ProduceWay             int         `json:"produce_way"               orm:"produce_way"               description:"1-租赁  2-续费 3-批量绑定"` // 1-租赁  2-续费 3-批量绑定
}
