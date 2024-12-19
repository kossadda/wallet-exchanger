package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Env         string
	TokenExpire string
	CacheExpire string
	Servers     map[string]Server
}

type Server struct {
	ServerName string
	Port       string
	Host       string
}

func NewServerConfig(env, tokenExp, cacheExp string, servers ...Server) *ServerConfig {
	serverMap := make(map[string]Server)
	for _, server := range servers {
		serverMap[server.ServerName] = server
	}

	return &ServerConfig{
		Env:         env,
		TokenExpire: tokenExp,
		CacheExpire: cacheExp,
		Servers:     serverMap,
	}
}

func NewServerEnvConfig(configPath string) *ServerConfig {
	err := godotenv.Load(configPath)
	if err != nil {
		panic(err)
	}

	serverMap := make(map[string]Server)

	serverPrefixes := []string{"APP", "GRPC", "CACHE"}

	for _, prefix := range serverPrefixes {
		host := os.Getenv(prefix + "_HOST")
		port := os.Getenv(prefix + "_PORT")

		if host != "" && port != "" {
			serverMap[prefix] = Server{
				ServerName: prefix,
				Port:       port,
				Host:       host,
			}
		}
	}

	env := os.Getenv("APP_ENV")
	if env != "" {
		env = DefaultENV
	}

	return &ServerConfig{
		Env:         env,
		TokenExpire: os.Getenv("TOKEN_EXPIRE"),
		CacheExpire: os.Getenv("CACHE_EXPIRE"),
		Servers:     serverMap,
	}
}
