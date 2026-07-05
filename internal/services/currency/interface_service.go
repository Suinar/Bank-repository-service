package currency

import (
	"Bank-repository-service/proto/repository/common"
	"Bank-repository-service/proto/repository/currency"
	"context"
)

//go:generate mockgen -source=interface_service.go -destination=../../../internal/mocks/service/currency.go -package=mocks

type ICurrencyService interface {
	GetAll(ctx context.Context, req *common.Empty) (*currency.CurrencyList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*currency.Currency, error)
	GetByIso(ctx context.Context, req *currency.IsoCodeRequest) (*currency.Currency, error)
	GetBySymbol(ctx context.Context, req *currency.SymbolRequest) (*currency.Currency, error)
	Create(ctx context.Context, req *currency.Currency) (*currency.Currency, error)
	Update(ctx context.Context, req *currency.UpdateCurrencyRequest) (*currency.Currency, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
