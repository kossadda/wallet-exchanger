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

type handler struct {
	services *service.Service
	logger   *slog.Logger
	config   *configs.ServerConfig
}

func newHandler(services *service.Service, logger *slog.Logger, config *configs.ServerConfig) *handler {
	return &handler{
		services: services,
		logger:   logger,
		config:   config,
	}
}

func (h *handler) Register(ctx *gin.Context) {
	const op = "handler.Register"

	var input model.User

	if err := ctx.BindJSON(&input); err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	log := h.logger.With(
		slog.String("operation", op),
		slog.String("input login", input.Username),
		slog.String("input password", input.Password),
		slog.String("input email", input.Email),
	)

	log.Info("register in service")

	err := h.services.CreateUser(input)
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, "Username or email already exists")
		return
	}

	ctx.JSON(http.StatusCreated, "User registered successfully")
}

func (h *handler) Login(ctx *gin.Context) {
	const op = "handler.Login"

	var input model.LogUser

	if err := ctx.BindJSON(&input); err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusUnauthorized, err.Error())
		return
	}

	log := h.logger.With(
		slog.String("operation", op),
		slog.String("input login", input.Username),
		slog.String("input password", input.Password),
	)

	log.Info("authorization to get bearer token")

	token, err := h.services.GenerateToken(input.Username, input.Password, h.config.Expire)
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	ctx.JSON(http.StatusOK, token)
}
