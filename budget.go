package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Budget is the top level budget struct
type Budget struct {
	Year       int
	Month      int
	Categories []BudgetCategory
}

// BudgetCategory is a category for assigning transactions in a budget
type BudgetCategory struct {
}

// NewBudget create a new budget and return the struct
func (b *Budget) NewBudget() *Budget {
	categories := make([]BudgetCategory, 0)

	return &Budget{
		Year:       2016,
		Month:      5,
		Categories: categories,
	}
}

// BudgetHandler http handler for the budget list page
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
