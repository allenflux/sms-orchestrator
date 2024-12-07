// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientFundFlow is the golang structure for table client_fund_flow.
type ClientFundFlow struct {
	Id              int         `json:"id"                orm:"id"                description:""`                                                           //
	FundType        int         `json:"fund_type"         orm:"fund_type"         description:"1-租赁 2-续费 3-退款 4-余额 5-授信  6-产生订单  7-取消订单 8-驳回订单 9-绑定 10-延期"` // 1-租赁 2-续费 3-退款 4-余额 5-授信  6-产生订单  7-取消订单 8-驳回订单 9-绑定 10-延期
	OrderId         int         `json:"order_id"          orm:"order_id"          description:"订单ID"`                                                       // 订单ID
	ClientId        int         `json:"client_id"         orm:"client_id"         description:"客户ID"`                                                       // 客户ID
	BalanceBefore   float64     `json:"balance_before"    orm:"balance_before"    description:"操作前余额"`                                                      // 操作前余额
	BalanceAfter    float64     `json:"balance_after"     orm:"balance_after"     description:"操作后余额"`                                                      // 操作后余额
	CreditBefore    float64     `json:"credit_before"     orm:"credit_before"     description:"操作前授信金额"`                                                    // 操作前授信金额
	CreditAfter     float64     `json:"credit_after"      orm:"credit_after"      description:"操作后授信金额"`                                                    // 操作后授信金额
	AdvanceBefore   float64     `json:"advance_before"    orm:"advance_before"    description:"操作前预扣金额"`                                                    // 操作前预扣金额
	AdvanceAfter    float64     `json:"advance_after"     orm:"advance_after"     description:"操作后预扣金额"`                                                    // 操作后预扣金额
	UseCreditBefore float64     `json:"use_credit_before" orm:"use_credit_before" description:"操作前已使用预扣金额"`                                                 // 操作前已使用预扣金额
	UseCreditAfter  float64     `json:"use_credit_after"  orm:"use_credit_after"  description:"操作后已使用预扣金额"`                                                 // 操作后已使用预扣金额
	Amount          float64     `json:"amount"            orm:"amount"            description:"操作数量"`                                                       // 操作数量
	CreatedBy       int         `json:"created_by"        orm:"created_by"        description:"操作人ID"`                                                      // 操作人ID
	CreatedAt       *gtime.Time `json:"created_at"        orm:"created_at"        description:"创建时间"`                                                       // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"        orm:"updated_at"        description:"更新时间"`                                                       // 更新时间
	DeletedAt       *gtime.Time `json:"deleted_at"        orm:"deleted_at"        description:"删除时间"`                                                       // 删除时间
	TransactionType int         `json:"transaction_type"  orm:"transaction_type"  description:"交易类型1-存入 2-支出"`                                              // 交易类型1-存入 2-支出
	Remark          string      `json:"remark"            orm:"remark"            description:"备注"`                                                         // 备注
	OrderNo         string      `json:"order_no"          orm:"order_no"          description:"订单号"`                                                        // 订单号
}
