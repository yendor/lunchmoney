package main

import "testing"

func TestTransactions(t *testing.T) {

	a := &Account{IsActive: true, Currency: "AUD"}

	trans := Transaction{Credit: 10, Debit: 0, IsCleared: true}

	a.AddTransaction(trans)

	if a.GetTotal() != 10 {
		t.Fail()
	}
	if a.GetClearedTotal() != 10 {
		t.Fail()
	}

	trans.Credit = 0
	trans.Debit = 5
	trans.IsCleared = false
	a.AddTransaction(trans)

	if a.GetTotal() != 5 {
		t.Fail()
	}

	if a.GetClearedTotal() != 10 {
		t.Fail()
	}

	tr := a.GetTransactions()

	if len(tr) != 2 {
		t.Fail()
	}
	// log.Printf("%v\n", tr)
}
