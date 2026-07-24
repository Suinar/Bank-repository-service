package credit

import (
	"context"

	common "github.com/Suinar/Bank-proto/repository/common"
	credit "github.com/Suinar/Bank-proto/repository/credit"
	service "github.com/kVinsom/Bank-repository-service/internal/services/credit"
)

// CreditHandler adapts gRPC requests to the application service contract.
type CreditHandler struct {
	credit.UnimplementedCreditRepositoryServer
	service service.ICreditService
}

// NewCreditHandler creates a ready-to-use credit handler.
func NewCreditHandler(service service.ICreditService) *CreditHandler {
	return &CreditHandler{service: service}
}

// GetAll returns all records available through CreditHandler.
func (h *CreditHandler) GetAll(ctx context.Context, req *common.Empty) (*credit.CreditList, error) {
	return h.service.GetAll(ctx, req)
}

// GetByUser returns records matching the requested user lookup.
func (h *CreditHandler) GetByUser(ctx context.Context, req *common.UserIdRequest) (*credit.CreditList, error) {
	return h.service.GetByUser(ctx, req)
}

// GetById returns records matching the requested id lookup.
func (h *CreditHandler) GetById(ctx context.Context, req *common.IdRequest) (*credit.Credit, error) {
	return h.service.GetById(ctx, req)
}

// Create persists a new record through CreditHandler.
func (h *CreditHandler) Create(ctx context.Context, req *credit.Credit) (*credit.Credit, error) {
	return h.service.Create(ctx, req)
}

// Repay applies a repayment to the requested credit through CreditHandler.
func (h *CreditHandler) Repay(ctx context.Context, req *common.AmountRequest) (*credit.Credit, error) {
	return h.service.Repay(ctx, req)
}

// Delete removes the requested record through CreditHandler.
func (h *CreditHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}
