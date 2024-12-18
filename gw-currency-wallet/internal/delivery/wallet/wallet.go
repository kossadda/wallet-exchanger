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

func (h *handler) GetBalance(ctx *gin.Context) {
	const op = "handler.GetBalance"

	userId, _ := ctx.Get("userId")

	log := h.logger.With(
		slog.String("operation", op),
		slog.Any("user ID", userId),
	)

	log.Info("get user balance")

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}

func (h *handler) Deposit(ctx *gin.Context) {
	const op = "handler.Deposit"

	userId, _ := ctx.Get("userId")
	input := &model.Operation{
		UserId: userId.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	log := h.logger.With(
		slog.String("operation", op),
		slog.Any("user ID", userId),
		slog.String("input currency", input.Currency),
		slog.Float64("input amount", input.Amount),
	)

	log.Info("deposit sum on account")

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

func (h *handler) Withdraw(ctx *gin.Context) {
	const op = "handler.Withdraw"

	userId, _ := ctx.Get("userId")
	input := &model.Operation{
		UserId: userId.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		util.NewErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	log := h.logger.With(
		slog.String("operation", op),
		slog.Any("user ID", userId),
		slog.String("input currency", input.Currency),
		slog.Float64("input amount", input.Amount),
	)

	log.Info("withdraw sum on account")

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
