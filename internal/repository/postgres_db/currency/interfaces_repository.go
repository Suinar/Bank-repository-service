package currency

import (
	"context"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interfaces_repository.go -destination=../../../mocks/repository/currency.go -package=mocks

// ICurrencyRepository defines the behavior required at this layer boundary.
type ICurrencyRepository interface {
	GetAll(ctx context.Context) ([]core.Currency, error)
	GetById(ctx context.Context, id int64) (*core.Currency, error)
	GetByIso(ctx context.Context, isoCode string) (*core.Currency, error)
	GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error)
	Create(ctx context.Context, input *core.Currency) (*core.Currency, error)
	Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error)
	Delete(ctx context.Context, id int64) error
}
