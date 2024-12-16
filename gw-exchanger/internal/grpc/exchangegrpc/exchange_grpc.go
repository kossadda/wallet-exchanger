package exchangegrpc

import (
	"context"

	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/service"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	gen.UnimplementedExchangerServer
	service *service.Service
}

func Register(gRPC *grpc.Server, service *service.Service) {
	gen.RegisterExchangerServer(gRPC, &serverAPI{
		service: service,
	})
}

func (s *serverAPI) Exchange(ctx context.Context, req *gen.ExchangeRequest) (*gen.ExchangeResponse, error) {
	if req.Sum <= 0.0 {
		return nil, status.Error(codes.InvalidArgument, "invalid converting sum")
	}

	supCurrency := map[string]struct{}{"USD": {}, "RUB": {}, "EUR": {}}
	if _, ok := supCurrency[req.InputCurrency]; !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid input currency")
	}
	if _, ok := supCurrency[req.OutputCurrency]; !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid output currency")
	}

	return s.service.Exchange(ctx, req)
}
