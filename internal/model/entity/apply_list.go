// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ApplyList is the golang structure for table apply_list.
type ApplyList struct {
	Id            int         `json:"id"              orm:"id"              description:""`                                                   //
	ApplyId       string      `json:"apply_id"        orm:"apply_id"        description:"申请单ID"`                                              // 申请单ID
	CardNo        string      `json:"card_no"         orm:"card_no"         description:"卡号"`                                                 // 卡号
	AppUserId     int         `json:"app_user_id"     orm:"app_user_id"     description:"平台用户ID"`                                             // 平台用户ID
	Status        int         `json:"status"          orm:"status"          description:"1-申请中 2-审核通过 3-已发卡 4-已激活 5-已实名 6-已驳回 7-已取消 8-供应商异常"` // 1-申请中 2-审核通过 3-已发卡 4-已激活 5-已实名 6-已驳回 7-已取消 8-供应商异常
	TotalQuota    float64     `json:"total_quota"     orm:"total_quota"     description:"总额度"`                                                // 总额度
	CreatedAt     *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`                                               // 创建时间
	UpdatedAt     *gtime.Time `json:"updated_at"      orm:"updated_at"      description:"更新时间"`                                               // 更新时间
	UserCardImage string      `json:"user_card_image" orm:"user_card_image" description:"后台上传的卡片照片"`                                          // 后台上传的卡片照片
	UserName      string      `json:"user_name"       orm:"user_name"       description:""`                                                   //
	PhoneNumber   string      `json:"phone_number"    orm:"phone_number"    description:"User Phone Number"`                                  // User Phone Number
	BankUid       string      `json:"bank_uid"        orm:"bank_uid"        description:"bank_uid"`                                           // bank_uid
	Balance       float64     `json:"balance"         orm:"balance"         description:"余额信息"`                                               // 余额信息
	Note          string      `json:"note"            orm:"note"            description:"备注信息"`                                               // 备注信息
}
