package main

import (
	"encoding/csv"
	"io"
	"log"
	"time"
)

type Account struct {
	ID                  int64
	Name                string
	CurrencyCode        string
	CurrencySymbolLeft  string
	CurrencySymbolRight string
	DecimalPlaces       int32
	Icon                string
	IsActive            bool
	Transactions        []Transaction
	clearedTotal        int64
	total               int64
}

func (a *Account) GetClearedTotal() int64 {
	return a.clearedTotal
}

func (a *Account) GetTotal() int64 {
	return a.total
}

func (a *Account) GetTransactions() []Transaction {
	return a.Transactions
}

func (a *Account) AddTransaction(t Transaction) {
	a.Transactions = append(a.Transactions, t)
	a.total += t.Credit
	a.total -= t.Debit

	if t.IsCleared {
		a.clearedTotal += t.Credit
		a.clearedTotal -= t.Debit
	}
}

func (a *Account) ImportTransactions(r io.Reader) {

	csvr := csv.NewReader(r)
	csvr.LazyQuotes = true
	csvr.TrimLeadingSpace = true

	records, err := csvr.ReadAll()
	if err != nil {
		log.Println(err)
	}

	for k, record := range records {
		if k == 0 {
			continue
		}

		transDate, err := time.Parse("02/01/2006", record[3])
		if err != nil {
			log.Println(err)
		}

		cleared := false
		reconciled := false

		if record[11] == "R" {
			cleared = true
			reconciled = true
		} else if record[11] == "C" {
			cleared = true
		}

		credit := currencyToInt(record[10], a)
		debit := currencyToInt(record[9], a)

		trans := Transaction{
			Credit:       credit,
			Debit:        debit,
			IsCleared:    cleared,
			IsReconciled: reconciled,
			Payee:        record[4],
			Date:         transDate,
		}
		a.AddTransaction(trans)
	}
}
