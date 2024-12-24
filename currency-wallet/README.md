# GW Currency Wallet Service

This is a microservice for wallet management and currency exchange, which supports user registration, account replenishment, withdrawals, getting exchange rates, and currency exchange. The service uses gRPC for retrieving exchange rates and JWT for authentication.

# Table of Contents

1. [GW Currency Wallet Service](#gw-currency-wallet-service)
2. [Technology Stack](#technology-stack)
3. [Installation and Setup](#installation-and-setup) \
   3.1. [Dependencies](#dependencies) \
   3.2. [When importing the package](#when-importing-the-package) \
   3.3. [When cloning the repository](#when-cloning-the-repository) \
   3.4. [Building with docker-compose](#building-with-docker-compose)
4. [API](#api) \
   4.1. [User Registration](#1-user-registration) \
   4.2. [User Authorization](#2-user-authorization) \
   4.3. [Get User Balance](#3-get-user-balance) \
   4.4. [Deposit Funds](#4-deposit-funds) \
   4.5. [Withdraw Funds](#5-withdraw-funds) \
   4.6. [Get Exchange Rates](#6-get-exchange-rates) \
   4.7. [Exchange Currency](#7-exchange-currency)
5. [Logging](#logging)
6. [Testing](#testing)
7. [Documentation](#documentation)

## Technology Stack

- **Go (Golang)** — programming language
- **Gin** — HTTP framework for developing REST API
- **gRPC** — for communication with the currency exchange service
- **JWT** — for authentication
- **PostgreSQL** — database for storing user data and balances
- **Docker** — for containerizing the service

## Installation and Setup

### Dependencies

For currency exchange operations to work, you need to deploy servers for `exchanger` and `redis`, and also specify their addresses in the config

The required `postgres` database relations are specified in the files of the `db/migrations` directory in the root of the repository

### When importing the package

Specify the startup parameters in `configs.ServerConfig`
> To perform currency exchange operations, you need to deploy servers for `exchanger` and `redis`, and provide their addresses in the config.

**Environment variables**
- For `Postgres`:
```env
DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE

```
- For servers:
```env
APP_ENV  // (local, dev, prod)
TOKEN_EXPIRE  // (in time package format)
CACHE_EXPIRE, APP_PORT, APP_HOST, GRPC_PORT, GRPC_HOST, CACHE_PORT, CACHE_HOST
```

Example of running the service with the path to `env` configurations or creating configs manually:

```go
package main

import (
	"github.com/kossadda/wallet-exchanger/currency-wallet/pkg/app"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
)

func main() {
    // You can specify the path to the configs
	//servConf := configs.NewServerEnvConfig("config/local.env")
	//dbConf := configs.NewEnvConfigDB("config/database.env")

    // Or define them manually
	servConf := &configs.ServerConfig{
		Env:         "local",
		TokenExpire: "10h",
		CacheExpire: "1m",
		Servers: map[string]configs.Server{
			"APP": configs.Server{
				Host: "localhost",
				Port: "8080",
			},
			"GRPC": configs.Server{
				Host: "localhost",
				Port: "44044",
			},
			"CACHE": configs.Server{
				Host: "localhost",
				Port: "6379",
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
	go application.Wallet.MustRun()

	_ = application.Wallet.Stop()
}
```

### When cloning the repository

```bash
go run cmd/main.go -serv=config/local.env -db=config/database.env
```

### Building with docker-compose

Navigate to the root of the repository and run the following command:

```bash
docker-compose up --build
```

## API

### 1. User Registration

- **Method:** `POST`
- **URL:** `/api/v1/register`
- **Request Body:**

    ```json
    {
      "username": "string",
      "password": "string",
      "email": "string"
    }
    ```

- **Response:**
    - Success (201 Created):

    ```json
    {
      "message": "User registered successfully"
    }
    ```

    - Error (400 Bad Request):

    ```json
    {
      "error": "Username or email already exists"
    }
    ```

### 2. User authorization

- **Method:** `POST`
- **URL:** `/api/v1/login`
- **Request Body:**

    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```

- **Response:**
    - Success (200 OK):

    ```json
    {
      "token": "JWT_TOKEN"
    }
    ```

    - Error (401 Unauthorized):

    ```json
    {
      "error": "Invalid username or password"
    }
    ```

### 3. Get User Balance

- **Method:** `GET`
- **URL:** `/api/v1/balance`
- **Headers:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Response:**
    - Success (200 OK):

    ```json
    {
      "balance": {
        "USD": "float",
        "RUB": "float",
        "EUR": "float"
      }
    }
    ```

### 4. Deposit Funds

- **Method:** `POST`
- **URL:** `/api/v1/wallet/deposit`
- **Headers:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Request Body:**

    ```json
    {
      "amount": 100.00,
      "currency": "USD"
    }
    ```

- **Response:**
    - Success (200 OK):

    ```json
    {
      "message": "Account topped up successfully",
      "new_balance": {
        "USD": "float",
        "RUB": "float",
        "EUR": "float"
      }
    }
    ```

    - Error (400 Bad Request):

    ```json
    {
      "error": "Invalid amount or currency"
    }
    ```

### 5. Withdraw Funds

- **Method:** `POST`
- **URL:** `/api/v1/wallet/withdraw`
- **Headers:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Request Body:**

    ```json
    {
      "amount": 50.00,
      "currency": "USD"
    }
    ```

- **Response:**
    - Success (200 OK):

    ```json
    {
      "message": "Withdrawal successful",
      "new_balance": {
        "USD": "float",
        "RUB": "float",
        "EUR": "float"
      }
    }
    ```

    - Error (400 Bad Request):

    ```json
    {
      "error": "Insufficient funds or invalid amount"
    }
    ```

### 6. Get Exchange Rates

- **Method:** `GET`
- **URL:** `/api/v1/exchange/rates`
- **Headers:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Response:**
    - Success (200 OK):

    ```json
    {
      "rates": {
        "USD": "float",
        "RUB": "float",
        "EUR": "float"
      }
    }
    ```

    - Error (500 Internal Server Error):

    ```json
    {
      "error": "Failed to retrieve exchange rates"
    }
    ```

### 7. Exchange Currency

- **Method:** `POST`
- **URL:** `/api/v1/exchange`
- **Headers:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Request Body:**

    ```json
    {
      "from_currency": "USD",
      "to_currency": "EUR",
      "amount": 100.00
    }
    ```

- **Response:**
    - Success (200 OK):

    ```json
    {
      "message": "Exchange successful",
      "exchanged_amount": 85.00,
      "new_balance": {
        "USD": 0.00,
        "EUR": 85.00
      }
    }
    ```

    - Error (400 Bad Request):

    ```json
    {
      "error": "Insufficient funds or invalid currencies"
    }
    ```

## Logging

The service uses logging to track errors and important operations. Logs can be viewed through standard console output or files, depending on the configuration.

## Testing

For testing the service, integration tests are used, and you can run them with:

```bash
go test ./...
```

## Documentation

The API documentation is provided in Swagger format (swagger.json or swagger.yaml) and is available at the following path:

```bash
/docs/swagger.json
```

When running from the cloned repository, a `swagger` server will be deployed at `http://localhost:8282/swagger/index.html`

