package deposit

import (
	mocks "Bank-repository-service/internal/mocks/service"
	"Bank-repository-service/internal/test/fixture"
	errors "Bank-repository-service/pkg"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDepositHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := fixture.NewDepositListProto()

	service.
		EXPECT().
		GetAll(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestDepositHandler_GetAll_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

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

func TestDepositHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequestProto()

	expected := fixture.NewDepositListProto()

	service.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestDepositHandler_GetByUser_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequestProto()

	service.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestDepositHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewDepositProto()

	service.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestDepositHandler_GetById_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

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

func TestDepositHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewDepositProto()

	expected := fixture.NewDepositProto()

	service.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Create(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestDepositHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewDepositProto()

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

func TestDepositHandler_Replenish_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewAmountRequestProto(fixture.TestAmount)

	expected := fixture.NewDepositProto()

	service.
		EXPECT().
		Replenish(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Replenish(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestDepositHandler_Replenish_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewAmountRequestProto(fixture.TestAmount)

	service.
		EXPECT().
		Replenish(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Replenish(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestDepositHandler_Delete_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewDeleteResponseProto()

	service.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestDepositHandler_Delete_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

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

func NewSUT(t *testing.T) (*mocks.MockIDepositService, *DepositHandler, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	service := mocks.NewMockIDepositService(ctrl)
	sut := NewDepositHandler(service)

	return service, sut, context.Background()
}
