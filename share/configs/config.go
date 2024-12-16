package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

func New(dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode, serverPort string) *Config {
	return &Config{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBSSLMode:  dbSSLMode,
		ServerPort: serverPort,
	}
}

func NewEnvConfig(configPath string) *Config {
	err := godotenv.Load(configPath)
	if err != nil {
		dflt := NewDefaultConfig()
		logrus.Error(err, "Use default config", dflt)
		return dflt
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
