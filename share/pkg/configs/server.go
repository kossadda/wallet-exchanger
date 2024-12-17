package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Env      string
	TokenTTL string
	Port     string
}

func NewServerConfig(env, tokenTTL, port string) *ServerConfig {
	return &ServerConfig{
		Env:      env,
		TokenTTL: tokenTTL,
		Port:     port,
	}
}

func NewServerEnvConfig(configPath string) *ServerConfig {
	err := godotenv.Load(configPath)
	if err != nil {
		panic(err)
	}

	return &ServerConfig{
		Env:      os.Getenv("APP_ENV"),
		TokenTTL: os.Getenv("APP_TOKEN_TTL"),
		Port:     os.Getenv("APP_PORT"),
	}
}
