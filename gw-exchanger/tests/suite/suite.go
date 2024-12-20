package suite

import (
	"context"
	"fmt"
	"testing"
	"time"

	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T
	Cfg    *configs.ServerConfig
	Client gen.ExchangeServiceClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	t.Cleanup(func() {
		t.Helper()
		cancel()
	})

	cfg := &configs.ServerConfig{
		Env:         "local",
		TokenExpire: "10h",
		CacheExpire: "1m",
		Servers: map[string]configs.Server{
			"GRPC": configs.Server{
				Host: "localhost",
				Port: "44044",
			},
		},
	}

	grpcAddr, ok := cfg.Servers["GRPC"]
	if !ok {
		panic(fmt.Errorf("gRPC address not found in config"))
	}

	conn, err := grpc.NewClient(grpcAddr.Host+":"+grpcAddr.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to gRPC server: %v", err))
	}

	return ctx, &Suite{
		T:      t,
		Cfg:    cfg,
		Client: gen.NewExchangeServiceClient(conn),
	}
}
