package wallet

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

// MainOperations is an interface that defines methods for handling wallet operations such as balance checking, deposit, and withdrawal.
type MainOperations interface {
	// GetBalance to get user balance
	GetBalance(ctx *gin.Context)

	// Deposit to deposit funds into wallet
	Deposit(ctx *gin.Context)

	// Withdraw to withdraw funds from wallet
	Withdraw(ctx *gin.Context)
}

// Wallet represents the wallet handler with methods for managing balance and transactions.
type Wallet struct {
	MainOperations
}

// New creates and returns a new instance of the Wallet handler with the provided services, logger, and config.
func New(services *service.Service, logger *slog.Logger, cfg *configs.ServerConfig) *Wallet {
	return &Wallet{
		MainOperations: newHandler(services, logger, cfg),
	}
}
