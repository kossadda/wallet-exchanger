package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Env         string
	StoragePath string
	TokenTTL    string
	Port        string
}

func NewServerConfig(env, storagePath, tokenTTL, port string) *ServerConfig {
	return &ServerConfig{
		Env:         env,
		StoragePath: storagePath,
		TokenTTL:    tokenTTL,
		Port:        port,
	}
}

func NewServerEnvConfig(configPath string) *ServerConfig {
	err := godotenv.Load(configPath)
	if err != nil {
		panic(err)
	}

	return &ServerConfig{
		Env:         os.Getenv("APP_ENV"),
		StoragePath: os.Getenv("APP_STORAGE_PATH"),
		TokenTTL:    os.Getenv("APP_TOKEN_TTL"),
		Port:        os.Getenv("APP_PORT"),
	}
}
