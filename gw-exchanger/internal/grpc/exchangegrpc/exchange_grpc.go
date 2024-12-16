package exchangegrpc

import (
	"context"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"google.golang.org/grpc"
)

const (
	usdRateRub = 103.85
	usdRateEur = 0.95
	rubRateUsd = 0.0096
	rubRateEur = 0.0092
	eurRateRub = 1.05
	eurRateUsd = 109.08
)

type serverAPI struct {
	gen.UnimplementedExchangerServer
}

func Register(gRPC *grpc.Server) {
	gen.RegisterExchangerServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Exchange(ctx context.Context, req *gen.ExchangeRequest) (*gen.ExchangeResponse, error) {
	//newSum := req.Sum
	//switch req.InputCurrency {
	//case "USD":
	//	switch req.OutputCurrency {
	//	case "RUB":
	//		newSum *= usdRateRub
	//	case "EUR":
	//
	//	}
	//case "EUR":
	//case "RUB":
	//}
	return &gen.ExchangeResponse{
		Sum: 1234,
	}, nil
}
