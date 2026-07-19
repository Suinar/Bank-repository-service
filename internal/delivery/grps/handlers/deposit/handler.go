package deposit

import (
	"context"

	common "github.com/Suinar/Bank-proto/repository/common"
	deposit "github.com/Suinar/Bank-proto/repository/deposit"
	service "github.com/kVinsom/Bank-repository-service/internal/services/deposit"
)

// DepositHandler adapts gRPC requests to the application service contract.
type DepositHandler struct {
	deposit.UnimplementedDepositRepositoryServer
	service service.IDepositService
}

// NewDepositHandler creates a ready-to-use deposit handler.
func NewDepositHandler(service service.IDepositService) *DepositHandler {
	return &DepositHandler{service: service}
}

// GetAll returns all records available through DepositHandler.
func (h *DepositHandler) GetAll(ctx context.Context, req *common.Empty) (*deposit.DepositList, error) {
	return h.service.GetAll(ctx, req)
}

// GetByUser returns records matching the requested user lookup.
func (h *DepositHandler) GetByUser(ctx context.Context, req *common.UserIdRequest) (*deposit.DepositList, error) {
	return h.service.GetByUser(ctx, req)
}

// GetById returns records matching the requested id lookup.
func (h *DepositHandler) GetById(ctx context.Context, req *common.IdRequest) (*deposit.Deposit, error) {
	return h.service.GetById(ctx, req)
}

// Create persists a new record through DepositHandler.
func (h *DepositHandler) Create(ctx context.Context, req *deposit.Deposit) (*deposit.Deposit, error) {
	return h.service.Create(ctx, req)
}

// Replenish adds funds to the requested deposit through DepositHandler.
func (h *DepositHandler) Replenish(ctx context.Context, req *common.AmountRequest) (*deposit.Deposit, error) {
	return h.service.Replenish(ctx, req)
}

// Delete removes the requested record through DepositHandler.
func (h *DepositHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}
