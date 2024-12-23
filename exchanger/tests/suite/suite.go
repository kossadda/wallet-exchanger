package suite

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/exchanger/internal/app"
	"github.com/kossadda/wallet-exchanger/exchanger/internal/service"
	"github.com/kossadda/wallet-exchanger/exchanger/internal/storage"
	"github.com/kossadda/wallet-exchanger/exchanger/internal/storage/exchange"
	"github.com/kossadda/wallet-exchanger/exchanger/internal/storage/exchange/mockex"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	GrpcTestPort = "44042"
)

type Suite struct {
	Ctrl   *gomock.Controller
	Client gen.ExchangeServiceClient
	App    *app.GRPCApp
}

type FakeDB struct{}

func (f *FakeDB) Transaction(fn func(tx *sqlx.Tx) error) error {
	return nil
}

func (f *FakeDB) Close() error {
	return nil
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	t.Cleanup(func() {
		t.Helper()
		cancel()
	})

	cfg := &configs.ServerConfig{
		Env: "local",
		Servers: map[string]configs.Server{
			"GRPCAPP": configs.Server{
				Host: "localhost",
				Port: GrpcTestPort,
			},
			"CLIENT": configs.Server{
				Host: "localhost",
				Port: GrpcTestPort,
			},
		},
	}

	ctrl, services := fakeService(t, ctx)

	return ctx, &Suite{
		Ctrl:   ctrl,
		Client: grpcClient(cfg),
		App:    grpcApp(services, cfg),
	}
}

func fakeService(t *testing.T, ctx context.Context) (*gomock.Controller, *service.Service) {
	ctrl := gomock.NewController(t)

	mockRepo := mockex.NewMockMainAPI(gomock.NewController(t))

	response := &gen.ExchangeRatesResponse{
		Rates: map[string]*gen.OneCurrencyRate{
			"usd": &gen.OneCurrencyRate{
				Rate: map[string]float32{
					"usd": 1.0,
					"rub": 0.0097,
					"eur": 1.05,
				},
			},
			"rub": &gen.OneCurrencyRate{
				Rate: map[string]float32{
					"usd": 103.6,
					"rub": 1.0,
					"eur": 108.89,
				},
			},
			"eur": &gen.OneCurrencyRate{
				Rate: map[string]float32{
					"usd": 0.95,
					"rub": 0.0092,
					"eur": 1.0,
				},
			},
		},
	}

	gomock.InOrder(
		mockRepo.EXPECT().GetExchangeRates(gomock.Any()).Return(response, nil).AnyTimes(),
	)

	return ctrl, service.New(&storage.Storage{
		DataBase: &FakeDB{},
		Exchange: &exchange.Exchange{
			MainAPI: mockRepo,
		},
	})
}

func grpcApp(services *service.Service, cfg *configs.ServerConfig) *app.GRPCApp {
	gRPCServer := grpc.NewServer()
	log := logger.SetupByEnv(cfg.Env)

	appAddr, ok := cfg.Servers["GRPCAPP"]
	if !ok {
		panic("Wrong grpc App Server")
	}

	return app.New(gRPCServer, log, services, appAddr.Port)
}

func grpcClient(cfg *configs.ServerConfig) gen.ExchangeServiceClient {
	grpcAddr, _ := cfg.Servers["CLIENT"]

	conn, err := grpc.NewClient(grpcAddr.Host+":"+grpcAddr.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to gRPC server: %v", err))
	}

	return gen.NewExchangeServiceClient(conn)
}
