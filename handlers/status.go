package handlers

import (
	"net/http"

	"go-nbxplorer/bitcoin"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	info, err := bitcoin.Client.GetBlockChainInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status := map[string]interface{}{
		"chain":                info.Chain,
		"blocks":               info.Blocks,
		"headers":              info.Headers,
		"bestblockhash":        info.BestBlockHash,
		"difficulty":           info.Difficulty,
		"mediantime":           info.MedianTime,
		"verificationprogress": info.VerificationProgress,
		"initialblockdownload": info.InitialBlockDownload,
	}

	c.JSON(http.StatusOK, status)
}
