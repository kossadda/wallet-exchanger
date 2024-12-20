package tests

import (
	"syscall"
	"testing"
	"time"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/pkg/app"
	"github.com/kossadda/wallet-exchanger/gw-exchanger/tests/suite"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// It is necessary to deploy a postgres server on localhost:5436
// with the relations specified in the file root/db/migrations
func TestAll(t *testing.T) {
	ctx, s := suite.New(t)
	dbConf := &configs.ConfigDB{
		DBHost:     "localhost",
		DBPort:     "5436",
		DBUser:     "postgres",
		DBPassword: "qwerty",
		DBName:     "postgres",
		DBSSLMode:  "disable",
	}

	log := logger.SetupByEnv(s.Cfg.Env)

	application := app.New(log, dbConf, s.Cfg)

	go application.GRPCSrv.MustRun()

	t.Run("exchangeRates", func(t *testing.T) {
		resp, err := s.Client.GetExchangeRates(ctx, &gen.Empty{})
		require.NoError(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("exchangeRateForCurrency", func(t *testing.T) {
		resp, err := s.Client.GetExchangeRateForCurrency(ctx, &gen.CurrencyRequest{
			FromCurrency: "USD",
			ToCurrency:   "RUB",
		})
		require.NoError(t, err)
		require.Equal(t, float32(103.6), resp.Rate)
	})

	t.Run("exchangeRateForCurrencyError", func(t *testing.T) {
		tests := []struct {
			name string
			from string
			to   string
		}{
			{
				name: "wrong input",
				from: "US",
				to:   "RUB",
			},
			{
				name: "wrong input",
				from: "USD",
				to:   "RU",
			},
			{
				name: "wrong input",
				from: "",
				to:   "",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				resp, err := s.Client.GetExchangeRateForCurrency(ctx, &gen.CurrencyRequest{
					FromCurrency: tt.from,
					ToCurrency:   tt.to,
				})
				assert.Error(t, err)
				assert.Empty(t, resp)
			})
		}
	})

	go func() {
		time.Sleep(1 * time.Millisecond)
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	_ = application.GRPCSrv.Stop()
}
