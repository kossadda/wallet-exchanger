package tests

import (
	"bytes"
	"io"
	"net/http"
	"syscall"
	"testing"
	"time"

	"github.com/kossadda/wallet-exchanger/currency-wallet/tests/suite"
	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	t.Parallel()
	_, s := suite.New(t)
	go s.App.MustRun()

	testRegister(t)
	testLogin(t)
	testGetBalance(t)
	testDeposit(t)
	testWithdraw(t)
	testExchangeSum(t)
	testGetExchangeRate(t)

	go func() {
		time.Sleep(1 * time.Millisecond)
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	_ = s.App.Stop()
	s.Ctrl.Finish()
}

func testRegister(t *testing.T) {
	t.Run("CreateUser", func(t *testing.T) {
		ctx, s := suite.New(t)
		loginData := []byte(`{"username": "test", "password": "qwerty", "email": "test@example.com"}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(loginData)),
		}

		s.Hnd.Register(ctx)
		require.Equal(t, 0, len(ctx.Errors))
	})

	t.Run("CreateUserError1", func(t *testing.T) {
		ctx, s := suite.New(t)
		loginData := []byte(`{""}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(loginData)),
		}

		s.Hnd.Register(ctx)
		require.NotEqual(t, 0, len(ctx.Errors))
	})

	t.Run("CreateUserError2", func(t *testing.T) {
		ctx, s := suite.New(t)
		loginData := []byte(`{"username": "1", "password": "2", "email": "3"}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(loginData)),
		}

		s.Hnd.Register(ctx)
		require.NotEqual(t, 0, len(ctx.Errors))
	})
}

func testLogin(t *testing.T) {
	t.Run("Login", func(t *testing.T) {
		ctx, s := suite.New(t)
		loginData := []byte(`{"username": "test", "password": "qwerty"}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(loginData)),
		}

		s.Hnd.Login(ctx)
		require.Equal(t, 0, len(ctx.Errors))
	})

	t.Run("LoginError", func(t *testing.T) {
		ctx, s := suite.New(t)
		loginData := []byte(`{""}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(loginData)),
		}

		s.Hnd.Login(ctx)
		require.NotEqual(t, 0, len(ctx.Errors))
	})
}

func testGetBalance(t *testing.T) {
	t.Run("GetBalance", func(t *testing.T) {
		ctx, s := suite.New(t)
		ctx.Set("userId", 1)
		s.Hnd.GetBalance(ctx)
		require.Equal(t, 0, len(ctx.Errors))
	})

	t.Run("GetBalanceError", func(t *testing.T) {
		ctx, s := suite.New(t)
		ctx.Set("userId", 2)
		s.Hnd.GetBalance(ctx)
		require.NotEqual(t, 0, len(ctx.Errors))
	})
}

func testDeposit(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		ctx, s := suite.New(t)
		depositData := []byte(`{"amount": 150.0, "currency": "USD"}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(depositData)),
		}
		ctx.Set("userId", 1)

		s.Hnd.Deposit(ctx)
		require.Equal(t, 0, len(ctx.Errors))
	})

	t.Run("DepositError1", func(t *testing.T) {
		ctx, s := suite.New(t)
		depositData := []byte(`{"amount": 150.0, "currency": "USD"}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(depositData)),
		}
		ctx.Set("userId", 2)

		s.Hnd.Deposit(ctx)
		require.NotEqual(t, 0, len(ctx.Errors))
	})

	t.Run("DepositError2", func(t *testing.T) {
		ctx, s := suite.New(t)
		depositData := []byte(`{}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(depositData)),
		}
		ctx.Set("userId", 2)

		s.Hnd.Deposit(ctx)
		require.NotEqual(t, 0, len(ctx.Errors))
	})
}

func testWithdraw(t *testing.T) {
	t.Run("Withdraw", func(t *testing.T) {
		ctx, s := suite.New(t)
		depositData := []byte(`{"amount": 150.0, "currency": "USD"}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(depositData)),
		}
		ctx.Set("userId", 1)

		s.Hnd.Withdraw(ctx)
		require.Equal(t, 0, len(ctx.Errors))
	})

	t.Run("WithdrawError1", func(t *testing.T) {
		ctx, s := suite.New(t)
		depositData := []byte(`{"amount": 150.0, "currency": "USD"}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(depositData)),
		}
		ctx.Set("userId", 2)

		s.Hnd.Withdraw(ctx)
		require.NotEqual(t, 0, len(ctx.Errors))
	})

	t.Run("WithdrawError2", func(t *testing.T) {
		ctx, s := suite.New(t)
		depositData := []byte(`{}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(depositData)),
		}
		ctx.Set("userId", 2)

		s.Hnd.Withdraw(ctx)
		require.NotEqual(t, 0, len(ctx.Errors))
	})
}

func testGetExchangeRate(t *testing.T) {
	t.Run("GetExchangeRate", func(t *testing.T) {
		ctx, s := suite.New(t)
		ctx.Set("userId", 1)
		s.Hnd.GetExchangeRates(ctx)
		require.Equal(t, 0, len(ctx.Errors))
	})
}

func testExchangeSum(t *testing.T) {
	t.Run("ExchangeSum", func(t *testing.T) {
		ctx, s := suite.New(t)
		exchangeData := []byte(`{"from_currency": "USD", "to_currency": "EUR", "amount": 100.00}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(exchangeData)),
		}
		ctx.Set("userId", 1)

		s.Hnd.ExchangeSum(ctx)

		require.Equal(t, 0, len(ctx.Errors))
	})

	t.Run("ExchangeSumError", func(t *testing.T) {
		ctx, s := suite.New(t)
		exchangeData := []byte(`{}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(exchangeData)),
		}
		ctx.Set("userId", 1)

		s.Hnd.ExchangeSum(ctx)

		require.NotEqual(t, 0, len(ctx.Errors))
	})

	t.Run("ExchangeSum", func(t *testing.T) {
		ctx, s := suite.New(t)
		exchangeData := []byte(`{"from_currency": "USD", "to_currency": "EUR", "amount": 100.00}`)
		ctx.Request = &http.Request{
			Body: io.NopCloser(bytes.NewReader(exchangeData)),
		}
		ctx.Set("userId", 2)

		s.Hnd.ExchangeSum(ctx)

		require.NotEqual(t, 0, len(ctx.Errors))
	})
}
