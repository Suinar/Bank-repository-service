package currency

import (
	"context"

	common "github.com/Suinar/Bank-proto/repository/common"
	currency "github.com/Suinar/Bank-proto/repository/currency"
	service "github.com/kVinsom/Bank-repository-service/internal/services/currency"
)

// CurrencyHandler adapts gRPC requests to the application service contract.
type CurrencyHandler struct {
	currency.UnimplementedCurrencyRepositoryServer
	service service.ICurrencyService
}

// NewCurrencyHandler creates a ready-to-use currency handler.
func NewCurrencyHandler(service service.ICurrencyService) *CurrencyHandler {
	return &CurrencyHandler{service: service}
}

// GetAll returns all records available through CurrencyHandler.
func (h *CurrencyHandler) GetAll(ctx context.Context, req *common.Empty) (*currency.CurrencyList, error) {
	return h.service.GetAll(ctx, req)
}

// GetById returns records matching the requested id lookup.
func (h *CurrencyHandler) GetById(ctx context.Context, req *common.IdRequest) (*currency.Currency, error) {
	return h.service.GetById(ctx, req)
}

// GetByIso returns records matching the requested iso lookup.
func (h *CurrencyHandler) GetByIso(ctx context.Context, req *currency.IsoCodeRequest) (*currency.Currency, error) {
	return h.service.GetByIso(ctx, req)
}

// GetBySymbol returns records matching the requested symbol lookup.
func (h *CurrencyHandler) GetBySymbol(ctx context.Context, req *currency.SymbolRequest) (*currency.Currency, error) {
	return h.service.GetBySymbol(ctx, req)
}

// Create persists a new record through CurrencyHandler.
func (h *CurrencyHandler) Create(ctx context.Context, req *currency.Currency) (*currency.Currency, error) {
	return h.service.Create(ctx, req)
}

// Update applies the requested changes through CurrencyHandler.
func (h *CurrencyHandler) Update(ctx context.Context, req *currency.UpdateCurrencyRequest) (*currency.Currency, error) {
	return h.service.Update(ctx, req)
}

// Delete removes the requested record through CurrencyHandler.
func (h *CurrencyHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}
