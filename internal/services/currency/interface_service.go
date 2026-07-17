package currency

import (
	"context"

	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/currency"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface_service.go -destination=../../../internal/mocks/services/currency.go -package=mocks

// ICurrencyService defines the behavior required at this layer boundary.
type ICurrencyService interface {
	GetAll(ctx context.Context, req *common.Empty) (*currency.CurrencyList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*currency.Currency, error)
	GetByIso(ctx context.Context, req *currency.IsoCodeRequest) (*currency.Currency, error)
	GetBySymbol(ctx context.Context, req *currency.SymbolRequest) (*currency.Currency, error)
	Create(ctx context.Context, req *currency.Currency) (*currency.Currency, error)
	Update(ctx context.Context, req *currency.UpdateCurrencyRequest) (*currency.Currency, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
