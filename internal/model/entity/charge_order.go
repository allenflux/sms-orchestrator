// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChargeOrder is the golang structure for table charge_order.
type ChargeOrder struct {
	Id                 int         `json:"id"                   orm:"id"                   description:""`           //
	ClientId           int         `json:"client_id"            orm:"client_id"            description:"客户id"`       // 客户id
	ChargeType         int         `json:"charge_type"          orm:"charge_type"          description:"加款类型1余额2授信"` // 加款类型1余额2授信
	ChargeAmount       float64     `json:"charge_amount"        orm:"charge_amount"        description:"添加金额"`       // 添加金额
	BalanceAfterCharge float64     `json:"balance_after_charge" orm:"balance_after_charge" description:"加款后余额"`      // 加款后余额
	CreditAfterCharge  float64     `json:"credit_after_charge"  orm:"credit_after_charge"  description:"加款后授信"`      // 加款后授信
	ChargeBank         string      `json:"charge_bank"          orm:"charge_bank"          description:"加款银行"`       // 加款银行
	Remark             string      `json:"remark"               orm:"remark"               description:"备注"`         // 备注
	CreatedBy          int         `json:"created_by"           orm:"created_by"           description:"操作人"`        // 操作人
	CreatedAt          *gtime.Time `json:"created_at"           orm:"created_at"           description:"加款时间"`       // 加款时间
	DeletedAt          *gtime.Time `json:"deleted_at"           orm:"deleted_at"           description:"删除时间"`       // 删除时间
	OrderNo            string      `json:"order_no"             orm:"order_no"             description:"订单号"`        // 订单号
}
