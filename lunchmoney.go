package main

import (
	"html/template"
	"log"
	"net/http"
	// "bytes"

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

	var account Account
	for _, acc := range accounts {
		if acc.ID == vars["accountId"] {
			account = acc
			break
		}
	}

	pageData := &PageData{Title: "Accounts", AccountId: vars["accountId"], Account: account, Accounts: accounts}

	err := templates.ExecuteTemplate(w, "accounts_index.html", pageData)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/accounts/{accountId:[0-9]+}", AccountList)

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
		ID:                 "1",
		Name:               "Savings",
		CurrencySymbolLeft: "$",
		Currency:           "AUD",
		IsActive:           true,
		Icon:               "dollar",
	}

	trans := Transaction{Credit: 0, Debit: 10000, IsCleared: true, Payee: "Supermarket"}
	acc1.AddTransaction(trans)

	acc2 := Account{
		ID:                 "2",
		Name:               "Credit Card",
		Currency:           "AUD",
		CurrencySymbolLeft: "$",
		IsActive:           true,
		Icon:               "credit-card",
	}
	accounts = append(accounts, acc1, acc2)
}
