package model

import "time"

type Price struct {
	Price     string
	Timestamp time.Time
	Source    string
	Address   string
}
