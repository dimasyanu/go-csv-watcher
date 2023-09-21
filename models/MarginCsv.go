package models

import (
	"time"
)

type MarginCsv struct {
	AbCode            string
	Date              time.Time
	TradingId         int64
	OrderNumber       int64
	TransactionNumber int64
	DateTime          time.Time
	BuySell           rune
	SecurityCode      string
	Board             string
	Price             float64
	Quantity          int64
	Value             float64
	IsMargin          bool
}
