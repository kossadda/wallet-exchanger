// Package auth contains the logic for handling user authentication-related operations,
// including user registration, login, and token management.
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

// handler is a struct that handles the user authentication operations.
// It contains references to the services layer, logger, and configuration.
type handler struct {
	services *service.Service
	logger   *slog.Logger
	config   *configs.ServerConfig
}

// newHandler creates and returns a new handler instance.
// It is responsible for setting up the dependencies for handling user authentication.
func newHandler(services *service.Service, logger *slog.Logger, config *configs.ServerConfig) *handler {
	return &handler{
		services: services,
		logger:   logger,
		config:   config,
	}
}

// Register handles the user registration process.
// @Summary User Registration
// @Description Registers a new user in the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body model.User true "User Registration"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {string} string "Username or email already exists"
// @Router /api/v1/register [post]
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

// Login handles the user login process, validates the input,
// and generates a JWT token for the authenticated user.
// @Summary User Login
// @Description Logs in the user and generates a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body model.LogUser true "Login Credentials"
// @Success 200 {string} string "JWT Token"
// @Failure 401 {string} string "Invalid username or password"
// @Router /api/v1/login [post]
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

	token, err := h.services.GenerateToken(input.Username, input.Password, h.config.TokenExpire)
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	ctx.JSON(http.StatusOK, token)
}
