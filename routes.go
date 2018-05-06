package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouting() *gin.Engine {
	// Routing
	//
	r := gin.Default()

	// html := template.Must(template.ParseGlob("templates/*"))
	// r.SetHTMLTemplate(html)

	r.LoadHTMLGlob("templates/*")
	r.Static("/public", "./public")
	// Routes consist of a path and a handler function.
	//
	r.GET("/", HomepageHandler)
	r.GET("/interest", InterestList)
	r.GET("/shares", SharesHandler)
	r.GET("/budget", BudgetHandler)
	r.GET("/accounts/:accountId", AccountList)
	r.POST("/accounts/:accountId", AccountsUpdater)
	r.POST("/import", AccountsImporter)
	return r
}

// HomepageHandler is the http handler for the homepage
func HomepageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":    "Home",
		"Accounts": &accounts,
		"Shares":   &shares,
		"NetWorth": netWorth,
	})
}
