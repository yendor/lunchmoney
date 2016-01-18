package main

import "time"

type Transaction struct {
	ID           int
	Date         time.Time
	Payee        string
	Memo         string
	Debit        int
	Credit       int
	CategoryID   int
	AccountID    int
	IsCleared    bool
	IsReconciled bool
}
