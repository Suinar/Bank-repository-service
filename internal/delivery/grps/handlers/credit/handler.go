package credit

import (
	"context"

	service "Bank-repository-service/internal/services/credit"
	common "Bank-repository-service/proto/repository/common"
	credit "Bank-repository-service/proto/repository/credit"
)

type CreditHandler struct {
	credit.UnimplementedCreditRepositoryServer
	service service.ICreditService
}

func NewCreditHandler(service service.ICreditService) *CreditHandler {
	return &CreditHandler{service: service}
}

func (h *CreditHandler) GetAll(ctx context.Context, req *common.Empty) (*credit.CreditList, error) {
	return h.service.GetAll(ctx, req)
}

func (h *CreditHandler) GetByUser(ctx context.Context, req *common.UserIdRequest) (*credit.CreditList, error) {
	return h.service.GetByUser(ctx, req)
}

func (h *CreditHandler) GetById(ctx context.Context, req *common.IdRequest) (*credit.Credit, error) {
	return h.service.GetById(ctx, req)
}

func (h *CreditHandler) Create(ctx context.Context, req *credit.Credit) (*credit.Credit, error) {
	return h.service.Create(ctx, req)
}

func (h *CreditHandler) Repay(ctx context.Context, req *common.AmountRequest) (*credit.Credit, error) {
	return h.service.Repay(ctx, req)
}

func (h *CreditHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}
