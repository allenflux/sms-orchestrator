// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RefundOrder is the golang structure for table refund_order.
type RefundOrder struct {
	Id           int         `json:"id"            orm:"id"            description:""`      //
	RefundAmount float64     `json:"refund_amount" orm:"refund_amount" description:"退款金额"`  // 退款金额
	ClientId     int         `json:"client_id"     orm:"client_id"     description:"客户ID"`  // 客户ID
	OrderNo      string      `json:"order_no"      orm:"order_no"      description:"订单号"`   // 订单号
	CreatedBy    int         `json:"created_by"    orm:"created_by"    description:"操作人ID"` // 操作人ID
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`  // 创建时间
}
