package cache

import (
	currency "github.com/kVinsom/Bank-repository-service/internal/repositories/cache/currency"

	"github.com/redis/go-redis/v9"
)

// Caches groups the Redis-backed cache dependencies used by services.
type Caches struct {
	Currency *currency.CurrencyCache
}

// InitCaches wires the dependencies required by caches.
func InitCaches(rdb *redis.Client) *Caches {
	return &Caches{
		Currency: currency.NewCurrencyCache(rdb),
	}
}
