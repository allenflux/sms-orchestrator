// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RechargeList is the golang structure for table recharge_list.
type RechargeList struct {
	Id            int         `json:"id"             orm:"id"             description:""`        //
	FundCardNo    string      `json:"fund_card_no"   orm:"fund_card_no"   description:"资金卡号"`    // 资金卡号
	CardNo        string      `json:"card_no"        orm:"card_no"        description:"对方卡号"`    // 对方卡号
	Amount        float64     `json:"amount"         orm:"amount"         description:"金额（hkd）"` // 金额（hkd）
	Name          string      `json:"name"           orm:"name"           description:"对方姓名"`    // 对方姓名
	ApplyListId   int         `json:"apply_list_id"  orm:"apply_list_id"  description:"申请列表id"`  // 申请列表id
	AdminUsername string      `json:"admin_username" orm:"admin_username" description:"操作账号"`    // 操作账号
	CreatedAt     *gtime.Time `json:"created_at"     orm:"created_at"     description:"充值时间"`    // 充值时间
	UpdatedAt     *gtime.Time `json:"updated_at"     orm:"updated_at"     description:""`        //
	DeletedAt     *gtime.Time `json:"deleted_at"     orm:"deleted_at"     description:""`        //
}
