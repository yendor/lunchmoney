package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const DebitCreditPrecision int = 2

var templates *template.Template

var accounts []Account

func YourHandler(w http.ResponseWriter, r *http.Request) {

	type PageData struct {
		Title    string
		Accounts []Account
	}

	pageData := &PageData{Title: "Home", Accounts: accounts}

	err := templates.ExecuteTemplate(w, "index.html", pageData)
	if err != nil {
		log.Fatal(err)
	}
	// w.Write(doc)
}

func AccountList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	type PageData struct {
		Title     string
		AccountId string
		Account   Account
		Accounts  []Account
	}

	acc_id, err := strconv.ParseInt(vars["accountId"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var account Account
	for _, acc := range accounts {
		if acc.ID == acc_id {
			account = acc
			break
		}
	}

	pageData := &PageData{Title: "Accounts", AccountId: vars["accountId"], Account: account, Accounts: accounts}

	err = templates.ExecuteTemplate(w, "accounts_index.html", pageData)
	if err != nil {
		log.Fatal(err)
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

	var account Account
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

	transaction.Payee = r.PostFormValue("Payee")
	transaction.Memo = r.PostFormValue("Memo")
	debit, err := strconv.ParseInt(r.PostFormValue("Debit"), 10, 64)
	if err != nil {
		log.Printf("Invalid Debit %v\n", err)
	} else {
		transaction.Debit = debit
	}
	credit, err := strconv.ParseInt(r.PostFormValue("Credit"), 10, 64)
	if err != nil {
		log.Printf("Invalid Credit %v\n", err)
	} else {
		transaction.Credit = credit
	}

	for trans_key, trans := range account.Transactions {
		if trans.ID == trans_id {
			account.Transactions[trans_key] = transaction
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

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/accounts/{accountId:[0-9]+}", AccountList).Methods("GET")
	r.HandleFunc("/accounts/{accountId:[0-9]+}", AccountsUpdater).Methods("POST")

	var err error
	// var doc []byte
	templates, err = template.ParseGlob("templates/*.html")

	if err != nil {
		log.Fatal(err)
	}

	loadData()

	// Bind to a port and pass our router in
	listen := ":8000"
	log.Printf("Starting server on %s\n", listen)
	http.ListenAndServe(listen, r)
}

func loadData() {
	acc1 := Account{
		ID:                 1,
		Name:               "Savings",
		CurrencySymbolLeft: "$",
		Currency:           "AUD",
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

	acc2 := Account{
		ID:                 2,
		Name:               "Credit Card",
		Currency:           "AUD",
		CurrencySymbolLeft: "$",
		IsActive:           true,
		Icon:               "credit-card",
	}
	accounts = append(accounts, acc1, acc2)
}
