// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUserDetail is the golang structure for table app_user_detail.
type AppUserDetail struct {
	Id                uint        `json:"id"                  orm:"id"                  description:""`                  //
	AppUserId         int         `json:"app_user_id"         orm:"app_user_id"         description:"关联app_user表的id"`    // 关联app_user表的id
	FirstnameCn       string      `json:"firstname_cn"        orm:"firstname_cn"        description:"中文姓"`               // 中文姓
	LastnameCn        string      `json:"lastname_cn"         orm:"lastname_cn"         description:"中文名"`               // 中文名
	FirstnamePinyin   string      `json:"firstname_pinyin"    orm:"firstname_pinyin"    description:"拼音姓"`               // 拼音姓
	LastnamePinyin    string      `json:"lastname_pinyin"     orm:"lastname_pinyin"     description:"拼音名"`               // 拼音名
	Birthday          *gtime.Time `json:"birthday"            orm:"birthday"            description:"出生日期"`              // 出生日期
	AnnualIncome      float64     `json:"annual_income"       orm:"annual_income"       description:"年收入"`               // 年收入
	Career            string      `json:"career"              orm:"career"              description:"职业"`                // 职业
	Position          string      `json:"position"            orm:"position"            description:"职位"`                // 职位
	Region            string      `json:"region"              orm:"region"              description:"地区"`                // 地区
	Addr              string      `json:"addr"                orm:"addr"                description:"地址"`                // 地址
	IdCardNo          string      `json:"id_card_no"          orm:"id_card_no"          description:"身份证号"`              // 身份证号
	IdCardExpiredAt   *gtime.Time `json:"id_card_expired_at"  orm:"id_card_expired_at"  description:"证件有效期"`             // 证件有效期
	LongTermEffective int         `json:"long_term_effective" orm:"long_term_effective" description:"证件是否长期有效1代表是2代表不是"` // 证件是否长期有效1代表是2代表不是
	IdCardFrontPic    string      `json:"id_card_front_pic"   orm:"id_card_front_pic"   description:"身份证正面照片"`           // 身份证正面照片
	IdCardBackPic     string      `json:"id_card_back_pic"    orm:"id_card_back_pic"    description:"身份证背面照片"`           // 身份证背面照片
	HasCommit         int         `json:"has_commit"          orm:"has_commit"          description:"是否提交到万事达,1否2是"`     // 是否提交到万事达,1否2是
	CreatedAt         *gtime.Time `json:"created_at"          orm:"created_at"          description:"创建时间"`              // 创建时间
	UpdatedAt         *gtime.Time `json:"updated_at"          orm:"updated_at"          description:"修改时间"`              // 修改时间
	DeletedAt         *gtime.Time `json:"deleted_at"          orm:"deleted_at"          description:"删除时间"`              // 删除时间
}
