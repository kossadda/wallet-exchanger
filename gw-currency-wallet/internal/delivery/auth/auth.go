package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

type Handler struct {
	services *service.Service
	logger   *slog.Logger
	config   *configs.ServerConfig
}

func newHandler(services *service.Service, logger *slog.Logger, config *configs.ServerConfig) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
		config:   config,
	}
}

func (h *Handler) Register(ctx *gin.Context) {
	var input model.User

	if err := ctx.BindJSON(&input); err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.CreateUser(input)
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, "Username or email already exists")
		return
	}

	ctx.JSON(http.StatusCreated, "User registered successfully")
}

func (h *Handler) Login(ctx *gin.Context) {
	var input model.LogUser

	if err := ctx.BindJSON(&input); err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := h.services.GenerateToken(input.Username, input.Password, h.config.TokenTTL)
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	ctx.JSON(http.StatusOK, token)
}
