# Wallet Service

Этот репозиторий предоставляет сервис для управления кошельками пользователей и обмена валют через gRPC. Сервис включает функциональность регистрации, авторизации, пополнения счета, вывода средств и обмена валют. API работает с использованием gRPC для обмена данными и позволяет интегрировать обмен валют с внешними сервисами.

## Содержание

1. [GW Currency Wallet Service](#wallet-service)
2. [Установка и запуск](#установка-и-запуск) \
   2.1. [При клонировании репозитория](#при-клонировании-репозитория) \
   2.2. [Сборка через docker-compose](#сборка-через-docker-compose)
3. [API](#api) \
   3.1. [Регистрация пользователя](#1-регистрация-пользователя) \
   3.2. [Авторизация пользователя](#2-авторизация-пользователя) \
   3.3. [Получение баланса пользователя](#3-получение-баланса-пользователя) \
   3.4. [Пополнение счета](#4-пополнение-счета) \
   3.5. [Вывод средств](#5-вывод-средств) \
   3.6. [Получение курса валют](#6-получение-курса-валют) \
   3.7. [Обмен валют](#7-обмен-валют)
4. [Логирование](#логирование)
5. [Документация](#документация)

## Установка и запуск

### При клонировании репозитория

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/your-repo/your-project.git
   cd your-project
    ```

2. Для локальной разработки каждого сервиса создайте файлы конфигурации, по аналогии с `config/local.env` для параметров сервера и `config/database.env` для настроек базы данных в каждом сервере, или используйте готовые примеры конфигураций.

### Сборка через docker-compose

Перейдите в корень репозитория и пропишите следующую команду:

```bash
docker-compose up --build
```

## API

### 1. Регистрация пользователя

- **Метод:** `POST`
- **URL:** `/api/v1/register`
- **Тело запроса:**

    ```json
    {
      "username": "string",
      "password": "string",
      "email": "string"
    }
    ```

- **Ответ:**
    - Успех (201 Created):

    ```json
    {
      "message": "User registered successfully"
    }
    ```

    - Ошибка (400 Bad Request):

    ```json
    {
      "error": "Username or email already exists"
    }
    ```

### 2. Авторизация пользователя

- **Метод:** `POST`
- **URL:** `/api/v1/login`
- **Тело запроса:**

    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```

- **Ответ:**
    - Успех (200 OK):

    ```json
    {
      "token": "JWT_TOKEN"
    }
    ```

    - Ошибка (401 Unauthorized):

    ```json
    {
      "error": "Invalid username or password"
    }
    ```

### 3. Получение баланса пользователя

- **Метод:** `GET`
- **URL:** `/api/v1/balance`
- **Заголовки:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Ответ:**
    - Успех (200 OK):

    ```json
    {
      "balance": {
        "USD": "float",
        "RUB": "float",
        "EUR": "float"
      }
    }
    ```

### 4. Пополнение счета

- **Метод:** `POST`
- **URL:** `/api/v1/wallet/deposit`
- **Заголовки:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Тело запроса:**

    ```json
    {
      "amount": 100.00,
      "currency": "USD"
    }
    ```

- **Ответ:**
    - Успех (200 OK):

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

    - Ошибка (400 Bad Request):

    ```json
    {
      "error": "Invalid amount or currency"
    }
    ```

### 5. Вывод средств

- **Метод:** `POST`
- **URL:** `/api/v1/wallet/withdraw`
- **Заголовки:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Тело запроса:**

    ```json
    {
      "amount": 50.00,
      "currency": "USD"
    }
    ```

- **Ответ:**
    - Успех (200 OK):

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

    - Ошибка (400 Bad Request):

    ```json
    {
      "error": "Insufficient funds or invalid amount"
    }
    ```

### 6. Получение курса валют

- **Метод:** `GET`
- **URL:** `/api/v1/exchange/rates`
- **Заголовки:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Ответ:**
    - Успех (200 OK):

    ```json
    {
      "rates": {
        "USD": "float",
        "RUB": "float",
        "EUR": "float"
      }
    }
    ```

    - Ошибка (500 Internal Server Error):

    ```json
    {
      "error": "Failed to retrieve exchange rates"
    }
    ```

### 7. Обмен валют

- **Метод:** `POST`
- **URL:** `/api/v1/exchange`
- **Заголовки:**

    ```text
    Authorization: Bearer JWT_TOKEN
    ```

- **Тело запроса:**

    ```json
    {
      "from_currency": "USD",
      "to_currency": "EUR",
      "amount": 100.00
    }
    ```

- **Ответ:**
    - Успех (200 OK):

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

    - Ошибка (400 Bad Request):

    ```json
    {
      "error": "Insufficient funds or invalid currencies"
    }
    ```

## Логирование

Сервис использует логирование для отслеживания ошибок и важнейших операций. Логи можно просматривать через стандартные
механизмы для вывода в консоль или в файлы, в зависимости от конфигурации.

## Документация

При запуске разворачиваются `swagger`-сервера по адресу:
- `http://localhost:8181/swagger/index.html` - для `exchanger` сервера
- `http://localhost:8282/swagger/index.html` - для `currency-wallet` сервера

