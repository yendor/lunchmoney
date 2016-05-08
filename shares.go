package main

import (
	"log"

	"github.com/FlashBoys/go-finance"
	"github.com/shopspring/decimal"
)

type Share struct {
	Code  string
	Qty   int64
	Value decimal.Decimal
}

func (s *Share) GetValue() {
	// 15-min delayed full quote for Apple.
	q, err := finance.GetQuote(s.Code)
	if err == nil {
		log.Println(q)
	}

	qty := decimal.New(s.Qty, 0)

	log.Println(qty)
	log.Println(q.LastTradePrice)

	s.Value = q.LastTradePrice.Mul(qty)
}
