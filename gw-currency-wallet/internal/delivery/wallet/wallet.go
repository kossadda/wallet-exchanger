package wallet

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

func (h *Handler) GetBalance(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}

func (h *Handler) Deposit(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	input := &model.Operation{
		UserId: userId.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Deposit(input)
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Account topped up successfully",
		"new_balance": balance,
	})
}

func (h *Handler) Withdraw(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	input := &model.Operation{
		UserId: userId.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Withdraw(input)
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Withdrawal successful",
		"new_balance": balance,
	})
}
