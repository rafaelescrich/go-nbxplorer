package handlers

import (
	"net/http"

	"go-nbxplorer/postgres"

	"github.com/gin-gonic/gin"
)

func GetUnusedAddress(c *gin.Context) {
	derivationScheme := c.Param("derivationScheme")

	address, err := postgres.GetUnusedAddressForScheme(postgres.DB, derivationScheme)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address})
}
