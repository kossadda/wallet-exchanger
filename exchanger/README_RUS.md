# Wallet Exchanger gRPC Service

Этот сервис предоставляет API для получения актуальных обменных курсов валют. Он работает через gRPC и может быть интегрирован с другими сервисами для обмена валют.

Сервис предоставляет два основных метода:

1. **GetExchangeRates** — возвращает текущие курсы валют.
2. **GetExchangeRateForCurrency** — возвращает курс обмена между двумя конкретными валютами.

# Содержание

1. [Wallet Exchanger gRPC Service](#wallet-exchanger-grpc-service)
2. [Установка и запуск](#установка-и-запуск) \
   2.1. [Зависимости](#зависимости) \
   2.2. [При импортировании пакета](#при-импортировании-пакета) \
   2.3. [При клонировании репозитория](#при-клонировании-репозитория) \
   2.4. [Сборка через docker-compose](#сборка-через-docker-compose)
3. [API](#api) \
   3.1. [Получение всех курсов валют](#получение-всех-курсов-валют) \
   3.2. [Получение курса для конкретных валют](#получение-курса-для-конкретных-валют)
4. [Тестирование](#тестирование)
5. [Документация](#документация)

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
APP_PORT, APP_HOST, GRPC_PORT, GRPC_HOST, CACHE_PORT, CACHE_HOST
```

Пример запуска сервиса с указанием пути к `env` конфигурациям или созданием конфигов вручную:

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

### Получение всех курсов валют

Чтобы получить все курсы валют, отправьте gRPC запрос `GetExchangeRates`.

**Запрос:**
```protobuf
message Empty {}
```
**Ответ:**
```protobuf
message ExchangeRatesResponse {
    map<string, OneCurrencyRate> rates = 1;
}

message OneCurrencyRate {
    map<string, float> rate = 1;
}
```

### Получение курса для конкретных валют

Чтобы получить курс обмена для конкретных валют, используйте метод `GetExchangeRateForCurrency`.

**Запрос:**
```protobuf
message CurrencyRequest {
    string fromCurrency = 1;
    string toCurrency = 2;
}
```
**Ответ:**
```protobuf
message ExchangeRateResponse {
    string fromCurrency = 1;
    string toCurrency = 2;
    float rate = 3;
}
```

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

При запуске в склонированном репозитории разворачивается `swagger`-сервер по адресу `http://localhost:8181/swagger/index.html`

