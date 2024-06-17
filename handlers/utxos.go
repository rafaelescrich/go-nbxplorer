package handlers

import (
	"net/http"

	"go-nbxplorer/bitcoin"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/gin-gonic/gin"
)

func ScanUTXOSet(c *gin.Context) {
	// Example block hash to start rescan
	startBlockHashStr := "0000000000000000000000000000000000000000000000000000000000000000" // Replace with a real block hash if needed
	startBlockHash, err := chainhash.NewHashFromStr(startBlockHashStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	startBlockHashes := []chainhash.Hash{*startBlockHash}

	rescannedBlocks, err := bitcoin.RescanBlocks(startBlockHashes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          "UTXO set scan completed",
		"rescannedBlocks": rescannedBlocks,
	})
}
