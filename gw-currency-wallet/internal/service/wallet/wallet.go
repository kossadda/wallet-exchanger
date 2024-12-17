package wallet

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
)

type service struct {
	repo *storage.Storage
}

func newService(repo *storage.Storage) *service {
	return &service{
		repo: repo,
	}
}

func (w *service) GetBalance(userId int) (*model.Currency, error) {
	return w.repo.GetBalance(userId)
}

func (w *service) Deposit(dep *model.Operation) error {
	return w.repo.Deposit(dep)
}

func (w *service) Withdraw(dep *model.Operation) error {
	return w.repo.Withdraw(dep)
}
