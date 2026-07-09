package deposit

import (
	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/deposit"
	"context"
)

//go:generate mockgen -source=interface_service.go -destination=../../../internal/mocks/service/deposit.go -package=mocks

type IDepositService interface {
	GetAll(ctx context.Context, req *common.Empty) (*deposit.DepositList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*deposit.DepositList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*deposit.Deposit, error)
	Create(ctx context.Context, req *deposit.Deposit) (*deposit.Deposit, error)
	Replenish(ctx context.Context, req *common.AmountRequest) (*deposit.Deposit, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}


