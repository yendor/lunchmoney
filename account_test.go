package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestTransactions(t *testing.T) {

	a := &Account{IsActive: true, CurrencyCode: "AUD"}

	credit10, _ := decimal.NewFromString("10.0")
	debit, _ := decimal.NewFromString("0.0")

	trans := Transaction{
		Credit:    credit10,
		Debit:     debit,
		IsCleared: true,
	}

	a.AddTransaction(trans)

	total := a.GetTotal()

	if total != credit10 {
		t.Errorf("Total not right, expected %v got %v", credit10, a.GetTotal())
	}
	if a.GetClearedTotal() != credit10 {
		t.Error("Cleared Total not right")
	}

	credit, _ := decimal.NewFromString("0.0")
	debit, _ = decimal.NewFromString("5.0")

	trans.Credit = credit
	trans.Debit = debit
	trans.IsCleared = false
	a.AddTransaction(trans)

	if a.GetTotal() != debit {
		t.Error("Total not right")
	}

	if a.GetClearedTotal() != credit10 {
		t.Error("Cleared total not right")
	}

	tr := a.GetTransactions()

	if len(tr) != 2 {
		t.Error("Number of transactions not right")
	}
	// log.Printf("%v\n", tr)
}
