package deposit

import (
	"context"
	"github.com/Suinar/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interfaces_repository.go -destination=../../../mocks/repository/deposit.go -package=mocks

// IDepositRepository defines the behavior required at this layer boundary.
type IDepositRepository interface {
	GetAll(ctx context.Context) ([]core.Deposit, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Deposit, error)
	GetById(ctx context.Context, id int64) (*core.Deposit, error)
	Create(ctx context.Context, input *core.Deposit) (*core.Deposit, error)
	Replenish(ctx context.Context, id int64, amount int64) (*core.Deposit, error)
	Delete(ctx context.Context, id int64) error
}
