package main

type Account struct {
	ID           int
	Name         string
	Currency     string
	IsActive     bool
	Transactions []Transaction
	clearedTotal int
	total        int
}

func (a *Account) GetClearedTotal() int {
	return a.clearedTotal
}

func (a *Account) GetTotal() int {
	return a.total
}

func (a *Account) GetList() []Account {
	var ret []Account
	newAccount := Account{ID: 1, Name: "Savings", Currency: "AUD", IsActive: true}
	ret = append(ret, newAccount)
	return ret
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
