package cache

import (
	configs "github.com/kVinsom/Bank-repository-service/internal/configs"
	"github.com/redis/go-redis/v9"
)

// NewRedisDB creates a bounded Redis client; connectivity is verified by the caller.
func NewRedisDB(cfg *configs.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Addr,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
	})
}
