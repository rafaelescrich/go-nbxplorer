package bitcoin

import (
	"log"

	"go-nbxplorer/config"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
)

var Client *rpcclient.Client

func InitBitcoinRPC() {
	connConfig := &rpcclient.ConnConfig{
		Host:         config.AppConfig.BTCRPCURL + ":" + config.AppConfig.BTCNodePort,
		User:         "yourrpcuser",
		Pass:         "yourrpcpassword",
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	var err error
	Client, err = rpcclient.New(connConfig, nil)
	if err != nil {
		log.Fatalf("Failed to connect to Bitcoin RPC: %v", err)
	}
	log.Println("Connected to Bitcoin RPC")
}

func GetBlockchainInfo() (*btcjson.GetBlockChainInfoResult, error) {
	info, err := Client.GetBlockChainInfo()
	if err != nil {
		return nil, err
	}
	return info, nil
}

func RescanBlocks(startBlockHashes []chainhash.Hash) ([]btcjson.RescannedBlock, error) {
	log.Printf("Starting block rescan from blocks: %v", startBlockHashes)
	return Client.RescanBlocks(startBlockHashes)
}
