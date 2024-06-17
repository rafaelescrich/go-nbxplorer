package handlers

import (
	"encoding/json"
	"go-nbxplorer/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeeRateResponse struct {
	FastestFee  int `json:"fastestFee"`
	HalfHourFee int `json:"halfHourFee"`
	HourFee     int `json:"hourFee"`
}

func GetFeeRate(c *gin.Context) {
	url := config.AppConfig.BTCRPCURL + "/fee-rate"
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var feeRateResponse FeeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&feeRateResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fee rate": feeRateResponse})
}
