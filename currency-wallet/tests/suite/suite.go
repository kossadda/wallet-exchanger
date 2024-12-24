package suite

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	walletapp "github.com/kossadda/wallet-exchanger/currency-wallet/internal/app"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service"
	srvauth "github.com/kossadda/wallet-exchanger/currency-wallet/internal/service/auth"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service/grpcclient"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service/grpcclient/mockclient"
	srvwallet "github.com/kossadda/wallet-exchanger/currency-wallet/internal/service/wallet"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage"
	strgauth "github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage/auth"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage/auth/mockauth"
	strgwallet "github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage/wallet"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage/wallet/mockwallet"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
)

const (
	AppTestPort = "44043"
)

type Suite struct {
	App  *walletapp.WalletApp
	Hnd  *delivery.Handler
	Ctrl *gomock.Controller
}

type FakeDB struct{}

func (f *FakeDB) Transaction(fn func(tx *sqlx.Tx) error) error {
	return nil
}

func (f *FakeDB) Close() error {
	return nil
}

func New(t *testing.T) (*gin.Context, *Suite) {
	t.Helper()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	t.Cleanup(func() {
		t.Helper()
		ctx.Done()
	})

	cfg := &configs.ServerConfig{
		Env:         "prod",
		TokenExpire: "10h",
		Servers: map[string]configs.Server{
			"APP": configs.Server{
				Host: "localhost",
				Port: AppTestPort,
			},
		},
	}

	log := logger.SetupByEnv(cfg.Env)
	ctrl := gomock.NewController(t)
	srv := fakeService(ctx, ctrl)
	hnd := delivery.New(srv, log, cfg)

	appAddr, _ := cfg.Servers["APP"]

	return ctx, &Suite{
		App:  walletapp.New(log, hnd, appAddr.Port),
		Hnd:  hnd,
		Ctrl: ctrl,
	}
}

func fakeService(ctx *gin.Context, ctrl *gomock.Controller) *service.Service {
	strg := fakeStorage(ctrl)

	return &service.Service{
		Storage:  strg,
		Auth:     srvauth.New(strg),
		Wallet:   srvwallet.New(strg),
		Exchange: fakeGrpcService(ctx, ctrl),
	}
}

func fakeGrpcService(ctx *gin.Context, ctrl *gomock.Controller) *grpcclient.Exchange {
	mockClient := mockclient.NewMockMainAPI(ctrl)

	response := &gen.ExchangeRatesResponse{
		Rates: map[string]*gen.OneCurrencyRate{
			"usd": &gen.OneCurrencyRate{
				Rate: map[string]float32{
					"usd": 1.0,
					"rub": 0.0097,
					"eur": 1.05,
				},
			},
			"rub": &gen.OneCurrencyRate{
				Rate: map[string]float32{
					"usd": 103.6,
					"rub": 1.0,
					"eur": 108.89,
				},
			},
			"eur": &gen.OneCurrencyRate{
				Rate: map[string]float32{
					"usd": 0.95,
					"rub": 0.0092,
					"eur": 1.0,
				},
			},
		},
	}

	mockClient.EXPECT().GetExchangeRates(ctx).Return(response, nil).AnyTimes()
	mockClient.EXPECT().ExchangeSum(ctx, &model.Exchange{UserId: 1, FromCurrency: "USD", ToCurrency: "EUR", Amount: 100.0}).Return([]float64{0, 0.95}, nil).AnyTimes()
	mockClient.EXPECT().ExchangeSum(ctx, &model.Exchange{UserId: 2, FromCurrency: "USD", ToCurrency: "EUR", Amount: 100.0}).Return(nil, fmt.Errorf("empty req")).AnyTimes()
	mockClient.EXPECT().CloseGRPC().Return(nil).AnyTimes()

	return &grpcclient.Exchange{
		MainAPI: mockClient,
	}
}

func fakeStorage(ctrl *gomock.Controller) *storage.Storage {
	return &storage.Storage{
		DataBase: &FakeDB{},
		Auth: &strgauth.Auth{
			MainAPI: fakeAuthStorage(ctrl),
		},
		Wallet: &strgwallet.Wallet{
			MainAPI: fakeWalletStorage(ctrl),
		},
	}
}

func fakeAuthStorage(ctrl *gomock.Controller) *mockauth.MockMainAPI {
	mockAuth := mockauth.NewMockMainAPI(ctrl)

	gomock.InOrder(
		mockAuth.EXPECT().CreateUser(model.User{
			Username: "test",
			Password: "74657374b1b3773a05c0ed0176787a4f1574ff0075f7521e",
			Email:    "test@example.com",
		}).Return(nil).AnyTimes(),
		mockAuth.EXPECT().GetUser("test", "74657374b1b3773a05c0ed0176787a4f1574ff0075f7521e").Return(&model.User{
			Id:       1,
			Username: "test",
			Password: "qwerty",
			Email:    "test@example.com",
		}, nil).AnyTimes(),
	)
	mockAuth.EXPECT().CreateUser(model.User{
		Username: "1",
		Password: "31da4b9237bacccdf19c0760cab7aec4a8359010b0",
		Email:    "3",
	}).Return(fmt.Errorf("empty user")).AnyTimes()

	return mockAuth
}

func fakeWalletStorage(ctrl *gomock.Controller) *mockwallet.MockMainAPI {
	mockWallet := mockwallet.NewMockMainAPI(ctrl)

	mockWallet.EXPECT().GetBalance(1).Return(&model.Currency{
		USD: 1.0,
		RUB: 100.0,
		EUR: 2.0,
	}, nil).AnyTimes()
	mockWallet.EXPECT().GetBalance(2).Return(nil, fmt.Errorf("user 2 does not exists")).AnyTimes()
	mockWallet.EXPECT().Deposit(&model.Operation{
		UserId:   1,
		Amount:   150.0,
		Currency: "USD",
	}).Return(150.0, nil).AnyTimes()
	mockWallet.EXPECT().Deposit(&model.Operation{
		UserId:   2,
		Amount:   150.0,
		Currency: "USD",
	}).Return(0.0, fmt.Errorf("user 2 does not exists")).AnyTimes()
	mockWallet.EXPECT().Withdraw(&model.Operation{
		UserId:   1,
		Amount:   150.0,
		Currency: "USD",
	}).Return(50.0, nil).AnyTimes()
	mockWallet.EXPECT().Withdraw(&model.Operation{
		UserId:   2,
		Amount:   150.0,
		Currency: "USD",
	}).Return(0.0, fmt.Errorf("user 2 does not exists")).AnyTimes()

	return mockWallet
}
