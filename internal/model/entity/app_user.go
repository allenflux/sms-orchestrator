// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUser is the golang structure for table app_user.
type AppUser struct {
	Id                  uint        `json:"id"                   orm:"id"                   description:""`                               //
	Password            string      `json:"password"             orm:"password"             description:"密码"`                             // 密码
	Salt                string      `json:"salt"                 orm:"salt"                 description:"盐"`                              // 盐
	TransactionPassword string      `json:"transaction_password" orm:"transaction_password" description:"交易密码"`                           // 交易密码
	TransactionSalt     string      `json:"transaction_salt"     orm:"transaction_salt"     description:"交易密码的盐"`                         // 交易密码的盐
	CreatedAt           *gtime.Time `json:"created_at"           orm:"created_at"           description:"创建时间"`                           // 创建时间
	UpdatedAt           *gtime.Time `json:"updated_at"           orm:"updated_at"           description:"修改时间"`                           // 修改时间
	DeletedAt           *gtime.Time `json:"deleted_at"           orm:"deleted_at"           description:"删除时间"`                           // 删除时间
	Uid                 string      `json:"uid"                  orm:"uid"                  description:"银行uid"`                          // 银行uid
	Mobile              string      `json:"mobile"               orm:"mobile"               description:"手机号"`                            // 手机号
	Status              int         `json:"status"               orm:"status"               description:"1-待KYC 2-待上传 3-待配卡 4-待申请 5-已申请"` // 1-待KYC 2-待上传 3-待配卡 4-待申请 5-已申请
	KycStatus           int         `json:"kyc_status"           orm:"kyc_status"           description:"1-待kyc 2-已kyc"`                  // 1-待kyc 2-已kyc
	UserCardStatus      int         `json:"user_card_status"     orm:"user_card_status"     description:"1-未上传 2-已上传"`                    // 1-未上传 2-已上传
	BankCardStatus      int         `json:"bank_card_status"     orm:"bank_card_status"     description:"1-未配卡 2-已配卡"`                    // 1-未配卡 2-已配卡
	BankImage           string      `json:"bank_image"           orm:"bank_image"           description:"万事达照片"`                          // 万事达照片
	CardLevelId         int         `json:"card_level_id"        orm:"card_level_id"        description:"卡片等级表ID"`                        // 卡片等级表ID
	UserCardImage       string      `json:"user_card_image"      orm:"user_card_image"      description:"用户卡照片"`                          // 用户卡照片
	Remark              string      `json:"remark"               orm:"remark"               description:"备注"`                             // 备注
	Email               string      `json:"email"                orm:"email"                description:"邮箱"`                             // 邮箱
}
