# Wallet Service

This repository provides a service for managing user wallets and currency exchange via gRPC. The service includes functionality for user registration, authentication, account deposit, withdrawal, and currency exchange. The API works using gRPC to exchange data and allows integration of currency exchange with external services.

## Contents

1. [GW Currency Wallet Service](#wallet-service)
2. [Installation and Setup](#installation-and-setup) \
   2.1. [When cloning the repository](#when-cloning-the-repository) \
   2.2. [Build with docker-compose](#build-with-docker-compose)
3. [API](#api) \
   3.1. [User Registration](#1-user-registration) \
   3.2. [User Authorization](#2-user-authorization) \
   3.3. [Get User Balance](#3-get-user-balance) \
   3.4. [Deposit Funds](#4-deposit-funds) \
   3.5. [Withdraw Funds](#5-withdraw-funds) \
   3.6. [Get Exchange Rates](#6-get-exchange-rates) \
   3.7. [Exchange Currency](#7-exchange-currency)
4. [Logging](#logging)
5. [Documentation](#documentation)

## Installation and Setup

### When cloning the repository

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/your-project.git
   cd your-project
    ```

2. For local development of each service, create configuration files similar to `config/local.env` for server parameters and `config/database.env` for database settings in each server, or use the provided example configurations.

### Build with docker-compose

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

The service uses logging to track errors and important operations. Logs can be viewed through standard mechanisms for console or file output, depending on the configuration.

## Documentation

When running, the `swagger` servers are deployed at the following addresses:
- `http://localhost:8181/swagger/index.html` - for the `exchanger` server
- `http://localhost:8282/swagger/index.html` - for the `currency-wallet` server

