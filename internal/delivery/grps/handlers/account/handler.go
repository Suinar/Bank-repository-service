package account

import (
	"context"

	account "github.com/Suinar/Bank-proto/repository/account"
	common "github.com/Suinar/Bank-proto/repository/common"
	service "github.com/Suinar/Bank-repository-service/internal/services/account"
)

// AccountHandler adapts gRPC requests to the application service contract.
type AccountHandler struct {
	account.UnimplementedAccountRepositoryServer
	service service.IAccountService
}

// NewAccountHandler creates a ready-to-use account handler.
func NewAccountHandler(service service.IAccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

// GetAll returns all records available through AccountHandler.
func (h *AccountHandler) GetAll(ctx context.Context, req *common.Empty) (*account.AccountList, error) {
	return h.service.GetAll(ctx, req)
}

// GetByUser returns records matching the requested user lookup.
func (h *AccountHandler) GetByUser(ctx context.Context, req *common.UserIdRequest) (*account.AccountList, error) {
	return h.service.GetByUser(ctx, req)
}

// GetById returns records matching the requested id lookup.
func (h *AccountHandler) GetById(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	return h.service.GetById(ctx, req)
}

// Create persists a new record through AccountHandler.
func (h *AccountHandler) Create(ctx context.Context, req *account.Account) (*account.Account, error) {
	return h.service.Create(ctx, req)
}

// Blocking moves the requested record to its blocked state through AccountHandler.
func (h *AccountHandler) Blocking(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	return h.service.Blocking(ctx, req)
}

// Close moves the requested account to its closed state through AccountHandler.
func (h *AccountHandler) Close(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	return h.service.Close(ctx, req)
}

// Update applies the requested changes through AccountHandler.
func (h *AccountHandler) Update(ctx context.Context, req *account.UpdateAccountRequest) (*account.Account, error) {
	return h.service.Update(ctx, req)
}

// Delete removes the requested record through AccountHandler.
func (h *AccountHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}
