package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type UserT struct {
	Id     int    `json:"id" dc:"用户id"`
	Uid    string `json:"uid" dc:"银行uid"`
	Mobile string `json:"mobile" dc:"手机号"`
	Email  string `json:"email" dc:"邮箱"`
	//Status            int         `json:"status" dc:"状态,1-待KYC 2-待上传 3-待配卡 4-待申请 5-已申请"`
	KycStatus         int         `json:"kyc_status" dc:"1-待kyc 2-已kyc"`     // 1-待kyc 2-已kyc
	UserCardStatus    int         `json:"user_card_status" dc:"1-未上传 2-已上传"` // 1-未上传 2-已上传
	BankCardStatus    int         `json:"bank_card_status" dc:"1-未配卡 2-已配卡"` // 1-未配卡 2-已配卡
	UserCardImage     string      `json:"user_card_image" dc:"用户卡照片"`
	BankImage         string      `json:"bank_image" dc:"万事达卡照片"`
	CardLevelId       int         `json:"card_level_id" dc:"万事达卡等级id"`
	Remark            string      `json:"remark" dc:"备注"`
	CreatedAt         *gtime.Time `json:"created_at"  dc:"用户注册时间"`
	FirstnameCn       string      `json:"firstname_cn" dc:"中文姓"`
	LastnameCn        string      `json:"lastname_cn"  dc:"中文名"`
	FirstnamePinyin   string      `json:"firstname_pinyin"  dc:"英文姓"`
	LastnamePinyin    string      `json:"lastname_pinyin"  dc:"英文名"`
	Birthday          string      `json:"birthday"  dc:"生日"`
	AnnualIncome      float64     `json:"annual_income"  dc:"年收入"`
	Career            string      `json:"career"  dc:"职业"`
	Position          string      `json:"position"  dc:"职位"`
	Region            string      `json:"region"  dc:"地区"`
	Addr              string      `json:"addr"  dc:"地址"`
	IdCardNo          string      `json:"id_card_no"  dc:"身份证号"`
	IdCardExpiredAt   string      `json:"id_card_expired_at"  dc:"身份证证有效期至,长期则为null"`
	LongTermEffective int         `json:"long_term_effective"  dc:"是否长期有效,1-是,2-否"`
	IdCardFrontPic    string      `json:"id_card_front_pic" dc:"身份证正面照片(头像面)"`
	IdCardBackPic     string      `json:"id_card_back_pic" dc:"身份证背面照片(国徽面)"`
	CardName          string      `json:"card_name" dc:"用户卡等级名称"`
	ApplyStatus       int         `json:"apply_status" dc:"0-没有记录 1-申请中 2-审核通过 3-已发卡 4-已激活 5-已实名 6-已驳回 7-已取消"`
}

type UserCardLevelInfo struct {
	gmeta.Meta `orm:"table:card_level"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
}

// --------------------------我-是-分-割-线-------------------------- //

type UpdateUserInfo struct {
	// gmeta.Meta        `orm:"table:app_user_detail"`
	AppUserId         int         `json:"app_user_id" dc:"用户id"`
	FirstnameCn       string      `json:"firstname_cn" dc:"中文姓"`
	LastnameCn        string      `json:"lastname_cn"  dc:"中文名"`
	FirstnamePinyin   string      `json:"firstname_pinyin"  dc:"英文姓"`
	LastnamePinyin    string      `json:"lastname_pinyin"  dc:"英文名"`
	Birthday          *gtime.Time `json:"birthday"  dc:"生日"`
	AnnualIncome      float64     `json:"annual_income"  dc:"年收入"`
	Career            string      `json:"career"  dc:"职业"`
	Position          string      `json:"position"  dc:"职位"`
	Region            *string     `json:"region"  dc:"地区"`
	Addr              string      `json:"addr"  dc:"地址"`
	IdCardNo          string      `json:"id_card_no"  dc:"身份证号"`
	IdCardExpiredAt   *gtime.Time `json:"id_card_expired_at"  dc:"身份证证有效期至,长期则为null"`
	LongTermEffective int         `json:"long_term_effective"  dc:"是否长期有效,1-是,2-否"`
	IdCardFrontPic    string      `json:"id_card_front_pic" dc:"身份证正面照片(头像面)"`
	IdCardBackPic     string      `json:"id_card_back_pic" dc:"身份证背面照片(国徽面)"`
}

// --------------------------我-是-分-割-线-------------------------- //

type CardLevel struct {
	gmeta.Meta `orm:"table:card_level"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
}

// --------------------------我-是-分-割-线-------------------------- //

type UpdateUserCardLevel struct {
	gmeta.Meta  `orm:"table:app_user"`
	Id          int `json:"id" dc:"用户id"`
	CardLevelId int `json:"card_level_id" dc:"卡等级id"`
}
