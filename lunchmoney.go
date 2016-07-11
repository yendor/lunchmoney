package main

import (
	"database/sql"
	"flag"
	"log"

	//_ "github.com/mattn/go-sqlite3"

	_ "github.com/lib/pq"
	"github.com/shopspring/decimal"
)

const DebitCreditPrecision int32 = 2

var accounts map[int64]*Account

var shares map[string]*Share

var netWorth decimal.Decimal

var db *sql.DB
var err error

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	// Data
	accounts = make(map[int64]*Account)
	shares = make(map[string]*Share)

	setupConfig()

	setupDB()

	loadData()

	r := setupRouting()

	// Bind to a port and pass our router in
	listen := ":8000"
	log.Printf("Starting server on %s\n", listen)
	r.Run(listen)
}

func setupDB() {
	// db, err = sql.Open("sqlite3", "./lunchmoney.db")
	db, err = sql.Open("postgres", "")
	if err != nil {
		log.Printf("Error connect to the db: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the db: %s\n", err)
	}
}

func setupConfig() {
	flag.Parse()
}

func loadData() {
	query := `SELECT
		id,
		name,
		currency_code,
		currency_symbol_left,
		currency_symbol_right,
		decimal_places,
		icon,
		is_active,
		cleared_total,
		total
	FROM accounts`

	var id int64
	var cleared_total, total decimal.Decimal
	var decimal_places int32
	var name, currency_code, currency_symbol_left, currency_symbol_right, icon string
	var is_active bool

	rows, err := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&id,
			&name,
			&currency_code,
			&currency_symbol_left,
			&currency_symbol_right,
			&decimal_places,
			&icon,
			&is_active,
			&cleared_total,
			&total,
		)
		if err != nil {
			log.Printf("%s\n", err)
			return
		}

		acc := &Account{
			ID:                  id,
			Name:                name,
			CurrencyCode:        currency_code,
			CurrencySymbolLeft:  currency_symbol_left,
			CurrencySymbolRight: currency_symbol_right,
			DecimalPlaces:       decimal_places,
			Icon:                icon,
			IsActive:            is_active,
			clearedTotal:        cleared_total,
			Total:               total,
		}

		acc.LoadTransactions()

		netWorth = netWorth.Add(acc.GetTotal())
		// log.Printf("%v\n", acc)
		accounts[acc.ID] = acc
	}
	err = rows.Err()
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	query = `SELECT stockcode, qty FROM shares`

	var stockcode string
	var qty int64

	rows, err = db.Query(query)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&stockcode,
			&qty,
		)
		if err != nil {
			log.Printf("%s\n", err)
			return
		}

		share := &Share{
			Code: stockcode,
			Qty:  qty,
		}
		share.GetValue()
		// log.Println(share)

		shares[stockcode] = share

		netWorth = netWorth.Add(share.Value)
	}

}
