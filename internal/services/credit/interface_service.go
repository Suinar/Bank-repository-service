package credit

import (
	"Bank-repository-service/proto/repository/common"
	"Bank-repository-service/proto/repository/credit"
	"context"
)

//go:generate mockgen -source=interface_service.go -destination=../../../internal/mocks/service/credit.go -package=mocks

type ICreditService interface {
	GetAll(ctx context.Context, req *common.Empty) (*credit.CreditList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*credit.CreditList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*credit.Credit, error)
	Create(ctx context.Context, req *credit.Credit) (*credit.Credit, error)
	Repay(ctx context.Context, req *common.AmountRequest) (*credit.Credit, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error)
}
