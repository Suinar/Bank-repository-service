package service

import (
	"Bank-repository-service/proto/repository/common"
	"Bank-repository-service/proto/repository/currency"
	"context"

	"Bank-repository-service/internal/repository/cache"
	repository "Bank-repository-service/internal/repository/postgres_db"
	"Bank-repository-service/pkg/core"
)

type CurrencyService struct {
	repository repository.ICurrencyRepository
	cache      cache.ICurrencyCache
}

func NewCurrencyService(repository repository.ICurrencyRepository,
	cache cache.ICurrencyCache,
) *CurrencyService {
	return &CurrencyService{
		repository: repository,
		cache:      cache,
	}
}

func (s *CurrencyService) GetAll(ctx context.Context) ([]core.Currency, error) {
	currencies, err := s.cache.GetAll(ctx)
	if err != nil || len(currencies) == 0 {
		currencies, err = s.repository.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		s.cache.SetAll(ctx, currencies)
	}

	return currencies, nil
}

func (s *CurrencyService) GetById(ctx context.Context, id int64) (*currency.Currency, error) {
	currency, err := s.cache.GetById(ctx, id)
	if err != nil || currency == nil {
		currency, err = s.repository.GetById(ctx, id)
		if err != nil {
			return nil, err
		}

		s.cache.Set(ctx, currency)
	}

	return s.toProto(currency), nil
}

func (s *CurrencyService) GetByIso(ctx context.Context, isoCode string) (*currency.Currency, error) {
	currency, err := s.cache.GetByIso(ctx, isoCode)
	if err != nil || currency == nil {
		currency, err = s.repository.GetByIso(ctx, isoCode)
		if err != nil {
			return nil, err
		}

		s.cache.Set(ctx, currency)
	}

	return s.toProto(currency), nil
}

func (s *CurrencyService) GetBySymbol(ctx context.Context, symbol rune) (*currency.Currency, error) {
	currency, err := s.cache.GetBySymbol(ctx, symbol)
	if err != nil || currency == nil {
		currency, err = s.repository.GetBySymbol(ctx, symbol)
		if err != nil {
			return nil, err
		}

		s.cache.Set(ctx, currency)
	}

	return s.toProto(currency), nil
}

func (s *CurrencyService) Create(ctx context.Context, input *core.Currency) (*currency.Currency, error) {
	currency, err := s.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	s.cache.Set(ctx, currency)

	return s.toProto(currency), nil
}

func (s *CurrencyService) Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*currency.Currency, error) {
	currency, err := s.repository.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}

	s.cache.Update(ctx, currency)

	return s.toProto(currency), nil
}

func (s *CurrencyService) Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error) {
	err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	s.cache.Delete(ctx, req.Id)

	return &common.DeleteResponse{
		EntityId: 0,
	}, nil
}

func (s *CurrencyService) toProto(input *core.Currency) *currency.Currency {
	if input == nil {
		return nil
	}

	return &currency.Currency{
		Id:         input.Id,
		Name:       input.Name,
		Symbol:     string(input.Symbol),
		IsoCode:    input.IsoCode,
		MinorUnits: int32(input.MinorUnits),
	}
}

func (s *CurrencyService) fromProto(input *currency.Currency) *core.Currency {
	if input == nil {
		return nil
	}

	var symbol rune
	if len([]rune(input.Symbol)) > 0 {
		symbol = []rune(input.Symbol)[0]
	}

	return &core.Currency{
		Id:         input.Id,
		Name:       input.Name,
		Symbol:     symbol,
		IsoCode:    input.IsoCode,
		MinorUnits: int8(input.MinorUnits),
	}
}
