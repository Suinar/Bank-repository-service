package cache

import (
	currency "github.com/Suinar/Bank-repository-service/internal/repository/cache/currency"

	"github.com/redis/go-redis/v9"
)

type Caches struct {
	Currency *currency.CurrencyCache
}

func InitCaches(rdb *redis.Client) *Caches {
	return &Caches{
		Currency: currency.NewCurrencyCache(rdb),
	}
}



