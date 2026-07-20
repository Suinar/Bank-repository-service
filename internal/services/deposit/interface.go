package deposit

import (
	"context"

	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/deposit"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/deposit.go -package=mocks

// IDepositService defines the behavior required at this layer boundary.
type IDepositService interface {
	GetAll(ctx context.Context, req *common.Empty) (*deposit.DepositList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*deposit.DepositList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*deposit.Deposit, error)
	Create(ctx context.Context, req *deposit.Deposit) (*deposit.Deposit, error)
	Replenish(ctx context.Context, req *common.AmountRequest) (*deposit.Deposit, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
