package serivce

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db"
	core "Bank-repository-service/pkg/core"
)

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repository repository.IAccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) GetAll(ctx context.Context) ([]core.Account, error) {
	return s.repository.GetAll(ctx)
}

func (s *AccountService) GetByUser(ctx context.Context, userId int64) ([]core.Account, error) {
	return s.repository.GetByUser(ctx, userId)
}

func (s *AccountService) GetById(ctx context.Context, id int64) (*core.Account, error) {
	return s.repository.GetById(ctx, id)
}

func (s *AccountService) Create(ctx context.Context, input *core.Account) (*core.Account, error) {
	return s.repository.Create(ctx, input)
}

func (s *AccountService) Blocking(ctx context.Context, id int64) (core.Account, error) {
	return s.repository.Blocking(ctx, id)
}

func (s *AccountService) Close(ctx context.Context, id int64) (core.Account, error) {
	return s.repository.Close(ctx, id)
}

func (s *AccountService) Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error) {
	return s.repository.Update(ctx, id, input)
}

func (s *AccountService) Delete(ctx context.Context, id int64) (int64, error) {
	return s.repository.Delete(ctx, id)
}
