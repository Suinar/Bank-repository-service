package credit

import (
	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/credit"
	"context"
)

type ICreditHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*credit.CreditList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*credit.CreditList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*credit.Credit, error)
	Create(ctx context.Context, req *credit.Credit) (*credit.Credit, error)
	Repay(ctx context.Context, req *common.AmountRequest) (*credit.Credit, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}


