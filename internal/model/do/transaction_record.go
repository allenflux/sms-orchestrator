// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TransactionRecord is the golang structure of table transaction_record for DAO operations like Where/Data.
type TransactionRecord struct {
	g.Meta           `orm:"table:transaction_record, do:true"`
	Id               interface{} //
	Xid              interface{} // 用户xid
	RefId            interface{} //
	CardAccountId    interface{} //
	CardLast4        interface{} //
	CardEmbossedName interface{} //
	CreatedAt        *gtime.Time // 创建时间
	Status           interface{} // 状态 pending 待处理  posted 已完成 declined 已拒绝  vodi 已取消
	PostedAt         *gtime.Time //
	Intent           interface{} // 交易目的/方式，charge收费 refund退款 topup充值 repay偿还 cashback现金返回 interest利息 fee费用 other其他
	Amount           interface{} //
	EntryType        interface{} //
	Currency         interface{} //
	Description      interface{} // 响应消息（例如：响应被接受）
}
