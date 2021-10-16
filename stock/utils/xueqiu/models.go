package xueqiu

type QuoteResp struct {
	Market Market `json:"market"`
	Quote  Quote  `json:"quote"`
}

type Market struct {
	Region   string `json:"region"`
	TimeZone string `json:"time_zone"`
}

type Quote struct {
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	AvgPrice  float32 `json:"avg_price"`
	Symbol    string  `json:"symbol"`
	High      float32 `json:"high"`
	Current   float32 `json:"current"`
	Low       float32 `json:"low"`
	Currency  string  `json:"currency"`
	LimitUp   float32 `json:"limit_up"`
	Open      float32 `json:"open"`
	LimitDown float32 `json:"limit_down"`
	Percent   float32 `json:"percent"`
}
