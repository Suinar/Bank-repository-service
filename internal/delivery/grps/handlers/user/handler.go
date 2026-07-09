package user

import (
	"context"

	service "github.com/Suinar/Bank-repository-service/internal/services/user"
	common "github.com/Suinar/Bank-proto/repository/common"
	user "github.com/Suinar/Bank-proto/repository/user"
)

type UserHandler struct {
	user.UnimplementedUserRepositoryServer
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(ctx context.Context, req *common.Empty) (*user.UserList, error) {
	return h.service.GetAll(ctx, req)
}

func (h *UserHandler) GetById(ctx context.Context, req *common.IdRequest) (*user.User, error) {
	return h.service.GetById(ctx, req)
}

func (h *UserHandler) GetByEmail(ctx context.Context, req *user.EmailRequest) (*user.User, error) {
	return h.service.GetByEmail(ctx, req)
}

func (h *UserHandler) GetByPhoneNumber(ctx context.Context, req *user.PhoneNumberRequest) (*user.User, error) {
	return h.service.GetByPhoneNumber(ctx, req)
}

func (h *UserHandler) Create(ctx context.Context, req *user.User) (*user.User, error) {
	return h.service.Create(ctx, req)
}

func (h *UserHandler) Update(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	return h.service.Update(ctx, req)
}

func (h *UserHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}



