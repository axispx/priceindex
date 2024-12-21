package model

type PriceResponse struct {
	Price     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
	Source    string `json:"source"`
	Address   string `json:"address"`
}

type HourlyPriceResponse struct {
	Timestamp int64  `json:"timestmap"`
	Address   string `json:"address"`
	Price     string `json:"price"`
}

type DailyPriceResponse struct {
	Timestamp int64  `json:"timestmap"`
	Address   string `json:"address"`
	Price     string `json:"price"`
}
