// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RechargeList is the golang structure of table recharge_list for DAO operations like Where/Data.
type RechargeList struct {
	g.Meta        `orm:"table:recharge_list, do:true"`
	Id            interface{} //
	FundCardNo    interface{} // 资金卡号
	CardNo        interface{} // 对方卡号
	Amount        interface{} // 金额（hkd）
	Name          interface{} // 对方姓名
	ApplyListId   interface{} // 申请列表id
	AdminUsername interface{} // 操作账号
	CreatedAt     *gtime.Time // 充值时间
	UpdatedAt     *gtime.Time //
	DeletedAt     *gtime.Time //
}
