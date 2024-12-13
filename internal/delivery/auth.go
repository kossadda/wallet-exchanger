package delivery

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/model"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input model.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	val, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Created user %d", val))
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input model.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	val, err := h.services.Login(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Succes login for %d", val))
}
