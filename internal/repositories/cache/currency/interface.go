package currency

import (
	"context"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../mocks/cache/currency.go -package=mocks

// ICurrencyCache defines the behavior required at this layer boundary.
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
