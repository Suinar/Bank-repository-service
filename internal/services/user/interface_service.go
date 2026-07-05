package user

import (
	"Bank-repository-service/proto/repository/common"
	"Bank-repository-service/proto/repository/user"
	"context"
)

//go:generate mockgen -source=interface_service.go -destination=../../../internal/mocks/service/user.go -package=mocks

type IUserService interface {
	GetAll(ctx context.Context, req *common.Empty) (*user.UserList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*user.User, error)
	GetByEmail(ctx context.Context, req *user.EmailRequest) (*user.User, error)
	GetByPhoneNumber(ctx context.Context, req *user.PhoneNumberRequest) (*user.User, error)
	Create(ctx context.Context, req *user.User) (*user.User, error)
	ChangePassword(ctx context.Context, req *user.ChangePasswordRequest) (*common.Empty, error)
	Update(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
