package grpcclient

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
)

type handler struct {
	services *service.Service
	logger   *slog.Logger
}

func newHandler(services *service.Service, logger *slog.Logger) *handler {
	return &handler{
		services: services,
		logger:   logger,
	}
}

func (c *handler) GetExchangeRates(ctx *gin.Context) {
	const op = "GRPCClient.GetExchangeRates"

	log := c.logger.With(
		slog.String("operation", op),
	)

	log.Info("get currency rates")

	response, err := c.services.GetExchangeRates(ctx)
	if err != nil {
		util.NewErrorResponse(ctx, c.logger, http.StatusInternalServerError, "Failed to retrieve exchange rates")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *handler) ExchangeSum(ctx *gin.Context) {
	const op = "handler.ExchangeSum"

	userId, _ := ctx.Get("userId")
	input := &model.Exchange{
		UserId: userId.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		util.NewErrorResponse(ctx, c.logger, http.StatusBadRequest, err.Error())
		return
	}

	log := c.logger.With(
		slog.String("operation", op),
		slog.Any("user ID", userId),
		slog.String("input currency", input.FromCurrency),
		slog.String("output currency", input.ToCurrency),
		slog.Float64("input amount", input.Amount),
	)

	log.Info("exchange of one currency for another")

	response, err := c.services.ExchangeSum(ctx, input)
	if err != nil {
		util.NewErrorResponse(ctx, c.logger, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":          "Exchange successful",
		"exchanged_amount": input.Amount,
		"new_balance": map[string]float64{
			input.FromCurrency: response[0],
			input.ToCurrency:   response[1],
		},
	})
}

func (c *handler) CloseGRPC() error {
	return c.services.CloseGRPC()
}
