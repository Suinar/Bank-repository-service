package user

import (
	"Bank-repository-service/pkg/core"
	"context"
)

//go:generate mockgen -source=interfaces_repository.go -destination=../../../mocks/repository/user.go -package=mocks

type IUserRepository interface {
	GetAll(ctx context.Context) ([]core.User, error)
	GetById(ctx context.Context, id int64) (*core.User, error)
	GetByEmail(ctx context.Context, email string) (*core.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error)
	Create(ctx context.Context, input *core.User) (*core.User, error)
	ChangePassword(ctx context.Context, id int64, newPassword string) error
	Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error)
	Delete(ctx context.Context, id int64) error
}
