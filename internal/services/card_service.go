package serivce

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db"
	core "Bank-repository-service/pkg/core"
)

type CardService struct {
	repository repository.ICardRepository
}

func NewCardService(repository repository.ICardRepository) *CardService {
	return &CardService{repository: repository}
}

func (s *CardService) GetAll(ctx context.Context) ([]core.Card, error) {
	return s.repository.GetAll(ctx)
}

func (s *CardService) GetByUser(ctx context.Context, idUser int64) ([]core.Card, error) {
	return s.repository.GetByUser(ctx, idUser)
}

func (s *CardService) GetById(ctx context.Context, id int64) (*core.Card, error) {
	return s.repository.GetById(ctx, id)
}

func (s *CardService) GetByNumber(ctx context.Context, number string) (*core.Card, error) {
	return s.repository.GetByNumber(ctx, number)
}

func (s *CardService) Blocking(ctx context.Context, id int64) (core.Card, error) {
	return s.repository.Blocking(ctx, id)
}

func (s *CardService) Create(ctx context.Context, input *core.CardCreateInput) (*core.Card, error) {
	return s.repository.Create(ctx, input)
}

func (s *CardService) Delete(ctx context.Context, id int64) (int64, error) {
	return s.repository.Delete(ctx, id)
}
