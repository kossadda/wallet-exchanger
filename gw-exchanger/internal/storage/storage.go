package storage

import (
	"context"

	"github.com/kossadda/wallet-exchanger/share/database"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
)

type Exchanger interface {
	Exchange(ctx context.Context, request *gen.ExchangeRequest) (*gen.ExchangeResponse, error)
}

type Storage struct {
	Exchanger
}

func New(db database.DataBase) *Storage {
	return &Storage{
		Exchanger: NewExchangeDB(db),
	}
}
