package account

import (
	"Bank-repository-service/proto/repository/account"
	"Bank-repository-service/proto/repository/common"
	"context"
)

type IAccountHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*account.AccountList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*account.AccountList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Create(ctx context.Context, req *account.Account) (*account.Account, error)
	Blocking(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Close(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Update(ctx context.Context, req *account.UpdateAccountRequest) (*account.Account, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error)
}
