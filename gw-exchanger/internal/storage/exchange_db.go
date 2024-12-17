package storage

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"

	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

type ExchangeDB struct {
	db database.DataBase
}

func NewExchangeDB(db database.DataBase) *ExchangeDB {
	return &ExchangeDB{
		db: db,
	}
}

func (e *ExchangeDB) Exchange(ctx context.Context, request *gen.ExchangeRequest) (*gen.ExchangeResponse, error) {
	var res gen.ExchangeResponse
	if err := e.db.Transaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("SELECT %s FROM %s WHERE output=$1", strings.ToLower(request.InputCurrency), database.CurrencyTable)
		var coefficient float64
		if err := tx.GetContext(ctx, &coefficient, query, strings.ToLower(request.OutputCurrency)); err != nil {
			return fmt.Errorf("failed to get exchange coefficient: %w", err)
		}

		res.Sum = request.Sum * coefficient
		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &res, nil
}
