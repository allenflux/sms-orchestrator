// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TransactionRecord is the golang structure for table transaction_record.
type TransactionRecord struct {
	Id               string      `json:"id"                 orm:"id"                 description:""`                                                                                //
	Xid              string      `json:"xid"                orm:"xid"                description:"用户xid"`                                                                           // 用户xid
	RefId            string      `json:"ref_id"             orm:"ref_id"             description:""`                                                                                //
	CardAccountId    string      `json:"card_account_id"    orm:"card_account_id"    description:""`                                                                                //
	CardLast4        string      `json:"card_last_4"        orm:"card_last_4"        description:""`                                                                                //
	CardEmbossedName string      `json:"card_embossed_name" orm:"card_embossed_name" description:""`                                                                                //
	CreatedAt        *gtime.Time `json:"created_at"         orm:"created_at"         description:"创建时间"`                                                                            // 创建时间
	Status           string      `json:"status"             orm:"status"             description:"状态 pending 待处理  posted 已完成 declined 已拒绝  vodi 已取消"`                               // 状态 pending 待处理  posted 已完成 declined 已拒绝  vodi 已取消
	PostedAt         *gtime.Time `json:"posted_at"          orm:"posted_at"          description:""`                                                                                //
	Intent           string      `json:"intent"             orm:"intent"             description:"交易目的/方式，charge收费 refund退款 topup充值 repay偿还 cashback现金返回 interest利息 fee费用 other其他"` // 交易目的/方式，charge收费 refund退款 topup充值 repay偿还 cashback现金返回 interest利息 fee费用 other其他
	Amount           float64     `json:"amount"             orm:"amount"             description:""`                                                                                //
	EntryType        string      `json:"entry_type"         orm:"entry_type"         description:""`                                                                                //
	Currency         string      `json:"currency"           orm:"currency"           description:""`                                                                                //
	Description      string      `json:"description"        orm:"description"        description:"响应消息（例如：响应被接受）"`                                                                  // 响应消息（例如：响应被接受）
}
