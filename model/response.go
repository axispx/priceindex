package model

type PriceResponse struct {
	Price     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
	Source    string `json:"source"`
	Address   string `json:"address"`
}
