package user

import (
	"context"

	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/user"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interface.go -destination=../../../internal/mocks/services/user.go -package=mocks

// IUserService defines the behavior required at this layer boundary.
type IUserService interface {
	GetAll(ctx context.Context, req *common.Empty) (*user.UserList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*user.User, error)
	GetByEmail(ctx context.Context, req *user.EmailRequest) (*user.User, error)
	GetByPhoneNumber(ctx context.Context, req *user.PhoneNumberRequest) (*user.User, error)
	Create(ctx context.Context, req *user.User) (*user.User, error)
	Update(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error)
	ChangePassword(ctx context.Context, req *user.ChangePasswordRequest) (*common.Empty, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
