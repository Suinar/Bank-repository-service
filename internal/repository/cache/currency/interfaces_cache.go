package currency

import (
	"Bank-repository-service/pkg/core"
	"context"
)

//go:generate mockgen -source=interfaces_cache.go -destination=../../../mocks/currency_cache_mock.go -package=mocks

type ICurrencyCache interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	GetById(ctx context.Context, id int64) (*core.Currency, error)
	GetByIso(ctx context.Context, iso string) (*core.Currency, error)
	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	Set(ctx context.Context, currency *core.Currency) error
	SetAll(ctx context.Context, currencies []core.Currency) error
	Update(ctx context.Context, currency *core.Currency) error
	Delete(ctx context.Context, id int64) error
}

