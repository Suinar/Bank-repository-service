package account

import (
	"Bank-repository-service/pkg/core"
	"context"
)

//go:generate mockgen -source=interfaces_repository.go -destination=../../../mocks/account_repository_mock.go -package=mocks

type IAccountRepository interface {
	GetAll(ctx context.Context) ([]core.Account, error)
	GetByUser(ctx context.Context, userId int64) ([]core.Account, error)
	GetById(ctx context.Context, id int64) (*core.Account, error)
	Create(ctx context.Context, input *core.Account) (*core.Account, error)
	Blocking(ctx context.Context, id int64) (*core.Account, error)
	Close(ctx context.Context, id int64) (*core.Account, error)
	Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
