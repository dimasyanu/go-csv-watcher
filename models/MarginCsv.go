package models

import (
	"time"
)

type MarginCsv struct {
	AbCode            string
	Date              time.Time
	TradingId         string
	OrderNumber       int64
	TransactionNumber int64
	DateTime          time.Time
	BuySell           string
	SecurityCode      string
	Board             string
	Price             float64
	Quantity          int64
	Value             float64
	IsMargin          bool
}

func CreateMarginCsv() *MarginCsv {
	return &MarginCsv{}
}

func (o MarginCsv) GetDateStr() string {
	return o.Date.Format(time.RFC3339)
}
