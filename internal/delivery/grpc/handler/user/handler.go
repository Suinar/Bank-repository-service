package user

import (
	"context"

	common "github.com/Suinar/Bank-proto/repository/common"
	user "github.com/Suinar/Bank-proto/repository/user"
	service "github.com/kVinsom/Bank-repository-service/internal/services/user"
)

// UserHandler adapts gRPC requests to the application service contract.
type UserHandler struct {
	user.UnimplementedUserRepositoryServer
	service service.IUserService
}

// NewUserHandler creates a ready-to-use user handler.
func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAll returns all records available through UserHandler.
func (h *UserHandler) GetAll(ctx context.Context, req *common.Empty) (*user.UserList, error) {
	return h.service.GetAll(ctx, req)
}

// GetById returns records matching the requested id lookup.
func (h *UserHandler) GetById(ctx context.Context, req *common.IdRequest) (*user.User, error) {
	return h.service.GetById(ctx, req)
}

// GetByEmail returns records matching the requested email lookup.
func (h *UserHandler) GetByEmail(ctx context.Context, req *user.EmailRequest) (*user.User, error) {
	return h.service.GetByEmail(ctx, req)
}

// GetByPhoneNumber returns records matching the requested phone number lookup.
func (h *UserHandler) GetByPhoneNumber(ctx context.Context, req *user.PhoneNumberRequest) (*user.User, error) {
	return h.service.GetByPhoneNumber(ctx, req)
}

// Create persists a new record through UserHandler.
func (h *UserHandler) Create(ctx context.Context, req *user.User) (*user.User, error) {
	return h.service.Create(ctx, req)
}

// Update applies the requested changes through UserHandler.
func (h *UserHandler) Update(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	return h.service.Update(ctx, req)
}

// ChangePassword updates the password for the requested user.
func (h *UserHandler) ChangePassword(ctx context.Context, req *user.ChangePasswordRequest) (*common.Empty, error) {
	return h.service.ChangePassword(ctx, req)
}

// Delete removes the requested record through UserHandler.
func (h *UserHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}
