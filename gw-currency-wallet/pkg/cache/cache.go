// Package cache provides an abstraction for caching mechanisms used by the application.
// It allows setting and getting data from a cache backend (e.g., Redis).
package cache

import (
	"context"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/pkg/cache/redis"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

// Cache defines the interface for caching operations. It provides methods for storing and retrieving data.
type Cache interface {
	// Set stores data in the cache with a specified key.
	Set(ctx context.Context, key string, data any) error

	// Get retrieves data from the cache by the specified key.
	Get(ctx context.Context, key string, v any) error
}

// NewRedis initializes a new Redis cache client and returns a Cache instance.
// It takes the context and server configuration to configure the Redis client.
func NewRedis(ctx context.Context, cfg *configs.ServerConfig) (Cache, error) {
	return redis.New(ctx, cfg)
}
