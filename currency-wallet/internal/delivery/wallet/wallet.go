// Package wallet manages user wallet operations, including balance retrieval, deposits, and withdrawals.
package wallet

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/util"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

// handler is a structure that holds the services, logger, and config
// required for handling wallet-related HTTP requests.
type handler struct {
	services *service.Service
	logger   *slog.Logger
	config   *configs.ServerConfig
}

// newHandler creates and returns a new instance of the handler with the provided services, logger, and config.
func newHandler(services *service.Service, logger *slog.Logger, config *configs.ServerConfig) *handler {
	return &handler{
		services: services,
		logger:   logger,
		config:   config,
	}
}

// GetBalance to get user balance
// @Summary Fetches the balance of the authenticated user.
// @Description This endpoint returns the current balance of the user’s wallet.
// @Tags Wallet
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{} "Current user balance"
// @Failure 400
// @Failure 401
// @Router /api/v1/balance [get]
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

// Deposit to deposit funds into wallet
// @Summary Deposits an amount into the authenticated user's wallet.
// @Description This endpoint adds a specified amount of currency to the user's wallet.
// @Tags Wallet
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param amount body model.Operation true "Amount to deposit"
// @Success 200 {object} map[string]interface{} "New wallet balance after deposit"
// @Failure 400
// @Failure 401
// @Router /api/v1/wallet/deposit [post]
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

	_, err := h.services.Deposit(input)
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

// Withdraw to withdraw funds from wallet
// @Summary Withdraws an amount from the authenticated user's wallet.
// @Description This endpoint deducts a specified amount from the user's wallet.
// @Tags Wallet
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param amount body model.Operation true "Amount to withdraw"
// @Success 200 {object} map[string]interface{} "New wallet balance after withdrawal"
// @Failure 400
// @Failure 401
// @Router /api/v1/wallet/withdraw [post]
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

	_, err := h.services.Withdraw(input)
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
