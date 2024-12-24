// Package redis provides a Redis-based caching implementation for the application.
// It includes functionality to store and retrieve data in a Redis cache.
package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/redis/go-redis/v9"
)

// Redis represents a Redis cache client. It includes the Redis client and cache expiry time.
type Redis struct {
	client *redis.Client
	expiry time.Duration
}

// New initializes a new Redis cache client using the provided server configuration.
// It returns a Redis instance or an error if the connection to Redis fails.
func New(ctx context.Context, addr, cacheExpire string) (*Redis, error) {
	db := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis server: %s", err.Error())
	}

	cacheExp, err := time.ParseDuration(cacheExpire)
	if err != nil {
		cacheExp = configs.DefaultCacheExpire
	}

	return &Redis{
		client: db,
		expiry: cacheExp,
	}, nil
}

// Set stores data in the Redis cache under the specified key. It serializes the data into JSON format
// and sets the expiration time based on the Redis's expiry.
func (c *Redis) Set(ctx context.Context, key string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to serialize data: %w", err)
	}
	return c.client.Set(ctx, key, jsonData, c.expiry).Err()
}

// Get retrieves data from the Redis cache using the specified key. It deserializes the cached data
// from JSON format into the provided value.
func (c *Redis) Get(ctx context.Context, key string, v any) error {
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
