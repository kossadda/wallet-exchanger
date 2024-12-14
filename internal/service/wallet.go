package service

import (
	"github.com/kossadda/wallet-exchanger/internal/repository"
	"github.com/kossadda/wallet-exchanger/model"
)

type WalletService struct {
	repo repository.Repository
}

func NewWalletService(repo repository.Repository) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (w *WalletService) GetBalance(userId int) (*model.Currency, error) {
	return w.repo.GetBalance(userId)
}

func (w *WalletService) DepositSum(dep *model.Deposit) error {
	return w.repo.DepositSum(dep)
}
