# go-nbxplorer

`go-nbxplorer` is a Go-based application designed to provide functionalities similar to NBXplorer for Bitcoin. This application uses Postgres for data storage, Gin Gonic for the web framework, Sentry for error tracking, and RabbitMQ for messaging.

## Features

- Track derivation schemes
- Fetch fee rates
- Scan UTXO set
- Connect to Bitcoin RPC and node
- Swagger for API documentation
- Sentry for error logging
- Dockerized for easy deployment

## Requirements

- Go 1.21 or higher
- Postgres database
- Bitcoin node with RPC enabled
- RabbitMQ server

## Environment Variables

Create a `.env` file in the root directory with the following environment variables:

```env
RABBITMQ_URL=amqp://user:password@localhost:5672/
SENTRY_DSN=<your_sentry_dsn>
BTC_RPC_URL=<your_bitcoin_rpc_url>
BTC_RPC_USER=<your_bitcoin_rpc_user>
BTC_RPC_PASS=<your_bitcoin_rpc_password>
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=nbxplorer
```

## Running Locally 

1. Clone the repository:


```sh
git clone <repository_url>
cd go-nbxplorer
```

1. Install dependencies:


```sh
go mod tidy
```

1. Run the application:


```sh
go run ./main.go
```

## Docker 

1. Build the Docker image:


```sh
docker build -t go-nbxplorer .
```

1. Run the Docker container:


```sh
docker run --env-file .env -p 8080:8080 go-nbxplorer
```

## API Documentation 
API documentation is generated using Swagger. Once the application is running, you can access the API docs at: `http://localhost:8080/swagger/index.html`
## Project Structure 


```python
.
├── main.go                    # Entry point of the application
├── config                     # Configuration related files
│   └── config.go
├── handlers                   # HTTP handlers
│   └── handlers.go
├── bitcoin                    # Bitcoin related functionalities
│   └── bitcoin.go
├── postgres                   # Postgres related functionalities
│   └── models.go
├── rabbitmq                   # RabbitMQ related functionalities
│   └── rabbitmq.go
├── logger                     # Logging related functionalities
│   └── logger.go
├── Dockerfile                 # Dockerfile for containerizing the application
├── go.mod                     # Go module file
├── go.sum                     # Go module dependencies
└── README.md                  # Project README
```

## Contributing 

Contributions are welcome! Please feel free to submit a Pull Request.

## License 

This project is licensed under the MIT License - see the LICENSE file for details.
