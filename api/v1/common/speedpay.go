package common

import "github.com/gogf/gf/v2/frame/g"

type GetCountryListReq struct {
	g.Meta `path:"/country/list"  summary:"获取国家列表" tags:"国家" method:"get"`
}

type GetCountryListRes struct {
	Result []Country `json:"result"`
}

type Country struct {
	TwoLetterCode          string `json:"two_letter_code"`
	ThreeLetterCode        string `json:"three_letter_code"`
	InitialLetter          string `json:"initial_letter"`
	NumberCode             int    `json:"number_code"`
	IsoCode                string `json:"iso_code"`
	EnglishName            string `json:"english_name"`
	ChineseName            string `json:"chinese_name"`
	ChineseTraditionalName string `json:"chinese_traditional_name"`
	PhoneCode              string `json:"phone_code"`
	SupportedCardType      []int  `json:"supported_card_type"`
}

type ApplyCallbackReq struct {
	g.Meta     `path:"/apply/callback"  summary:"申请回调" tags:"申请" method:"post"`
	EventId    string `json:"event_id"`
	RelationId string `json:"relation_id"`
	CardNo     string `json:"card_no"`
}

type ApplyCallbackRes struct {
}
