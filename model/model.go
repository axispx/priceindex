package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Price struct {
	Timestamp time.Time       `gorm:"primaryKey;->"`
	Price     decimal.Decimal `gorm:"type:decimal(30,30);"`
	Source    string
	Address   string
}

type HourlyPrice struct {
	Hour     time.Time       `json:"hour"`
	Address  string          `json:"address"`
	AvgPrice decimal.Decimal `json:"avg_price"`
}

type DailyPrice struct {
	Day      time.Time       `json:"day"`
	Address  string          `json:"address"`
	AvgPrice decimal.Decimal `json:"avg_price"`
}
