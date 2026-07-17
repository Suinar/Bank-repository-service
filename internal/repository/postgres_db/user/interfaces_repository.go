package user

import (
	"context"
	"github.com/Suinar/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interfaces_repository.go -destination=../../../mocks/repository/user.go -package=mocks

// IUserRepository defines the behavior required at this layer boundary.
type IUserRepository interface {
	GetAll(ctx context.Context) ([]core.User, error)
	GetById(ctx context.Context, id int64) (*core.User, error)
	GetByEmail(ctx context.Context, email string) (*core.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error)
	Create(ctx context.Context, input *core.User) (*core.User, error)
	Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error)
	Delete(ctx context.Context, id int64) error
}
