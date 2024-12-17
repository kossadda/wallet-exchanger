package auth

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

type MainAPI interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type Auth struct {
	MainAPI
}

func New(services *service.Service, logger *slog.Logger, cfg *configs.ServerConfig) *Auth {
	return &Auth{
		MainAPI: newHandler(services, logger, cfg),
	}
}
