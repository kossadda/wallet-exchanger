# Wallet Exchanger gRPC Service

This service provides an API to fetch the current exchange rates for currencies. It operates through gRPC and can be integrated with other services for currency exchange.

The service provides two main methods:

1. **GetExchangeRates** — returns the current exchange rates for all currencies.
2. **GetExchangeRateForCurrency** — returns the exchange rate between two specific currencies.

# Contents

1. [Wallet Exchanger gRPC Service](#wallet-exchanger-grpc-service)
2. [Installation and Setup](#installation-and-setup) \
   2.1. [Dependencies](#dependencies) \
   2.2. [When importing the package](#when-importing-the-package) \
   2.3. [When cloning the repository](#when-cloning-the-repository) \
   2.4. [Build via docker-compose](#build-via-docker-compose)
3. [API](#api) \
   3.1. [Fetching all currency exchange rates](#fetching-all-currency-exchange-rates) \
   3.2. [Fetching exchange rates for specific currencies](#fetching-exchange-rates-for-specific-currencies)
4. [Testing](#testing)
5. [Documentation](#documentation)

## Installation and Setup

### Dependencies

For currency exchange operations to work, you need to deploy servers for `exchanger` and `redis`, and also specify their addresses in the config

The required `postgres` database relations are specified in the files of the `db/migrations` directory in the root of the repository

### When importing the package

Specify the startup parameters in `configs.ServerConfig`

**Environment Variables**
- For `Postgres`:
```env
DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE
```
- For servers:
```env
APP_ENV  // (local, dev, prod)
APP_PORT, APP_HOST, GRPC_PORT, GRPC_HOST, CACHE_PORT, CACHE_HOST
```

Example of running the service by specifying the path to the `env` configuration or creating configs manually:

```go
package main

import (
	"github.com/kossadda/wallet-exchanger/exchanger/pkg/app"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
)

func main() {
	// Можно указать путь до конфигов
	//servConf := configs.NewServerEnvConfig("config/local.env")
	//dbConf := configs.NewEnvConfigDB("config/database.env")

	// Или забить их вручную
	servConf := &configs.ServerConfig{
		Env:         "local",
		Servers: map[string]configs.Server{
			"APP": configs.Server{
				Host: "localhost",
				Port: "44044",
			},
		},
	}

	log := logger.SetupByEnv(servConf.Env)
	dbConf := &configs.ConfigDB{
		DBHost:     "localhost",
		DBPort:     "5436",
		DBUser:     "postgres",
		DBPassword: "qwerty",
		DBName:     "postgres",
		DBSSLMode:  "disable",
	}

	application := app.New(log, dbConf, servConf)
	go application.GRPCSrv.MustRun()

	_ = application.GRPCSrv.Stop()
}
```

### When cloning the repository

```bash
go run cmd/main.go -serv=config/local.env -db=config/database.env
```

### Build via docker-compose

Navigate to the root of the repository and run the following command:

```bash
docker-compose up --build
```

## API

### Fetching all currency exchange rates

To get all currency exchange rates, send a gRPC request to `GetExchangeRates`.

**Request:**
```protobuf
message Empty {}
```
**Response:**
```protobuf
message ExchangeRatesResponse {
    map<string, OneCurrencyRate> rates = 1;
}

message OneCurrencyRate {
    map<string, float> rate = 1;
}
```

### Fetching exchange rates for specific currencies

To get the exchange rate for specific currencies, use the `GetExchangeRateForCurrency` method.

**Request:**
```protobuf
message CurrencyRequest {
    string fromCurrency = 1;
    string toCurrency = 2;
}
```
**Response:**
```protobuf
message ExchangeRateResponse {
    string fromCurrency = 1;
    string toCurrency = 2;
    float rate = 3;
}
```

## Testing

Integration tests are used to test the service, and can be run using:

```bash
go test ./...
```

## Documentation

API documentation is provided in Swagger format (swagger.json or swagger.yaml) and is available at the following path:

```bash
/docs/swagger.json
```

When running in the cloned repository, a `swagger` server is launched at `http://localhost:8181/swagger/index.html`

