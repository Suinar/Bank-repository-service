package credit

import (
	"Bank-repository-service/pkg/core"
	"context"
)

//go:generate mockgen -source=interfaces_repository.go -destination=../../../mocks/repository/credit.go -package=mocks

type ICreditRepository interface {
	GetAll(ctx context.Context) ([]core.Credit, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Credit, error)
	GetById(ctx context.Context, id int64) (*core.Credit, error)
	Create(ctx context.Context, input *core.Credit) (*core.Credit, error)
	Repay(ctx context.Context, id int64, amount int64) (*core.Credit, error)
	Delete(ctx context.Context, id int64) error
}
