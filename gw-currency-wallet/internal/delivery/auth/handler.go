package auth

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

// MainAPI is an interface that defines the methods required for user authentication.
type MainAPI interface {
	// Register to handle user registration
	Register(ctx *gin.Context)

	// Login to handle user login
	Login(ctx *gin.Context)
}

// Auth represents the authentication handler with methods to register and login users.
type Auth struct {
	MainAPI
}

// New creates and returns a new instance of the Auth handler with the provided services, logger, and config.
func New(services *service.Service, logger *slog.Logger, cfg *configs.ServerConfig) *Auth {
	return &Auth{
		MainAPI: newHandler(services, logger, cfg),
	}
}
