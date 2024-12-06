package common

// 利率
type BankRateResponse struct {
	Result struct {
		BaseCurrency string             `json:"baseCurrency"`
		Rates        map[string]float64 `json:"rates"`
	} `json:"result"`
}

type CardInfoRequest struct {
	CardNoList []string `json:"card_no_list"`
}

type CardInfoRow struct {
	CardNo             string  `json:"card_no"`
	EmbossedName       string  `json:"embossed_name"`
	Status             int     `json:"status"`
	ExpiryMonth        string  `json:"expiryMonth"`
	ExpiryYear         string  `json:"expiryYear"`
	CardLimit          float64 `json:"cardLimit"`
	DailyAtmLimit      float64 `json:"dailyAtmLimit"`
	DailyPurchaseLimit float64 `json:"dailyPurchaseLimit"`
	AvailableLimit     float64 `json:"availableLimit"`
	Balance            float64 `json:"balance"`
	Available          float64 `json:"available"`
}

type CardInfoResponse struct {
	Result struct {
		CardInfoList []*CardInfoRow `json:"card_info_list"`
	} `json:"result"`
}
