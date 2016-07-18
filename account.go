package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
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
	clearedTotal        decimal.Decimal
	Total               decimal.Decimal
}

func AccountList(c *gin.Context) {
	acc_id, err := strconv.ParseInt(c.Param("accountId"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var account Account
	for _, acc := range accounts {
		if acc.ID == acc_id {
			account = *acc
			break
		}
	}

	c.HTML(http.StatusOK, "accounts_index.html", gin.H{
		"Title":     "Accounts",
		"AccountId": acc_id,
		"Account":   &account,
		"Accounts":  &accounts,
		"Shares":    &shares,
		"NetWorth":  netWorth,
	})
}

func AccountsUpdater(c *gin.Context) {
	acc_id, err := strconv.ParseInt(c.Param("accountId"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	trans_id, err := strconv.ParseInt(c.PostForm("ID"), 10, 64)
	if err != nil {
		log.Printf("Invalid ID: %v - %v\n", trans_id, err)
		return
	}

	var account *Account
	for _, acc := range accounts {
		if acc.ID == acc_id {
			account = acc
			break
		}
	}

	var transaction Transaction
	for _, trans := range account.Transactions {
		if trans.ID == trans_id {
			transaction = trans
		}
	}

	transaction.Payee = c.PostForm("Payee")
	transaction.Memo = c.PostForm("Memo")
	debit, err := decimal.NewFromString(c.PostForm("Debit"))
	if err != nil {
		log.Printf("Invalid Debit %v\n", err)
	} else {
		transaction.Debit = debit
	}
	credit, err := decimal.NewFromString(c.PostForm("Credit"))
	if err != nil {
		log.Printf("Invalid Credit %v\n", err)
	} else {
		transaction.Credit = credit
	}

	for trans_key, trans := range account.Transactions {
		if trans.ID == trans_id {
			account.Transactions[trans_key] = transaction

			jsonResponse, err := json.Marshal(transaction)
			if err != nil {
				log.Printf("Json marshaling error: %v\n", err)
				return
			}
			c.JSON(http.StatusOK, jsonResponse)
		}
	}
}

func AccountsImporter(c *gin.Context) {
	acc_id, err := strconv.ParseInt(c.PostForm("account_id"), 10, 64)

	if err != nil {
		log.Println(err)
		return
	}

	file, _, err := c.Request.FormFile("upload_file")

	if err != nil {
		log.Println(err)
	}

	accounts[acc_id].ImportTransactions(file)
	// log.Println(accounts[acc_id].Transactions)

	c.Redirect(http.StatusFound, fmt.Sprintf("/accounts/%d", acc_id))
}

func (a *Account) GetClearedTotal() decimal.Decimal {
	return a.clearedTotal
}

func (a *Account) GetTotal() decimal.Decimal {
	return a.Total
}

func (a *Account) GetFormattedTotal() string {
	return a.Total.StringFixed(a.DecimalPlaces)
}

func (a *Account) GetFormattedAmount(amount decimal.Decimal) string {
	return amount.StringFixed(a.DecimalPlaces)
}

func (a *Account) GetFormattedClearedTotal() string {
	return a.clearedTotal.StringFixed(a.DecimalPlaces)
}

func (a *Account) LoadTransactions() {
	// log.Printf("Loading transactions for account id %d\n", a.ID)

	query := `SELECT id, occurred, payee, memo, debit, credit
	FROM transactions
	WHERE account_id = $1`

	var id int64
	var credit, debit decimal.Decimal
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

		a.Total = a.Total.Add(credit)
		a.Total = a.Total.Sub(debit)

		if t.IsCleared {
			a.clearedTotal = a.clearedTotal.Add(credit)
			a.clearedTotal = a.clearedTotal.Sub(debit)
		}

		a.Transactions = append(a.Transactions, *t)
	}

}

func (a *Account) GetTransactions() []Transaction {
	return a.Transactions
}

func (a *Account) AddTransaction(t Transaction) {
	a.Transactions = append(a.Transactions, t)

	a.Total = a.Total.Add(t.Credit)
	a.Total = a.Total.Sub(t.Debit)

	if t.IsCleared {
		a.clearedTotal = a.clearedTotal.Add(t.Credit)
		a.clearedTotal = a.clearedTotal.Sub(t.Debit)
	}
}

func (a *Account) Create(db *sql.DB) {
	query := `INSERT INTO accounts (name, currency_code, currency_symbol_left, currency_symbol_right, decimal_places, icon, is_active, cleared_total, total) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: %s\n", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		a.Name,
		a.CurrencyCode,
		a.CurrencySymbolLeft,
		a.CurrencySymbolRight,
		a.DecimalPlaces,
		a.Icon,
		a.IsActive,
		a.GetClearedTotal(),
		a.GetTotal(),
	)
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

		transDate, err := time.Parse("02/01/2006", record[1])
		if err != nil {
			log.Println(err)
		}

		cleared := true
		reconciled := false

		credit, err := decimal.NewFromString(record[4])
		if err != nil {
			log.Print(err)
		}
		debit, err := decimal.NewFromString(record[3])
		if err != nil {
			log.Print(err)
		}

		trans := Transaction{
			Credit:       credit,
			Debit:        debit,
			IsCleared:    cleared,
			IsReconciled: reconciled,
			Payee:        record[2],
			Date:         transDate,
		}
		a.AddTransaction(trans)
	}
}
