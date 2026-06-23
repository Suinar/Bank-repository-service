package serivce

import (
	"context"

	"Bank-repository-service/internal/repository/cache"
	repository "Bank-repository-service/internal/repository/postgres_db"
	core "Bank-repository-service/pkg/core"
)

type CurrencyService struct {
	repository repository.ICurrencyRepository
	cache      cache.ICurrencyCache
}

func NewCurrencyService(repository repository.ICurrencyRepository, cache cache.ICurrencyCache) *CurrencyService {
	return &CurrencyService{
		repository: repository,
		cache:      cache,
	}
}

func (s *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {
	currencies, err := s.cache.GetAll(ctx)
	if err != nil || len(currencies) == 0 {
		return s.repository.GetAll(ctx)
	}
	return currencies, nil
}

func (s *CurrencyService) GetById(ctx context.Context, id int64) (*core.Currency, error) {
	currency, err := s.cache.GetById(ctx, id)
	if err != nil || currency == nil {
		return s.repository.GetById(ctx, id)
	}
	return currency, nil
}

func (s *CurrencyService) GetByIso(ctx context.Context, isoCode string) (*core.Currency, error) {
	currency, err := s.cache.GetByIso(ctx, isoCode)
	if err != nil || currency == nil {
		return s.repository.GetByIso(ctx, isoCode)
	}
	return currency, nil
}

func (s *CurrencyService) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	currency, err := s.cache.GetBySymbol(ctx, symbol)
	if err != nil || currency == nil {
		return s.repository.GetBySymbol(ctx, symbol)
	}
	return currency, nil
}

func (s *CurrencyService) Create(ctx context.Context, input *core.Currency) (*core.Currency, error) {
	return s.repository.Create(ctx, input)
}

func (s *CurrencyService) Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
	return s.repository.Update(ctx, id, input)
}

func (s *CurrencyService) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
