package grpcclient

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Exchange struct {
	conn   *grpc.ClientConn
	ex     gen.ExchangeServiceClient
	logger *slog.Logger
}

func New(address string, log *slog.Logger) *Exchange {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Warn("Failed to connect to gRPC server", "error", err)
		return nil
	}

	client := gen.NewExchangeServiceClient(conn)
	return &Exchange{
		conn:   conn,
		ex:     client,
		logger: log,
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

func (c *Exchange) Close() error {
	return c.conn.Close()
}
