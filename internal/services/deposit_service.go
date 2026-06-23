package serivce

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db"
	core "Bank-repository-service/pkg/core"
)

type DepositService struct {
	repo repository.IDepositRepository
}

func NewDepositService(repo repository.IDepositRepository) *DepositService {
	return &DepositService{repo: repo}
}

func (s *DepositService) GetAll(ctx context.Context) ([]core.Deposit, error) {
	return s.repo.GetAll(ctx)
}

func (s *DepositService) GetByUser(ctx context.Context, idUser int64) ([]core.Deposit, error) {
	return s.repo.GetByUser(ctx, idUser)
}

func (s *DepositService) GetById(ctx context.Context, id int64) (*core.Deposit, error) {
	return s.repo.GetById(ctx, id)
}

func (s *DepositService) Create(ctx context.Context, input *core.Deposit) (*core.Deposit, error) {
	return s.repo.Create(ctx, input)
}

func (s *DepositService) Replenish(ctx context.Context, id int64, amount int) (*core.Deposit, error) {
	return s.repo.Replenish(ctx, id, amount)
}

func (s *DepositService) Delete(ctx context.Context, id int64) (int64, error) {
	return s.repo.Delete(ctx, id)
}
