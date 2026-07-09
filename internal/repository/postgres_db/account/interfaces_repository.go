package account

import (
	"github.com/Suinar/Bank-exhange-rate-service/pkg/core"
	"context"
)

//go:generate mockgen -source=interfaces_repository.go -destination=../../../mocks/repository/account.go -package=mocks

type IAccountRepository interface {
	GetAll(ctx context.Context) ([]core.Account, error)
	GetByUser(ctx context.Context, userId int64) ([]core.Account, error)
	GetById(ctx context.Context, id int64) (*core.Account, error)
	Create(ctx context.Context, input *core.Account) (*core.Account, error)
	Blocking(ctx context.Context, id int64) (*core.Account, error)
	Close(ctx context.Context, id int64) (*core.Account, error)
	Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error)
	Delete(ctx context.Context, id int64) error
}


