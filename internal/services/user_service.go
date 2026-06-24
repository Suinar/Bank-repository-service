package service

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db"
	core "Bank-repository-service/pkg/core"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll(ctx context.Context) ([]core.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) GetById(ctx context.Context, id int64) (*core.User, error) {
	return s.repo.GetById(ctx, id)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*core.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *UserService) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error) {
	return s.repo.GetByPhoneNumber(ctx, phoneNumber)
}

func (s *UserService) Create(ctx context.Context, input *core.User) (*core.User, error) {
	return s.repo.Create(ctx, input)
}

func (s *UserService) ChangePassword(ctx context.Context, id int64, newPassword string) error {
	return s.repo.ChangePassword(ctx, id, newPassword)
}

func (s *UserService) Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error) {
	return s.repo.Update(ctx, id, input)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
