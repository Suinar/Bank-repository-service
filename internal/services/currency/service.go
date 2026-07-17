package currency

import (
	"context"
	common "github.com/Suinar/Bank-proto/repository/common"
	currency "github.com/Suinar/Bank-proto/repository/currency"
	cache "github.com/Suinar/Bank-repository-service/internal/repository/cache/currency"
	repository "github.com/Suinar/Bank-repository-service/internal/repository/postgres_db/currency"
	errors "github.com/Suinar/Bank-repository-service/pkg"
	"unicode/utf8"

	"github.com/Suinar/Bank-repository-service/pkg/core"
)

// CurrencyService coordinates the application use cases for its domain.
type CurrencyService struct {
	repository repository.ICurrencyRepository
	cache      cache.ICurrencyCache
}

// NewCurrencyService creates a ready-to-use currency service.
func NewCurrencyService(repository repository.ICurrencyRepository, cache cache.ICurrencyCache) *CurrencyService {
	return &CurrencyService{
		repository: repository,
		cache:      cache,
	}
}

// GetAll returns all records available through CurrencyService.
func (s *CurrencyService) GetAll(ctx context.Context, req *common.Empty) (*currency.CurrencyList, error) {
	currencies, err := s.cache.GetAll(ctx)
	if err != nil || len(currencies) == 0 {
		currencies, err = s.repository.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		s.cache.SetAll(ctx, currencies)
	}

	res := &currency.CurrencyList{
		Currencies: make([]*currency.Currency, len(currencies)),
	}

	for i, curr := range currencies {
		res.Currencies[i] = s.toProto(&curr)
	}

	return res, nil
}

// GetById returns records matching the requested id lookup.
func (s *CurrencyService) GetById(ctx context.Context, req *common.IdRequest) (*currency.Currency, error) {
	currency, err := s.cache.GetById(ctx, req.Id)
	if err != nil || currency == nil {
		currency, err = s.repository.GetById(ctx, req.Id)
		if err != nil {
			return nil, err
		}

		s.cache.Set(ctx, currency)
	}

	return s.toProto(currency), nil
}

// GetByIso returns records matching the requested iso lookup.
func (s *CurrencyService) GetByIso(ctx context.Context, req *currency.IsoCodeRequest) (*currency.Currency, error) {
	currency, err := s.cache.GetByIso(ctx, req.IsoCode)
	if err != nil || currency == nil {
		currency, err = s.repository.GetByIso(ctx, req.IsoCode)
		if err != nil {
			return nil, err
		}

		s.cache.Set(ctx, currency)
	}

	return s.toProto(currency), nil
}

// GetBySymbol returns records matching the requested symbol lookup.
func (s *CurrencyService) GetBySymbol(ctx context.Context, req *currency.SymbolRequest) (*currency.Currency, error) {
	symbol, _ := utf8.DecodeRuneInString(req.Symbol)

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

// Create persists a new record through CurrencyService.
func (s *CurrencyService) Create(ctx context.Context, input *currency.Currency) (*currency.Currency, error) {
	currency, err := s.repository.Create(ctx, s.fromProto(input))
	if err != nil {
		return nil, err
	}

	err = s.cache.Set(ctx, currency)
	if err != nil {
		return nil, err
	}

	return s.toProto(currency), nil
}

// Update applies the requested changes through CurrencyService.
func (s *CurrencyService) Update(ctx context.Context, req *currency.UpdateCurrencyRequest) (*currency.Currency, error) {
	symbol, size := utf8.DecodeRuneInString(*req.Input.Symbol)
	if size == 0 {
		return nil, errors.BadRequest
	}

	mirrorUnits := int8(*req.Input.MinorUnits)

	currency, err := s.repository.Update(ctx, req.Id, &core.CurrencyUpdateInput{
		Name:       req.Input.Name,
		Symbol:     &symbol,
		IsoCode:    req.Input.IsoCode,
		MinorUnits: &mirrorUnits,
	})
	if err != nil {
		return nil, err
	}

	err = s.cache.Update(ctx, currency)
	if err != nil {
		return nil, err
	}

	return s.toProto(currency), nil
}

// Delete removes the requested record through CurrencyService.
func (s *CurrencyService) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	err = s.cache.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

// toProto maps the domain model to its protobuf representation.
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

// fromProto maps a protobuf message to the domain model.
func (s *CurrencyService) fromProto(input *currency.Currency) *core.Currency {
	if input == nil {
		return nil
	}

	var symbol rune
	if input.Symbol != "" {
		symbol, _ = utf8.DecodeRuneInString(input.Symbol)
	}

	return &core.Currency{
		Id:         input.Id,
		Name:       input.Name,
		Symbol:     symbol,
		IsoCode:    input.IsoCode,
		MinorUnits: int8(input.MinorUnits),
	}
}
