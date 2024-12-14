package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getBalance(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	balance, err := h.services.GetBalance(userId.(int))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, balance)
}
