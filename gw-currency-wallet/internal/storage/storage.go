// Package storage provides access to both authentication and wallet storage layers, making them available to the main application.
package storage

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage/auth"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage/wallet"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

// Storage aggregates the authentication and wallet storage layers into a single struct.
type Storage struct {
	*auth.Auth     // Embeds Auth for authentication-related operations
	*wallet.Wallet // Embeds Wallet for wallet-related operations
}

// New creates and returns a new instance of Storage, initializing both authentication and wallet subsystems.
func New(db database.DataBase) *Storage {
	return &Storage{
		Auth:   auth.New(db),
		Wallet: wallet.New(db),
	}
}
