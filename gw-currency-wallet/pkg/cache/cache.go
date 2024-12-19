package cache

import (
	"context"

	r "github.com/kossadda/wallet-exchanger/gw-currency-wallet/pkg/cache/redis"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

type Cache interface {
	Set(ctx context.Context, key string, data any) error
	Get(ctx context.Context, key string, v any) error
}

func NewRedis(ctx context.Context, cfg *configs.ServerConfig) (Cache, error) {
	return r.New(ctx, cfg)
}
