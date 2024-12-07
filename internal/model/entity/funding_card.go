// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FundingCard is the golang structure for table funding_card.
type FundingCard struct {
	Id            int         `json:"id"              orm:"id"              description:""`        //
	FundingCardNo string      `json:"funding_card_no" orm:"funding_card_no" description:"资金卡号"`    // 资金卡号
	BankUid       string      `json:"bank_uid"        orm:"bank_uid"        description:"银行uid"`   // 银行uid
	Balance       float64     `json:"balance"         orm:"balance"         description:"余额（HKD）"` // 余额（HKD）
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`    // 创建时间
	UpdatedAt     *gtime.Time `json:"updated_at"      orm:"updated_at"      description:"修改时间"`    // 修改时间
	DeletedAt     *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:"删除时间"`    // 删除时间
}
