package wallet

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

type MainOperations interface {
	GetBalance(ctx *gin.Context)
	Deposit(ctx *gin.Context)
	Withdraw(ctx *gin.Context)
}

type Wallet struct {
	MainOperations
}

func NewWallet(services *service.Service, logger *slog.Logger, cfg *configs.ServerConfig) *Wallet {
	return &Wallet{
		MainOperations: newHandler(services, logger, cfg),
	}
}
