package handler

import (
	"context"

	service "Bank-repository-service/internal/services"
	common "Bank-repository-service/proto/repository/common"
	deposit "Bank-repository-service/proto/repository/deposit"
)

type DepositHandler struct {
	deposit.UnimplementedDepositRepositoryServer
	service service.IDepositService
}

func NewDepositHandler(service service.IDepositService) *DepositHandler {
	return &DepositHandler{service: service}
}

func (h *DepositHandler) GetAll(ctx context.Context, req *common.Empty) (*deposit.DepositList, error) {
	return h.service.GetAll(ctx, req)
}

func (h *DepositHandler) GetByUser(ctx context.Context, req *common.UserIdRequest) (*deposit.DepositList, error) {
	return h.service.GetByUser(ctx, req)
}

func (h *DepositHandler) GetById(ctx context.Context, req *common.IdRequest) (*deposit.Deposit, error) {
	return h.service.GetById(ctx, req)
}

func (h *DepositHandler) Create(ctx context.Context, req *deposit.Deposit) (*deposit.Deposit, error) {
	return h.service.Create(ctx, req)
}

func (h *DepositHandler) Replenish(ctx context.Context, req *common.AmountRequest) (*deposit.Deposit, error) {
	return h.service.Replenish(ctx, req)
}

func (h *DepositHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error) {
	return h.service.Delete(ctx, req)
}
