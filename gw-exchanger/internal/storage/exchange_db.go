package storage

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kossadda/wallet-exchanger/share/database"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
)

type ExchangeDB struct {
	db database.DataBase
}

func (e *ExchangeDB) Exchange(ctx context.Context, request *gen.ExchangeRequest) (*gen.ExchangeResponse, error) {
	return nil, status.Error(codes.Unavailable, "not implemented")
}

func NewExchangeDB(db database.DataBase) *ExchangeDB {
	return &ExchangeDB{
		db: db,
	}
}
