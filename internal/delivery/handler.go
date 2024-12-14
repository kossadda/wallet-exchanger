package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/internal/service"
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

	auth := router.Group("/auth/v1")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.Login)
	}

	return router
}
