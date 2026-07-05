package deposit

import (
	"Bank-repository-service/proto/repository/common"
	"Bank-repository-service/proto/repository/deposit"
	"context"
)

type IDepositHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*deposit.DepositList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*deposit.DepositList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*deposit.Deposit, error)
	Create(ctx context.Context, req *deposit.Deposit) (*deposit.Deposit, error)
	Replenish(ctx context.Context, req *common.AmountRequest) (*deposit.Deposit, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
