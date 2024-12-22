// Package grpcclient provides gRPC client functionality for exchanging rates and performing wallet operations.
package grpcclient

import (
	"context"
	"fmt"
	"strings"

	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service/wallet"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/currency-wallet/pkg/cache"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ratesCacheKey = "rates-cache" // The cache key used to store and retrieve exchange rates.
)

// service represents the gRPC client that interacts with the exchange service, handles wallet operations, and caches results.
type service struct {
	wallet *wallet.Wallet
	conn   *grpc.ClientConn
	rpc    gen.ExchangeServiceClient
	ch     cache.Cache
}

// newService creates and returns a new instance of service with the provided repository and configuration.
func newService(repo *storage.Storage, cfg *configs.ServerConfig) *service {
	grpcAddr, ok := cfg.Servers["GRPC"]
	if !ok {
		panic(fmt.Errorf("gRPC address not found in config"))
	}

	conn, err := grpc.NewClient(grpcAddr.Host+":"+grpcAddr.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to gRPC server: %v", err))
	}
	client := gen.NewExchangeServiceClient(conn)

	ch, err := cache.NewRedis(context.Background(), cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to redis server: %v", err))
	}

	return &service{
		wallet: wallet.New(repo),
		conn:   conn,
		rpc:    client,
		ch:     ch,
	}
}

// GetExchangeRates retrieves exchange rates from the cache or fetches them from the gRPC server if not cached.
func (s *service) GetExchangeRates(ctx context.Context) (*gen.ExchangeRatesResponse, error) {
	var response *gen.ExchangeRatesResponse
	if err := s.ch.Get(ctx, ratesCacheKey, &response); err != nil {
		response, err = s.rpc.GetExchangeRates(ctx, &gen.Empty{})
		if err != nil {
			return nil, err
		}
		_ = s.ch.Set(ctx, ratesCacheKey, response)
	}

	return response, nil
}

// ExchangeSum performs the currency exchange by withdrawing from the user's wallet in one currency and depositing in another.
func (s *service) ExchangeSum(ctx context.Context, input *model.Exchange) ([]float64, error) {
	r, err := s.GetExchangeRates(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get exchange rates from gRPC server")
	}

	toRates, ok := r.Rates[strings.ToLower(input.ToCurrency)]
	if !ok {
		return nil, fmt.Errorf("invalid output currency %s", input.ToCurrency)
	}
	resRate, ok := toRates.Rate[strings.ToLower(input.FromCurrency)]
	if !ok {
		return nil, fmt.Errorf("invalid input currency %s", input.FromCurrency)
	}

	updateFrom, err := s.wallet.Withdraw(&model.Operation{
		UserId:   input.UserId,
		Currency: input.FromCurrency,
		Amount:   input.Amount,
	})
	if err != nil {
		return nil, err
	}

	updateTo, err := s.wallet.Deposit(&model.Operation{
		UserId:   input.UserId,
		Currency: input.ToCurrency,
		Amount:   input.Amount * float64(resRate),
	})
	if err != nil {
		return nil, err
	}

	return []float64{updateFrom, updateTo}, nil
}

// CloseGRPC closes the gRPC connection.
func (s *service) CloseGRPC() error {
	return s.conn.Close()
}
