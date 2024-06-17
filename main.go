package main

import (
	"time"

	"go-nbxplorer/config"
	_ "go-nbxplorer/docs" // Add this import to include the generated Swagger docs
	"go-nbxplorer/handlers"
	"go-nbxplorer/logger"
	"go-nbxplorer/postgres"
	"go-nbxplorer/rabbitmq"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title NBXplorer API
// @version 1.0
// @description This is the NBXplorer API

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	config.InitConfig()
	logger.InitLogger()

	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.AppConfig.SentryDSN,
	})
	if err != nil {
		logger.LogFatal("sentry.Init", err)
	}
	defer sentry.Flush(2 * time.Second)

	postgres.InitDB()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Connect to RabbitMQ
	conn, err := rabbitmq.ConnectToRabbitMQ(config.AppConfig.RabbitMQURL)
	if err != nil {
		logger.LogError("Failed to connect to RabbitMQ", err)
		return
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	channel, err := rabbitmq.CreateChannel(conn)
	if err != nil {
		logger.LogError("Failed to create RabbitMQ channel", err)
		return
	}
	defer channel.Close()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/v1/cryptos/:cryptoCode/status", handlers.GetStatus)
	r.POST("/v1/cryptos/:cryptoCode/derivations/:derivationScheme", handlers.TrackDerivationScheme)
	r.GET("/v1/cryptos/:cryptoCode/derivations/:derivationScheme/transactions", handlers.QueryTransactions)
	r.GET("/v1/cryptos/:cryptoCode/derivations/:derivationScheme/balance", handlers.GetBalance)
	r.GET("/v1/cryptos/:cryptoCode/derivations/:derivationScheme/addresses/unused", handlers.GetUnusedAddress)
	r.POST("/v1/cryptos/:cryptoCode/transactions", handlers.BroadcastTransaction)
	r.GET("/v1/cryptos/:cryptoCode/fees/:blockCount", handlers.GetFeeRate)
	r.POST("/v1/cryptos/:cryptoCode/derivations/:derivationScheme/utxos/scan", handlers.ScanUTXOSet)

	logger.Info.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Error.Fatalf("Error starting server: %v", err)
	}
}
