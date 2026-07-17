package currency

import (
	"context"
	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/currency"
)

// ICurrencyHandler defines the behavior required at this layer boundary.
type ICurrencyHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*currency.CurrencyList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*currency.Currency, error)
	GetByIso(ctx context.Context, req *currency.IsoCodeRequest) (*currency.Currency, error)
	GetBySymbol(ctx context.Context, req *currency.SymbolRequest) (*currency.Currency, error)
	Create(ctx context.Context, req *currency.Currency) (*currency.Currency, error)
	Update(ctx context.Context, req *currency.UpdateCurrencyRequest) (*currency.Currency, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
