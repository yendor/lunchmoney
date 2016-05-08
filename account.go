package main

import (
	"database/sql"
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
	clearedTotal        float64
	total               float64
}

func (a *Account) GetClearedTotal() float64 {
	return a.clearedTotal
}

func (a *Account) GetTotal() float64 {
	return a.total
}

func (a *Account) GetFormattedTotal() string {
	return currencyToStr(a.total, a)
}

func (a *Account) GetFormattedAmount(amount float64) string {
	return currencyToStr(amount, a)
}

func (a *Account) GetFormattedClearedTotal() string {
	return currencyToStr(a.clearedTotal, a)
}

func (a *Account) LoadTransactions() {
	// log.Printf("Loading transactions for account id %d\n", a.ID)

	query := `SELECT id, occurred, payee, memo, debit, credit
	FROM transactions
	WHERE account_id = $1`

	var id int64
	var credit, debit float64
	var payee, memoStr string
	var memo sql.NullString
	var occurred time.Time

	rows, err := db.Query(query, a.ID)
	if err != nil {
		log.Printf("QUERY ERROR: %s\n", err)
		return
	}
	defer rows.Close()

	if memo.Valid {
		memoStr = memo.String
	} else {
		memoStr = ""
	}

	for rows.Next() {
		err := rows.Scan(
			&id,
			&occurred,
			&payee,
			&memo,
			&debit,
			&credit,
		)
		if err != nil {
			log.Printf("%s\n", err)
			return
		}

		t := &Transaction{
			ID:        id,
			Date:      occurred,
			Payee:     payee,
			Memo:      memoStr,
			Debit:     debit,
			Credit:    credit,
			AccountID: a.ID,
		}

		a.total += t.Credit
		a.total -= t.Debit

		if t.IsCleared {
			a.clearedTotal += t.Credit
			a.clearedTotal -= t.Debit
		}

		a.Transactions = append(a.Transactions, *t)
	}

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

func (a *Account) Create(db *sql.DB) {
	query := `INSERT INTO accounts (name, currency_code, currency_symbol_left, currency_symbol_right, decimal_places, icon, is_active, cleared_total, total) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: %s\n", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(a.Name, a.CurrencyCode, a.CurrencySymbolLeft, a.CurrencySymbolRight, a.DecimalPlaces, a.Icon, a.IsActive, a.GetClearedTotal(), a.GetTotal())
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	a.ID, err = result.LastInsertId()
	if err != nil {
		log.Printf("Error: %s\n", err)
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
