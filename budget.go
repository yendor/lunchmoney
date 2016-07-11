package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Budget struct {
	Year       int
	Month      int
	Categories []BudgetCategory
}

type BudgetCategory struct {
}

func (b *Budget) NewBudget() *Budget {
	categories := make([]BudgetCategory, 0)

	return &Budget{
		Year:       2016,
		Month:      5,
		Categories: categories,
	}
}

func BudgetHandler(c *gin.Context) {
	var b Budget

	budget := b.NewBudget()

	c.HTML(http.StatusOK, "budget.html", gin.H{
		"Title":    "Budget",
		"Budget":   budget,
		"Accounts": &accounts,
		"Shares":   &shares,
		"NetWorth": netWorth,
	})
}
