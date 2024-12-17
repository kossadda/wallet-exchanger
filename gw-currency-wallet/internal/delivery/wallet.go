package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
)

func (h *Handler) getBalance(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		newErrorResponse(ctx, h.logger, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}

func (h *Handler) deposit(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	input := &model.Operation{
		UserId: userId.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		newErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Deposit(input)
	if err != nil {
		newErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		newErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Account topped up successfully",
		"new_balance": balance,
	})
}

func (h *Handler) withdraw(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	input := &model.Operation{
		UserId: userId.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		newErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Withdraw(input)
	if err != nil {
		newErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		newErrorResponse(ctx, h.logger, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Withdrawal successful",
		"new_balance": balance,
	})
}
