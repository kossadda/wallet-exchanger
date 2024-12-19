package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	expiry time.Duration
}

func New(ctx context.Context, cfg *configs.ServerConfig) (*Cache, error) {
	cacheAddr, ok := cfg.Servers["CACHE"]
	if !ok {
		return nil, fmt.Errorf("can't find cache server config")
	}

	db := redis.NewClient(&redis.Options{
		Addr: cacheAddr.Host + ":" + cacheAddr.Port,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis server: %s", err.Error())
	}

	cacheExp, err := time.ParseDuration(cfg.CacheExpire)
	if err != nil {
		cacheExp = configs.DefaultCacheExpire
	}

	return &Cache{
		client: db,
		expiry: cacheExp,
	}, nil
}

func (c *Cache) Set(ctx context.Context, key string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to serialize data: %w", err)
	}
	return c.client.Set(ctx, key, jsonData, c.expiry).Err()
}

func (c *Cache) Get(ctx context.Context, key string, v any) error {
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
