package account

import (
	"context"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interfaces.go -destination=../../../mocks/repository/account.go -package=mocks

// IAccountRepository defines the behavior required at this layer boundary.
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
