package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/captncraig/temple"
)

const DebitCreditPrecision int32 = 2

var devMode = flag.Bool("dev", true, "activate dev mode for templates")
var templateManager temple.TemplateStore

var accounts map[int64]*Account

var shares map[string]*Share

var db *sql.DB
var err error

func main() {
	// Data
	accounts = make(map[int64]*Account)
	shares = make(map[string]*Share)

	setupConfig()
	setupTemplates()

	setupDB()

	loadData()

	if len(accounts) == 0 {
		createAccounts()
	}

	r := setupRouting()

	// Bind to a port and pass our router in
	listen := ":8000"
	log.Printf("Starting server on %s\n", listen)
	log.Printf("Template auto reloading: %t\n", *devMode)
	http.ListenAndServe(listen, r)
}

func setupDB() {
	db, err = sql.Open("sqlite3", "./lunchmoney.db")
	if err != nil {
		log.Printf("Error connect to the db: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the db: %s\n", err)
	}

	// goose migrate here
}

func setupConfig() {
	flag.Parse()
}

func setupTemplates() {
	var err error

	// Templates
	templateManager, err = temple.New(*devMode, templates, "templates")
	if err != nil {
		log.Fatal(err)
	}
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

	var id, cleared_total, total int64
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
			total:               total,
		}
		log.Printf("%v\n", acc)

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
		log.Println(share)

		shares[stockcode] = share
	}

}

func createAccounts() {

	acc1 := &Account{
		Name:               "Savings",
		CurrencySymbolLeft: "$",
		CurrencyCode:       "AUD",
		DecimalPlaces:      DebitCreditPrecision,
		IsActive:           true,
		Icon:               "dollar",
	}
	acc1.Create(db)

	acc2 := &Account{
		Name:               "Credit Card",
		CurrencyCode:       "AUD",
		CurrencySymbolLeft: "$",
		DecimalPlaces:      DebitCreditPrecision,
		IsActive:           true,
		Icon:               "credit-card",
	}
	acc2.Create(db)

	accounts[acc1.ID] = acc1
	accounts[acc2.ID] = acc2
}
