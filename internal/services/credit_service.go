package serivce

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db"
	core "Bank-repository-service/pkg/core"
)

type CreditService struct {
	repository repository.ICreditRepository
}

func NewCreditService(repository repository.ICreditRepository) *CreditService {
	return &CreditService{repository: repository}
}

func (s *CreditService) GetAll(ctx context.Context) ([]core.Credit, error) {
	return s.repository.GetAll(ctx)
}

func (s *CreditService) GetByUser(ctx context.Context, idUser int64) ([]core.Credit, error) {
	return s.repository.GetByUser(ctx, idUser)
}

func (s *CreditService) GetById(ctx context.Context, id int64) (*core.Credit, error) {
	return s.repository.GetById(ctx, id)
}

func (s *CreditService) Create(ctx context.Context, input *core.Credit) (*core.Credit, error) {
	return s.repository.Create(ctx, input)
}

func (s *CreditService) Repay(ctx context.Context, id int64, amount int) (*core.Credit, error) {
	return s.repository.Repay(ctx, id, amount)
}

func (s *CreditService) Delete(ctx context.Context, id int64) (int64, error) {
	return s.repository.Delete(ctx, id)
}
