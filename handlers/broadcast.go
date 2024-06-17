package handlers

import (
	"bytes"
	"encoding/hex"
	"net/http"

	"go-nbxplorer/bitcoin"

	"github.com/btcsuite/btcd/wire"
	"github.com/gin-gonic/gin"
)

func BroadcastTransaction(c *gin.Context) {
	var req struct {
		Hex string `json:"hex"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txBytes, err := hex.DecodeString(req.Hex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tx wire.MsgTx
	err = tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := bitcoin.Client.SendRawTransaction(&tx, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"txid": txHash.String()})
}
