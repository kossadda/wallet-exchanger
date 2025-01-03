package configs

import (
	"os"

	"github.com/joho/godotenv"
)

// ConfigDB holds the configuration settings for the PostgreSQL database connection.
type ConfigDB struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// NewConfigDB creates a new ConfigDB instance with the provided values.
func NewConfigDB(dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode, serverPort string) *ConfigDB {
	return &ConfigDB{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBSSLMode:  dbSSLMode,
	}
}

// NewEnvConfigDB loads configuration from a .env file and returns a ConfigDB instance.
func NewEnvConfigDB(configPath string) *ConfigDB {
	err := godotenv.Load(configPath)
	if err != nil {
		panic(err)
	}

	return &ConfigDB{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}
}
