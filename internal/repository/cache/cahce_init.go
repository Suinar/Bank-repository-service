package cache

import currency "Bank-repository-service/internal/repository/cache/currency"

type Caches struct {
	currency currency.ICurrencyCache
}

func InitCaches(currency currency.ICurrencyCache) *Caches {
	return &Caches{currency: currency}
}
