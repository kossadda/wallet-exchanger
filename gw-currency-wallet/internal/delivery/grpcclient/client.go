// Package grpcclient facilitates communication with external gRPC services, specifically for exchange rate retrieval and currency exchange operations.
package grpcclient

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
)

// handler handles interactions with the gRPC service for currency exchange operations.
type handler struct {
	services *service.Service
	logger   *slog.Logger
}

// newHandler creates and returns a new handler for the gRPC client.
func newHandler(services *service.Service, logger *slog.Logger) *handler {
	return &handler{
		services: services,
		logger:   logger,
	}
}

// GetExchangeRates fetches exchange rates from an external service via gRPC.
// @Summary Get Exchange Rates
// @Description Fetches current exchange rates from an external service
// @Tags Exchange
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Exchange rates"
// @Failure 500 {string} string "Failed to retrieve exchange rates"
// @Router /api/v1/exchange/rates [get]
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

// ExchangeSum handles the exchange of one currency for another.
// @Summary Exchange Currency
// @Description Exchange one currency for another and update the user's balance
// @Tags Exchange
// @Accept json
// @Produce json
// @Param input body model.Exchange true "Currency exchange details"
// @Success 200 {object} map[string]interface{} "Exchange result"
// @Failure 400 {string} string "Exchange error"
// @Router /api/v1/exchange [post]
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

// CloseGRPC closes the gRPC client connection.
func (c *handler) CloseGRPC() error {
	return c.services.CloseGRPC()
}
