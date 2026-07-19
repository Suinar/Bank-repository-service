package currency

import (
	"context"
	currency "github.com/Suinar/Bank-proto/repository/currency"
	mocksCech "github.com/kVinsom/Bank-repository-service/internal/mocks/cache"
	mocksRep "github.com/kVinsom/Bank-repository-service/internal/mocks/repository"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"testing"
	"unicode/utf8"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrencyService_GetAll_Success(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := []core.Currency{
		fixture.NewCurrencyCore(),
		fixture.NewCurrencyCore(),
	}

	cache.
		EXPECT().
		GetAll(gomock.Any()).
		Return(nil, nil)

	repository.
		EXPECT().
		GetAll(gomock.Any()).
		Return(expected, nil).
		Times(1)

	cache.
		EXPECT().
		SetAll(gomock.Any(), expected)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	require.Len(t, result.Currencies, len(expected))

	for i := range expected {
		AssertCurrencyEqual(t, &expected[i], result.Currencies[i])
	}
}

func TestCurrencyService_GetById_Success(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewCurrencyCore()

	cache.
		EXPECT().
		GetById(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	repository.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req.Id)).
		Return(&expected, nil).
		Times(1)

	cache.
		EXPECT().
		Set(gomock.Any(), gomock.Any()).
		Return(nil)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	AssertCurrencyEqual(t, &expected, result)
}

func TestCurrencyService_GetByIso_Success(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewIsoCodeRequestProto()

	expected := fixture.NewCurrencyCore()

	cache.
		EXPECT().
		GetByIso(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	repository.
		EXPECT().
		GetByIso(gomock.Any(), gomock.Eq(req.IsoCode)).
		Return(&expected, nil).
		Times(1)

	cache.
		EXPECT().
		Set(gomock.Any(), gomock.Any()).
		Return(nil)

	result, err := sut.GetByIso(ctx, req)

	require.NoError(t, err)

	AssertCurrencyEqual(t, &expected, result)
}

func TestCurrencyService_GetBySymbol_Success(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewSymbolRequestProto()
	expected := fixture.NewCurrencyCore()

	cache.
		EXPECT().
		GetBySymbol(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	symbol, _ := utf8.DecodeRuneInString(req.Symbol)
	repository.
		EXPECT().
		GetBySymbol(gomock.Any(), gomock.Eq(symbol)).
		Return(&expected, nil).
		Times(1)

	cache.
		EXPECT().
		Set(gomock.Any(), gomock.Any()).
		Return(nil)

	result, err := sut.GetBySymbol(ctx, req)

	require.NoError(t, err)

	AssertCurrencyEqual(t, &expected, result)
}

func TestCurrencyService_Create_Success(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	reqProto := fixture.NewCurrencyProto()
	reqCore := fixture.NewCurrencyCore()

	expected := fixture.NewCurrencyCore()

	repository.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(&reqCore)).
		Return(&expected, nil).
		Times(1)

	cache.
		EXPECT().
		Set(gomock.Any(), gomock.Any()).
		Return(nil)

	result, err := sut.Create(ctx, reqProto)

	require.NoError(t, err)

	AssertCurrencyEqual(t, &expected, result)
}

func TestCurrencyService_Update_Success(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateCurrencyRequestProto()

	expected := fixture.NewCurrencyCore()

	expectedInput := fixture.NewCurrencyUpdateInputCore()

	repository.
		EXPECT().
		Update(
			gomock.Any(),
			gomock.Eq(req.Id),
			gomock.AssignableToTypeOf(&core.CurrencyUpdateInput{}),
		).
		DoAndReturn(func(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
			AssertCurrencyUpdateInputEqual(t, expectedInput, input)

			return &expected, nil
		}).
		Times(1)

	cache.
		EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	result, err := sut.Update(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)

	AssertCurrencyEqual(t, &expected, result)
}

func TestCurrencyService_Delete_Success(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewEmptyProto()

	repository.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req.Id)).
		Return(nil).
		Times(1)

	cache.
		EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	result, err := sut.Delete(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, expected, result)
}

func TestCurrencyService_GetAll_CacheHit(t *testing.T) {
	t.Parallel()

	_, cache, sut, ctx := NewSUT(t)
	expected := []core.Currency{fixture.NewCurrencyCore()}
	cache.EXPECT().GetAll(gomock.Any()).Return(expected, nil).Times(1)

	result, err := sut.GetAll(ctx, fixture.NewEmptyProto())

	require.NoError(t, err)
	require.Len(t, result.Currencies, 1)
	AssertCurrencyEqual(t, &expected[0], result.Currencies[0])
}

func TestCurrencyService_GetById_CacheHit(t *testing.T) {
	t.Parallel()

	_, cache, sut, ctx := NewSUT(t)
	req := fixture.NewIdRequestProto()
	expected := fixture.NewCurrencyCore()
	cache.EXPECT().GetById(gomock.Any(), req.Id).Return(&expected, nil).Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)
	AssertCurrencyEqual(t, &expected, result)
}

func TestCurrencyService_GetByIso_CacheHit(t *testing.T) {
	t.Parallel()

	_, cache, sut, ctx := NewSUT(t)
	req := fixture.NewIsoCodeRequestProto()
	expected := fixture.NewCurrencyCore()
	cache.EXPECT().GetByIso(gomock.Any(), req.IsoCode).Return(&expected, nil).Times(1)

	result, err := sut.GetByIso(ctx, req)

	require.NoError(t, err)
	AssertCurrencyEqual(t, &expected, result)
}

func TestCurrencyService_GetBySymbol_CacheHit(t *testing.T) {
	t.Parallel()

	_, cache, sut, ctx := NewSUT(t)
	req := fixture.NewSymbolRequestProto()
	expected := fixture.NewCurrencyCore()
	symbol, _ := utf8.DecodeRuneInString(req.Symbol)
	cache.EXPECT().GetBySymbol(gomock.Any(), symbol).Return(&expected, nil).Times(1)

	result, err := sut.GetBySymbol(ctx, req)

	require.NoError(t, err)
	AssertCurrencyEqual(t, &expected, result)
}

func TestCurrencyService_toProto(t *testing.T) {
	t.Parallel()

	_, _, sut, _ := NewSUT(t)

	req := fixture.NewCurrencyCore()

	expected := fixture.NewCurrencyProto()

	result := sut.toProto(&req)

	assert.Equal(t, expected, result)
}

func TestCurrencyService_toCore(t *testing.T) {
	t.Parallel()

	_, _, sut, _ := NewSUT(t)

	req := fixture.NewCurrencyProto()

	expected := fixture.NewCurrencyCore()

	result := sut.fromProto(req)

	assert.Equal(t, &expected, result)
}

func NewSUT(t *testing.T) (*mocksRep.MockICurrencyRepository, *mocksCech.MockICurrencyCache, *CurrencyService, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := mocksRep.NewMockICurrencyRepository(ctrl)
	cache := mocksCech.NewMockICurrencyCache(ctrl)

	sut := NewCurrencyService(repository, cache)

	return repository, cache, sut, context.Background()

}

func AssertCurrencyEqual(t *testing.T, expected *core.Currency, actual *currency.Currency) {
	t.Helper()

	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, string(expected.Symbol), actual.Symbol)
	assert.Equal(t, expected.IsoCode, actual.IsoCode)
	assert.Equal(t, int32(expected.MinorUnits), actual.MinorUnits)
}

func AssertCurrencyUpdateInputEqual(t *testing.T, expected *core.CurrencyUpdateInput, actual *core.CurrencyUpdateInput) {
	t.Helper()

	assert.Equal(t, *expected.Name, *actual.Name)
	assert.Equal(t, *expected.IsoCode, *actual.IsoCode)
	assert.Equal(t, *expected.Symbol, *actual.Symbol)
	assert.Equal(t, *expected.MinorUnits, *actual.MinorUnits)
}
