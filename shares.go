package main

import (
	"log"
	"net/http"

	"github.com/FlashBoys/go-finance"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type Share struct {
	Code  string
	Qty   int64
	Value decimal.Decimal
}

func (s *Share) GetValue() {
	// 15-min delayed full quote for Apple.
	q, err := finance.GetQuote(s.Code)
	if err != nil {
		log.Println(q)
	}

	qty := decimal.New(s.Qty, 0)

	s.Value = q.LastTradePrice.Mul(qty)
}

func SharesHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "shares_index.html", gin.H{
		"Title":    "Shares",
		"Shares":   &shares,
		"Accounts": &accounts,
		"NetWorth": netWorth,
	})
}
