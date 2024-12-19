package delivery

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery/auth"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery/grpcclient"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery/middleware"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery/wallet"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

type Handler struct {
	*auth.Auth
	*wallet.Wallet
	*grpcclient.Exchange
	*middleware.Middleware
}

func New(services *service.Service, logger *slog.Logger, cfg *configs.ServerConfig) *Handler {
	return &Handler{
		Auth:       auth.New(services, logger, cfg),
		Wallet:     wallet.New(services, logger, cfg),
		Middleware: middleware.New(services, logger),
		Exchange:   grpcclient.New(services, logger),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	authorize := router.Group("/api/v1")
	{
		authorize.POST("/register", h.Register)
		authorize.POST("/login", h.Login)
	}

	api := router.Group("/api/v1", h.UserIdentity)
	{
		api.GET("/balance", h.GetBalance)
		api.POST("/wallet/deposit", h.Deposit)
		api.POST("/wallet/withdraw", h.Withdraw)
		api.POST("/exchange", h.ExchangeSum)
		api.GET("/exchange/rates", h.GetExchangeRates)
	}

	return router
}
