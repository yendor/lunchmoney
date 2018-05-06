package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/shopspring/decimal"
)

// DebitCreditPrecision is the number of decimal places to display
const DebitCreditPrecision int32 = 2

var accounts map[int64]*Account

var shares map[string]*Share

var netWorth decimal.Decimal

var db *sql.DB
var err error

var configfile string

var cfg Config

// Config is the config for the lunchmoney app
type Config struct {
	Listen string
	DSN    string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	flag.Parse()
	configfile = flag.Arg(0)
	if configfile == "" {
		configfile = "config.json"
	}

	file, _ := os.Open(configfile)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Data
	accounts = make(map[int64]*Account)
	shares = make(map[string]*Share)

	setupDB()

	loadData()

	r := setupRouting()

	// Bind to a port and pass our router in
	listen := cfg.Listen
	log.Printf("Starting server on %s\n", listen)
	r.Run(listen)
}

func setupDB() {
	parts := strings.Split(cfg.DSN, "://")
	driver := parts[0]
	dsn := parts[1]

	// db, err = sql.Open("sqlite3", "./lunchmoney.db")
	db, err = sql.Open(driver, dsn)
	if err != nil {
		log.Fatalf("Error connecting to the db using driver %s at %s: %s", driver, dsn, err)
	}

	// if driver != "sqlite3" {
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the db using driver %s at %s: %s", driver, dsn, err)
	}
	// }
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
	var clearedTotal, total decimal.Decimal
	var decimalPlaces int32
	var name, currencyCode, currencySymbolLeft, currencySymbolRight, icon string
	var isActive bool

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatalf("Unable to load data: %s", err)
	}

	for rows.Next() {
		err := rows.Scan(
			&id,
			&name,
			&currencyCode,
			&currencySymbolLeft,
			&currencySymbolRight,
			&decimalPlaces,
			&icon,
			&isActive,
			&clearedTotal,
			&total,
		)
		if err != nil {
			log.Printf("%s\n", err)
			return
		}

		acc := &Account{
			ID:                  id,
			Name:                name,
			CurrencyCode:        currencyCode,
			CurrencySymbolLeft:  currencySymbolLeft,
			CurrencySymbolRight: currencySymbolRight,
			DecimalPlaces:       decimalPlaces,
			Icon:                icon,
			IsActive:            isActive,
			clearedTotal:        clearedTotal,
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
	if err != nil {
		log.Fatalf("Unable to do query: %s", err)
	}

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
