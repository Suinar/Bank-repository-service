package currency

import (
	"context"

	service "github.com/Suinar/Bank-repository-service/internal/services/currency"
	common "github.com/Suinar/Bank-proto/repository/common"
	currency "github.com/Suinar/Bank-proto/repository/currency"
)

type CurrencyHandler struct {
	currency.UnimplementedCurrencyRepositoryServer
	service service.ICurrencyService
}

func NewCurrencyHandler(service service.ICurrencyService) *CurrencyHandler {
	return &CurrencyHandler{service: service}
}

func (h *CurrencyHandler) GetAll(ctx context.Context, req *common.Empty) (*currency.CurrencyList, error) {
	return h.service.GetAll(ctx, req)
}

func (h *CurrencyHandler) GetById(ctx context.Context, req *common.IdRequest) (*currency.Currency, error) {
	return h.service.GetById(ctx, req)
}

func (h *CurrencyHandler) GetByIso(ctx context.Context, req *currency.IsoCodeRequest) (*currency.Currency, error) {
	return h.service.GetByIso(ctx, req)
}

func (h *CurrencyHandler) GetBySymbol(ctx context.Context, req *currency.SymbolRequest) (*currency.Currency, error) {
	return h.service.GetBySymbol(ctx, req)
}

func (h *CurrencyHandler) Create(ctx context.Context, req *currency.Currency) (*currency.Currency, error) {
	return h.service.Create(ctx, req)
}

func (h *CurrencyHandler) Update(ctx context.Context, req *currency.UpdateCurrencyRequest) (*currency.Currency, error) {
	return h.service.Update(ctx, req)
}

func (h *CurrencyHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}



