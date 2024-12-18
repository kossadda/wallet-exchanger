package storage

import (
	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/storage/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

type Storage struct {
	*exchange.Exchange
}

func New(db database.DataBase) *Storage {
	return &Storage{
		Exchange: exchange.New(db),
	}
}
