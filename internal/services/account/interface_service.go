package account

import (
	"context"

	"github.com/Suinar/Bank-proto/repository/account"
	"github.com/Suinar/Bank-proto/repository/common"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface_service.go -destination=../../../internal/mocks/services/account.go -package=mocks

// IAccountService defines the behavior required at this layer boundary.
type IAccountService interface {
	GetAll(ctx context.Context, req *common.Empty) (*account.AccountList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*account.AccountList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Create(ctx context.Context, req *account.Account) (*account.Account, error)
	Blocking(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Close(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Update(ctx context.Context, req *account.UpdateAccountRequest) (*account.Account, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
