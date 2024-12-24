# GW Currency Wallet Service

Это микросервис для управления кошельком и обмена валют, который поддерживает регистрацию пользователей, пополнение
счета, вывод средств, получение курсов валют и обмен валют. Сервис использует gRPC для получения данных о курсах валют и
JWT для авторизации.

1. [GW Currency Wallet Service](#gw-currency-wallet-service)
2. [Стек технологий](#стек-технологий)
3. [Установка и запуск](#установка-и-запуск) \
   3.1. [Зависимости](#зависимости) \
   3.2. [При импортировании пакета](#при-импортировании-пакета) \
   3.3. [При клонировании репозитория](#при-клонировании-репозитория) \
   3.4. [Сборка через docker-compose](#сборка-через-docker-compose)
4. [API](#api) \
   4.1. [Регистрация пользователя](#1-регистрация-пользователя) \
   4.2. [Авторизация пользователя](#2-авторизация-пользователя) \
   4.3. [Получение баланса пользователя](#3-получение-баланса-пользователя) \
   4.4. [Пополнение счета](#4-пополнение-счета) \
   4.5. [Вывод средств](#5-вывод-средств) \
   4.6. [Получение курса валют](#6-получение-курса-валют) \
   4.7. [Обмен валют](#7-обмен-валют)
5. [Логирование](#логирование)
6. [Тестирование](#тестирование)
7. [Документация](#документация)

## Стек технологий

- **Go (Golang)** — язык программирования
- **Gin** — HTTP-фреймворк для разработки REST API
- **gRPC** — для общения с сервисом обмена валют
- **JWT** — для авторизации
- **PostgreSQL** — база данных для хранения данных о пользователях и их балансах
- **Docker** — для контейнеризации сервиса

## Установка и запуск

### Зависимости

Для работы операций по обмену валют необходимо развернуть сервера для `exchanger` и `redis`, а также указать их адреса в конфиге

Требуемые отношения базы данных `postgres` указаны в файлах директории `db/migrations` в корне репозитория

### При импортировании пакета

Укажите параметры запуска в `configs.ServerConfig`

**Переменные окружения**
- Для `Postgres`:
```env
DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE
```
- Для серверов:
```env
APP_ENV  // (local, dev, prod)
TOKEN_EXPIRE  // (в формате пакета time)
CACHE_EXPIRE, APP_PORT, APP_HOST, GRPC_PORT, GRPC_HOST, CACHE_PORT, CACHE_HOST
```

Пример запуска сервиса с указанием пути к `env` конфигурациям или созданием конфигов вручную:

```go
package main

import (
	"github.com/kossadda/wallet-exchanger/currency-wallet/pkg/app"
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

### При клонировании репозитория

```bash
go run cmd/main.go -serv=config/local.env -db=config/database.env
```

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

## Тестирование

Для тестирования сервиса используются интеграционные тесты, которые можно запустить с помощью:

```bash
go test ./...
```

## Документация

Документация API предоставляется в формате Swagger (swagger.json или swagger.yaml), и доступна по следующему пути:

```bash
/docs/swagger.json
```

При запуске в склонированном репозитории разворачивается `swagger`-сервер по адресу `http://localhost:8282/swagger/index.html`

