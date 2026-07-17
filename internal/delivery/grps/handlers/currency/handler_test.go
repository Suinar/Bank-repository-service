package currency

import (
	"context"
	mocks "github.com/Suinar/Bank-repository-service/internal/mocks/services"
	"github.com/Suinar/Bank-repository-service/internal/test/fixture"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrencyHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := fixture.NewCurrencyListProto()

	service.
		EXPECT().
		GetAll(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewCurrencyProto()

	service.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_GetByIso_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIsoCodeRequestProto()

	expected := fixture.NewCurrencyProto()

	service.
		EXPECT().
		GetByIso(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByIso(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_GetBySymbol_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)
	req := fixture.NewSymbolRequestProto()
	expected := fixture.NewCurrencyProto()

	service.EXPECT().GetBySymbol(gomock.Any(), gomock.Eq(req)).Return(expected, nil).Times(1)

	result, err := sut.GetBySymbol(ctx, req)

	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewCurrencyProto()

	expected := fixture.NewCurrencyProto()

	service.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Create(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_Update_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateCurrencyRequestProto()

	expected := fixture.NewCurrencyProto()

	service.
		EXPECT().
		Update(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_Delete_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewEmptyProto()

	service.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func NewSUT(t *testing.T) (*mocks.MockICurrencyService, *CurrencyHandler, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	service := mocks.NewMockICurrencyService(ctrl)
	sut := NewCurrencyHandler(service)

	return service, sut, context.Background()
}
