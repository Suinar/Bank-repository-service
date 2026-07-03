package deposit

import (
	"Bank-repository-service/pkg/core"
	"context"
)

//go:generate mockgen -source=interfaces_repository.go -destination=../../../mocks/reposit/deposit.go -package=mocks

type IDepositRepository interface {
	GetAll(ctx context.Context) ([]core.Deposit, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Deposit, error)
	GetById(ctx context.Context, id int64) (*core.Deposit, error)
	Create(ctx context.Context, input *core.Deposit) (*core.Deposit, error)
	Replenish(ctx context.Context, id int64, amount int64) (*core.Deposit, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
