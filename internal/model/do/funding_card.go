// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FundingCard is the golang structure of table funding_card for DAO operations like Where/Data.
type FundingCard struct {
	g.Meta        `orm:"table:funding_card, do:true"`
	Id            interface{} //
	FundingCardNo interface{} // 资金卡号
	BankUid       interface{} // 银行uid
	Balance       interface{} // 余额（HKD）
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 修改时间
	DeletedAt     *gtime.Time // 删除时间
}
