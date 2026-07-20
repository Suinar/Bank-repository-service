package credit

import (
	"context"

	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/credit"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/credit.go -package=mocks

// ICreditService defines the behavior required at this layer boundary.
type ICreditService interface {
	GetAll(ctx context.Context, req *common.Empty) (*credit.CreditList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*credit.CreditList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*credit.Credit, error)
	Create(ctx context.Context, req *credit.Credit) (*credit.Credit, error)
	Repay(ctx context.Context, req *common.AmountRequest) (*credit.Credit, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
