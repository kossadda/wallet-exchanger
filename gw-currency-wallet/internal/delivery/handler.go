package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/api/v1")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.Login)
	}

	api := router.Group("/api/v1", h.userIdentity)
	{
		api.GET("/balance", h.getBalance)
		api.POST("/wallet/deposit", h.deposit)
		api.POST("/wallet/withdraw", h.withdraw)
	}

	return router
}
