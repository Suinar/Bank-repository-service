package user

import (
	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/user"
	"context"
)

type IUserHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*user.UserList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*user.User, error)
	GetByEmail(ctx context.Context, req *user.EmailRequest) (*user.User, error)
	GetByPhoneNumber(ctx context.Context, req *user.PhoneNumberRequest) (*user.User, error)
	Create(ctx context.Context, req *user.User) (*user.User, error)
	Update(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}


