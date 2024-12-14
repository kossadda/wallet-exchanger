package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/model"
)

func (h *Handler) getBalance(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}

func (h *Handler) depositSum(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	input := &model.Deposit{
		UserId: userId.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.DepositSum(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Account topped up successfully",
		"new_balance": balance,
	})
}
