package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Env     string
	Expire  string
	Servers map[string]Server
}

type Server struct {
	ServerName string
	Port       string
	Host       string
}

func NewServerConfig(env, expire string, servers ...Server) *ServerConfig {
	serverMap := make(map[string]Server)
	for _, server := range servers {
		serverMap[server.ServerName] = server
	}

	return &ServerConfig{
		Env:     env,
		Expire:  expire,
		Servers: serverMap,
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

	return &ServerConfig{
		Env:     os.Getenv("APP_ENV"),
		Expire:  os.Getenv("EXPIRE"),
		Servers: serverMap,
	}
}
