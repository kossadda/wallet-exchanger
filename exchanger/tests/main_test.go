package tests

import (
	"syscall"
	"testing"
	"time"

	"github.com/kossadda/wallet-exchanger/exchanger/tests/suite"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	ctx, s := suite.New(t)

	go s.App.MustRun()

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

	_ = s.App.Stop()
	s.Ctrl.Finish()
}
