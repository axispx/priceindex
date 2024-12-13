package model

import "time"

type Price struct {
	Timestamp time.Time `gorm:"primaryKey;->"`
	Price     string
	Source    string
	Address   string
}
