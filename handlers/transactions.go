package handlers

import (
	"net/http"

	"go-nbxplorer/postgres"

	"github.com/gin-gonic/gin"
)

func QueryTransactions(c *gin.Context) {
	derivationScheme := c.Param("derivationScheme")

	transactions, err := postgres.GetTransactionsForScheme(postgres.DB, derivationScheme)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
