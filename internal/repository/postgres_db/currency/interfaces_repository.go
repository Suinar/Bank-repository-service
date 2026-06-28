package currency

import (
	"Bank-repository-service/pkg/core"
	"context"
)

//go:generate mockgen -source=interfaces_repository.go -destination=../../../mocks/currency_repository_mock.go -package=mocks

type ICurrencyRepository interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	GetById(ctx context.Context, id int64) (*core.Currency, error)
	GetByIso(ctx context.Context, isoCode string) (*core.Currency, error)
	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	Create(ctx context.Context, input *core.Currency) (*core.Currency, error)
	Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error)
	Delete(ctx context.Context, id int64) error
}
