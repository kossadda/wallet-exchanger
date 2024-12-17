package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

type Middleware struct {
	services *service.Service
	logger   *slog.Logger
}

func New(services *service.Service, logger *slog.Logger) *Middleware {
	return &Middleware{
		services: services,
		logger:   logger,
	}
}

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
