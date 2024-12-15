package service

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/model"
)

type WalletService struct {
	repo storage.Repository
}

func NewWalletService(repo storage.Repository) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (w *WalletService) GetBalance(userId int) (*model.Currency, error) {
	return w.repo.GetBalance(userId)
}

func (w *WalletService) Deposit(dep *model.Operation) error {
	return w.repo.Deposit(dep)
}

func (w *WalletService) Withdraw(dep *model.Operation) error {
	return w.repo.Withdraw(dep)
}
