package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

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

var configfile string

var cfg Config

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
	// db, err = sql.Open("sqlite3", "./lunchmoney.db")
	db, err = sql.Open("postgres", cfg.DSN)
	if err != nil {
		log.Printf("Error connect to the db: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the db: %s\n", err)
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
