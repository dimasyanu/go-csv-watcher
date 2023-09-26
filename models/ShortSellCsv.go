package models

import (
	"time"
)

type ShortSellCsv struct {
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
	IsShortSell       bool
}

func CreateShortSellCsv() *ShortSellCsv {
	return &ShortSellCsv{}
}

func (o ShortSellCsv) GetDateStr() string {
	return o.Date.Format(time.RFC3339)
}
