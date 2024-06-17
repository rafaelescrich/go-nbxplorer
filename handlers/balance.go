package handlers

import (
	"go-nbxplorer/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBalance gets the balance for a derivation scheme
func GetBalance(c *gin.Context) {
	derivationScheme := c.Param("derivationScheme")
	addresses, err := postgres.GetAddressesForScheme(postgres.DB, derivationScheme)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transactions, err := postgres.GetTransactionsForScheme(postgres.DB, derivationScheme)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate balance from transactions
	var balance float64
	for _, tx := range transactions {
		if tx.Confirmed {
			balance += tx.Amount
		}
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance, "addresses": addresses})
}
