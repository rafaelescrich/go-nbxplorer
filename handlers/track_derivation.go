package handlers

import (
	"bytes"
	"encoding/json"
	"go-nbxplorer/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddressInfo struct {
	Address string `json:"address"`
	Used    bool   `json:"used"`
}

type TrackDerivationSchemeRequest struct {
	DerivationScheme string `json:"derivation_scheme"`
}

func TrackDerivationScheme(c *gin.Context) {
	url := config.AppConfig.BTCRPCURL + "/track-derivation-scheme"
	jsonReq, err := json.Marshal(c.Param("derivationScheme"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var addressInfos []AddressInfo
	if err := json.NewDecoder(resp.Body).Decode(&addressInfos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": addressInfos})
}
