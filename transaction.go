package main

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID           int64
	Date         time.Time
	Payee        string
	Memo         string
	Debit        decimal.Decimal
	Credit       decimal.Decimal
	CategoryID   int64
	AccountID    int64
	IsCleared    bool
	IsReconciled bool
}
