package main

type Account struct {
	ID                  int64
	Name                string
	Currency            string
	CurrencySymbolLeft  string
	CurrencySymbolRight string
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
