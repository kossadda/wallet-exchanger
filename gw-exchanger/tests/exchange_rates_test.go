package tests

import (
	"context"
	"syscall"
	"testing"
	"time"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/pkg/app"
	"github.com/kossadda/wallet-exchanger/gw-exchanger/tests/suite"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
)

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
		exchangeRates(t, s, ctx)
	})

	t.Run("exchangeRateForCurrency", func(t *testing.T) {
		resp, err := exchangeRateForCurrency(t, s, ctx, "USD", "RUB")
		if err != nil && resp != nil {
			t.Fatalf("error getting exchange rates: %v", err)
		}
		if expected := float32(103.6); resp.Rate != expected {
			t.Fatalf("expected rate USD to RUB %f, got %f", expected, resp.Rate)
		}
	})

	t.Run("exchangeRateForCurrencyError1", func(t *testing.T) {
		_, err := exchangeRateForCurrency(t, s, ctx, "US", "RUB")
		if err == nil {
			t.Fatal("excepted error, but got nil")
		}
	})
	t.Run("exchangeRateForCurrencyError2", func(t *testing.T) {
		_, err := exchangeRateForCurrency(t, s, ctx, "USD", "RU")
		if err == nil {
			t.Fatal("excepted error, but got nil")
		}
	})
	t.Run("exchangeRateForCurrencyEmpty", func(t *testing.T) {
		_, err := exchangeRateForCurrency(t, s, ctx, "", "")
		if err == nil {
			t.Fatal("excepted error, but got nil")
		}
	})

	go func() {
		time.Sleep(1 * time.Millisecond)
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	_ = application.GRPCSrv.Stop()
}

func exchangeRates(t *testing.T, s *suite.Suite, ctx context.Context) {
	resp, err := s.Client.GetExchangeRates(ctx, &gen.Empty{})
	if err != nil && resp != nil {
		t.Fatalf("error getting exchange rates: %v", err)
	}
}

func exchangeRateForCurrency(t *testing.T, s *suite.Suite, ctx context.Context, from, to string) (*gen.ExchangeRateResponse, error) {
	return s.Client.GetExchangeRateForCurrency(ctx, &gen.CurrencyRequest{
		FromCurrency: from,
		ToCurrency:   to,
	})
}
