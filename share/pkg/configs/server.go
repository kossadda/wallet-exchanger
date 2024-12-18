package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Env      string
	TokenTTL string
	Port     string
	GrpcPort string
}

func NewServerConfig(env, tokenTTL, port, grpcPort string) *ServerConfig {
	return &ServerConfig{
		Env:      env,
		TokenTTL: tokenTTL,
		Port:     port,
		GrpcPort: grpcPort,
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
		GrpcPort: os.Getenv("GRPC_PORT"),
	}
}
