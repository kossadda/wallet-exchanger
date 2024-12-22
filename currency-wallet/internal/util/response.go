package util

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, logger *slog.Logger, statusCode int, message string) {
	logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, response{Message: message})
}
