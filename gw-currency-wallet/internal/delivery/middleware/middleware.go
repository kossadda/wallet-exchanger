// Package middleware provides middleware for handling common tasks like user identity verification (token validation).
package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
)

// Constants used for authorization and user context handling.
const (
	authorizationHeader = "Authorization" // Authorization header key
	userCtx             = "userId"        // Context key for storing user ID
)

// Middleware handles HTTP request interception and user authorization.
type Middleware struct {
	services *service.Service
	logger   *slog.Logger
}

// New creates and returns a new instance of Middleware with the provided services and logger.
func New(services *service.Service, logger *slog.Logger) *Middleware {
	return &Middleware{
		services: services,
		logger:   logger,
	}
}

// UserIdentity is a middleware function that checks the authorization header and parses the user ID.
// @Summary Verifies the user identity from the authorization token.
// @Description This middleware validates the presence and correctness of the authorization token in the HTTP request header.
// @Tags Middleware
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Failure 401
// @Router /api/v1 [get]
func (h *Middleware) UserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		util.NewErrorResponse(ctx, h.logger, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		util.NewErrorResponse(ctx, h.logger, http.StatusUnauthorized, "invalid authorization header")
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusUnauthorized, err.Error())
	}

	ctx.Set(userCtx, userId)
}
