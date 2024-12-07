// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApplyList is the golang structure of table apply_list for DAO operations like Where/Data.
type ApplyList struct {
	g.Meta        `orm:"table:apply_list, do:true"`
	Id            interface{} //
	ApplyId       interface{} // 申请单ID
	CardNo        interface{} // 卡号
	AppUserId     interface{} // 平台用户ID
	Status        interface{} // 1-申请中 2-审核通过 3-已发卡 4-已激活 5-已实名 6-已驳回 7-已取消 8-供应商异常
	TotalQuota    interface{} // 总额度
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	UserCardImage interface{} // 后台上传的卡片照片
	UserName      interface{} //
	PhoneNumber   interface{} // User Phone Number
	BankUid       interface{} // bank_uid
	Balance       interface{} // 余额信息
	Note          interface{} // 备注信息
}
