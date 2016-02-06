package main

import "io"

// https://en.wikipedia.org/wiki/Quicken_Interchange_Format

type QifTransaction struct {
	Date  string
	Total float64
	Payee string
	Memo  string
}

type QifImport struct {
	Type         string
	Transactions []QifTransaction
}

func (i *QifImport) Import(r io.Reader) {

}
