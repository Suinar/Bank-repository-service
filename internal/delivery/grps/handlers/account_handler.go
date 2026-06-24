package handler

import (
	"context"

	service "Bank-repository-service/internal/services"
	account "Bank-repository-service/proto/repository/account"
	common "Bank-repository-service/proto/repository/common"
)

type AccountHandler struct {
	account.UnimplementedAccountRepositoryServer
	service service.IAccountService
}

func NewAccountHandler(service service.IAccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) GetAll(ctx context.Context, req *common.Empty) (*account.AccountList, error) {
	return h.service.GetAll(ctx, req)
}

func (h *AccountHandler) GetByUser(ctx context.Context, req *common.UserIdRequest) (*account.AccountList, error) {
	return h.service.GetByUser(ctx, req)
}

func (h *AccountHandler) GetById(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	return h.service.GetById(ctx, req)
}

func (h *AccountHandler) Create(ctx context.Context, req *account.Account) (*account.Account, error) {
	return h.service.Create(ctx, req)
}

func (h *AccountHandler) Blocking(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	return h.service.Blocking(ctx, req)
}

func (h *AccountHandler) Close(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	return h.service.Close(ctx, req)
}

func (h *AccountHandler) Update(ctx context.Context, req *account.UpdateAccountRequest) (*account.Account, error) {
	return h.service.Update(ctx, req)
}

func (h *AccountHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error) {
	return h.service.Delete(ctx, req)
}
