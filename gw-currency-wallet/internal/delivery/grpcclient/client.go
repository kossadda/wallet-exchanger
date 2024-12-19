package grpcclient

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Exchange struct {
	conn     *grpc.ClientConn
	ex       gen.ExchangeServiceClient
	services *service.Service
	logger   *slog.Logger
}

func New(address string, services *service.Service, log *slog.Logger) *Exchange {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Warn("Failed to connect to gRPC server", "error", err)
		return nil
	}

	client := gen.NewExchangeServiceClient(conn)
	return &Exchange{
		conn:     conn,
		ex:       client,
		services: services,
		logger:   log,
	}
}

func (c *Exchange) GetExchangeRates(ctx *gin.Context) {
	const op = "GRPCClient.GetExchangeRates"

	log := c.logger.With(
		slog.String("operation", op),
	)

	log.Info("get currency rates")

	response, err := c.ex.GetExchangeRates(ctx, &gen.Empty{})
	if err != nil {
		util.NewErrorResponse(ctx, c.logger, http.StatusInternalServerError, "Failed to retrieve exchange rates")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Exchange) ExchangeSum(ctx *gin.Context) {
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

	response, err := c.ex.GetExchangeRateForCurrency(ctx, &gen.CurrencyRequest{
		FromCurrency: input.FromCurrency,
		ToCurrency:   input.ToCurrency,
	})
	if err != nil {
		util.NewErrorResponse(ctx, c.logger, http.StatusBadRequest, err.Error())
		return
	}

	updateFrom, err := c.services.Withdraw(&model.Operation{
		UserId:   userId.(int),
		Currency: input.FromCurrency,
		Amount:   input.Amount,
	})
	if err != nil {
		util.NewErrorResponse(ctx, c.logger, http.StatusBadRequest, err.Error())
		return
	}

	updateTo, err := c.services.Deposit(&model.Operation{
		UserId:   userId.(int),
		Currency: input.ToCurrency,
		Amount:   input.Amount * float64(response.Rate),
	})
	if err != nil {
		util.NewErrorResponse(ctx, c.logger, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":          "Exchange successful",
		"exchanged_amount": input.Amount,
		"new_balance": map[string]float64{
			input.FromCurrency: updateFrom,
			input.ToCurrency:   updateTo,
		},
	})
}

func (c *Exchange) Close() error {
	return c.conn.Close()
}
