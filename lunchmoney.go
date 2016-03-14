package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/captncraig/temple"
	"github.com/gorilla/mux"
)

const DebitCreditPrecision int32 = 2

var devMode = flag.Bool("dev", false, "activate dev mode for templates")
var templateManager temple.TemplateStore

var accounts map[int64]*Account

func main() {
	setupConfig()
	setupTemplates()

	// Routing
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/accounts/{accountId:[0-9]+}", AccountList).Methods("GET")
	r.HandleFunc("/accounts/{accountId:[0-9]+}", AccountsUpdater).Methods("POST")
	r.HandleFunc("/accounts/import", AccountsImporter).Methods("POST")

	// Data
	accounts = make(map[int64]*Account)
	loadData()

	// Bind to a port and pass our router in
	listen := ":8000"
	log.Printf("Starting server on %s\n", listen)
	log.Printf("Template auto reloading: %t\n", *devMode)
	http.ListenAndServe(listen, r)
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

func YourHandler(w http.ResponseWriter, r *http.Request) {

	type PageData struct {
		Title     string
		Accounts  map[int64]*Account
		AccountId string
	}

	pageData := &PageData{Title: "Home", Accounts: accounts}

	err := templateManager.Execute(w, pageData, "index.html")
	if err != nil {
		log.Println(err)
	}
}

func AccountList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	type PageData struct {
		Title     string
		AccountId string
		Account   Account
		Accounts  map[int64]*Account
	}

	acc_id, err := strconv.ParseInt(vars["accountId"], 10, 64)
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

	pageData := &PageData{
		Title:     "Accounts",
		AccountId: vars["accountId"],
		Account:   account,
		Accounts:  accounts,
	}

	err = templateManager.Execute(w, pageData, "accounts_index.html")
	if err != nil {
		log.Println(err)
	}
}

func AccountsUpdater(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	acc_id, err := strconv.ParseInt(vars["accountId"], 10, 64)
	if err != nil {
		log.Printf("Invalid Account ID: %v - %v\n", acc_id, err)
		return
	}

	trans_id, err := strconv.ParseInt(r.PostFormValue("ID"), 10, 64)
	if err != nil {
		log.Printf("Invalid ID: %v - %v\n", trans_id, err)
		return
	}

	var transaction Transaction
	for _, trans := range accounts[acc_id].Transactions {
		if trans.ID == trans_id {
			transaction = trans
		}
	}

	transaction.Payee = r.PostFormValue("Payee")
	transaction.Memo = r.PostFormValue("Memo")
	debit := currencyToInt(r.PostFormValue("Debit"), accounts[acc_id])
	transaction.Debit = debit
	credit := currencyToInt(r.PostFormValue("Credit"), accounts[acc_id])
	transaction.Credit = credit

	for trans_key, trans := range accounts[acc_id].Transactions {
		if trans.ID == trans_id {
			accounts[acc_id].Transactions[trans_key] = transaction
			w.Header().Set("Content-Type", "text/json")

			jsonResponse, err := json.Marshal(transaction)
			if err != nil {
				log.Printf("Json marshaling error: %v\n", err)
				return
			}
			w.Write(jsonResponse)
		}
	}
}

func AccountsImporter(w http.ResponseWriter, r *http.Request) {
	acc_id, err := strconv.ParseInt(r.FormValue("account_id"), 10, 64)

	if err != nil {
		log.Println(err)
		return
	}

	file, _, err := r.FormFile("upload_file")

	if err != nil {
		log.Println(err)
	}

	accounts[acc_id].ImportTransactions(file)
	log.Println(accounts[acc_id].Transactions)

	http.Redirect(w, r, fmt.Sprintf("/accounts/%d", acc_id), 302)
}

func loadData() {
	acc1 := &Account{
		ID:                 1,
		Name:               "Savings",
		CurrencySymbolLeft: "$",
		CurrencyCode:       "AUD",
		DecimalPlaces:      DebitCreditPrecision,
		IsActive:           true,
		Icon:               "dollar",
	}

	trans := Transaction{
		ID:        1,
		Credit:    0,
		Debit:     10000,
		IsCleared: true,
		Payee:     "Supermarket",
	}
	acc1.AddTransaction(trans)

	acc2 := &Account{
		ID:                 2,
		Name:               "Credit Card",
		CurrencyCode:       "AUD",
		CurrencySymbolLeft: "$",
		DecimalPlaces:      DebitCreditPrecision,
		IsActive:           true,
		Icon:               "credit-card",
	}

	accounts[1] = acc1
	accounts[2] = acc2
}
