package exchange

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// storage struct holds the database connection.
type storage struct {
	db database.DataBase
}

// newStorage creates a new storage instance.
func newStorage(db database.DataBase) *storage {
	return &storage{
		db: db,
	}
}

// GetExchangeRates fetches exchange rates from the database and returns them.
// @Summary Get exchange rates
// @Description Get the current exchange rates for various currencies from the database.
// @Tags Exchange
// @Accept json
// @Produce json
// @Success 200
// @Failure 500
func (e *storage) GetExchangeRates(ctx context.Context) (*gen.ExchangeRatesResponse, error) {
	var res gen.ExchangeRatesResponse

	if err := e.db.Transaction(func(tx *sqlx.Tx) error {
		query := "SELECT output, usd, rub, eur FROM currency"
		var rates []struct {
			Output string  `db:"output"`
			Usd    float32 `db:"usd"`
			Rub    float32 `db:"rub"`
			Eur    float32 `db:"eur"`
		}

		if err := tx.SelectContext(ctx, &rates, query); err != nil {
			return fmt.Errorf("failed to query exchange rates: %w", err)
		}

		res.Rates = make(map[string]*gen.OneCurrencyRate)

		for _, rate := range rates {
			res.Rates[rate.Output] = &gen.OneCurrencyRate{
				Rate: map[string]float32{
					"usd": rate.Usd,
					"rub": rate.Rub,
					"eur": rate.Eur,
				},
			}
		}

		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &res, nil
}
