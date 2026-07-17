package currency

import (
	"context"
	fixture "github.com/Suinar/Bank-repository-service/internal/test/fixture"
	errors "github.com/Suinar/Bank-repository-service/pkg"
	core "github.com/Suinar/Bank-repository-service/pkg/core"
	"testing"
	"unicode/utf8"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrencyService_GetAll_Error(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	cache.
		EXPECT().
		GetAll(gomock.Any()).
		Return(nil, nil)

	repository.
		EXPECT().
		GetAll(gomock.Any()).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyService_GetById_Error(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	cache.
		EXPECT().
		GetById(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	repository.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req.Id)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyService_GetByIso_Error(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewIsoCodeRequestProto()

	cache.
		EXPECT().
		GetByIso(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	repository.
		EXPECT().
		GetByIso(gomock.Any(), gomock.Eq(req.IsoCode)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetByIso(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyService_GetBySymbol_Error(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)

	req := fixture.NewSymbolRequestProto()

	cache.
		EXPECT().
		GetBySymbol(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	symbol, _ := utf8.DecodeRuneInString(req.Symbol)
	repository.
		EXPECT().
		GetBySymbol(gomock.Any(), gomock.Eq(symbol)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetBySymbol(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyService_Create_Error(t *testing.T) {
	t.Parallel()

	repository, _, sut, ctx := NewSUT(t)

	reqProto := fixture.NewCurrencyProto()
	reqCore := fixture.NewCurrencyCore()

	repository.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(&reqCore)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Create(ctx, reqProto)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyService_Update_Error(t *testing.T) {
	t.Parallel()

	repository, _, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateCurrencyRequestProto()

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

			return nil, errors.TestError
		}).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.Error(t, err)
	require.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyService_Delete_Error(t *testing.T) {
	t.Parallel()

	repository, _, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	repository.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req.Id)).
		Return(errors.TestError).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.Error(t, err)
	require.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyService_Create_CacheError(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)
	req := fixture.NewCurrencyProto()
	created := fixture.NewCurrencyCore()
	repository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&created, nil).Times(1)
	cache.EXPECT().Set(gomock.Any(), &created).Return(errors.TestError).Times(1)

	result, err := sut.Create(ctx, req)

	require.ErrorIs(t, err, errors.TestError)
	assert.Nil(t, result)
}

func TestCurrencyService_Update_EmptySymbol(t *testing.T) {
	t.Parallel()

	_, _, sut, ctx := NewSUT(t)
	req := fixture.NewUpdateCurrencyRequestProto()
	empty := ""
	req.Input.Symbol = &empty

	result, err := sut.Update(ctx, req)

	require.ErrorIs(t, err, errors.BadRequest)
	assert.Nil(t, result)
}

func TestCurrencyService_Update_CacheError(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)
	req := fixture.NewUpdateCurrencyRequestProto()
	updated := fixture.NewCurrencyCore()
	repository.EXPECT().Update(gomock.Any(), req.Id, gomock.Any()).Return(&updated, nil).Times(1)
	cache.EXPECT().Update(gomock.Any(), &updated).Return(errors.TestError).Times(1)

	result, err := sut.Update(ctx, req)

	require.ErrorIs(t, err, errors.TestError)
	assert.Nil(t, result)
}

func TestCurrencyService_Delete_CacheError(t *testing.T) {
	t.Parallel()

	repository, cache, sut, ctx := NewSUT(t)
	req := fixture.NewIdRequestProto()
	repository.EXPECT().Delete(gomock.Any(), req.Id).Return(nil).Times(1)
	cache.EXPECT().Delete(gomock.Any(), req.Id).Return(errors.TestError).Times(1)

	result, err := sut.Delete(ctx, req)

	require.ErrorIs(t, err, errors.TestError)
	assert.Nil(t, result)
}

func TestCurrencyService_Mappers_NilAndEmptySymbol(t *testing.T) {
	t.Parallel()

	_, _, sut, _ := NewSUT(t)
	assert.Nil(t, sut.toProto(nil))
	assert.Nil(t, sut.fromProto(nil))

	input := fixture.NewCurrencyProto()
	input.Symbol = ""
	result := sut.fromProto(input)
	require.NotNil(t, result)
	assert.Equal(t, rune(0), result.Symbol)
}
