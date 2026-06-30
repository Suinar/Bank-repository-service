package currency

import (
	mocks "Bank-repository-service/internal/mocks/service"
	fixture "Bank-repository-service/internal/test/fixture"
	errors "Bank-repository-service/pkg"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrencyHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmpty()

	expected := fixture.NewCurrencyList(
		fixture.TestIsoCode, fixture.TestIsoCode,
		fixture.TestCurrencyName, fixture.TestCurrencyName,
		fixture.TestSymbol, fixture.TestSymbol,
		fixture.TestMinorUnits, fixture.TestMinorUnits)

	service.
		EXPECT().
		GetAll(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_GetAll_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmpty()

	service.
		EXPECT().
		GetAll(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	expected := fixture.NewCurrency(fixture.TestIsoCode, fixture.TestCurrencyName,
		fixture.TestSymbol, fixture.TestMinorUnits)

	service.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_GetById_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	service.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyHandler_GetByIso_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIsoCodeRequest(fixture.TestIsoCode)

	expected := fixture.NewCurrency(fixture.TestIsoCode, fixture.TestCurrencyName,
		fixture.TestSymbol, fixture.TestMinorUnits)

	service.
		EXPECT().
		GetByIso(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByIso(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_GetByIso_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIsoCodeRequest(fixture.TestIsoCode)

	service.
		EXPECT().
		GetByIso(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetByIso(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewCurrency(fixture.TestIsoCode, fixture.TestCurrencyName,
		fixture.TestSymbol, fixture.TestMinorUnits)

	expected := fixture.NewCurrency(fixture.TestIsoCode, fixture.TestCurrencyName,
		fixture.TestSymbol, fixture.TestMinorUnits)

	service.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Create(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewCurrency(fixture.TestIsoCode, fixture.TestCurrencyName,
		fixture.TestSymbol, fixture.TestMinorUnits)

	service.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Create(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyHandler_Update_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateCurrencyRequest()

	expected := fixture.NewCurrency(fixture.TestUpdateIsoCode, fixture.TestUpdateCurrencyName,
		fixture.TestUpdateSymbol, fixture.TestUpdateMinorUnits)

	service.
		EXPECT().
		Update(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_Update_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateCurrencyRequest()

	service.
		EXPECT().
		Update(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestCurrencyHandler_Delete_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	expected := fixture.NewDeleteResponse()

	service.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCurrencyHandler_Delete_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	service.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func NewSUT(t *testing.T) (*mocks.MockICurrencyService, *CurrencyHandler, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	service := mocks.NewMockICurrencyService(ctrl)
	sut := NewCurrencyHandler(service)

	return service, sut, context.Background()
}
