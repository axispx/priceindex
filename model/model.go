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
	Hour     time.Time       `gorm:"->"`
	Address  string          `gorm:"->"`
	AvgPrice decimal.Decimal `gorm:"->"`
}

type DailyPrice struct {
	Day      time.Time       `gorm:"->"`
	Address  string          `gorm:"->"`
	AvgPrice decimal.Decimal `gorm:"->"`
}
