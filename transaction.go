package main

import "time"

type Transaction struct {
	ID           int64
	Date         time.Time
	Payee        string
	Memo         string
	Debit        int64
	Credit       int64
	CategoryID   int64
	AccountID    int64
	IsCleared    bool
	IsReconciled bool
}
