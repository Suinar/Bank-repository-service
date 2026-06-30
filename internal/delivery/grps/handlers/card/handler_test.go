package card

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

func TestCardHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmpty()

	expected := fixture.NewCardList()

	service.
		EXPECT().
		GetAll(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCardHandler_GetAll_Error(t *testing.T) {
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

func TestCardHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequest()

	expected := fixture.NewCardList()

	service.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCardHandler_GetByUser_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequest()

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

func TestCardHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	expected := fixture.NewCard()

	service.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCardHandler_GetById_Error(t *testing.T) {
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

func TestCardHandler_GetByNumber_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewCardNumberRequest()

	expected := fixture.NewCard()

	service.
		EXPECT().
		GetByNumber(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByNumber(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCardHandler_GetByNumber_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewCardNumberRequest()

	service.
		EXPECT().
		GetByNumber(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetByNumber(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}
func TestCardHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewCard()

	expected := fixture.NewCard()

	service.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Create(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCardHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewCard()

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

func TestCardHandler_Blocking_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	expected := fixture.NewCard()

	service.
		EXPECT().
		Blocking(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Blocking(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCardHandler_Blocking_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	service.
		EXPECT().
		Blocking(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Blocking(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestCardHandler_Delete_Success(t *testing.T) {
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

func TestCardHandler_Delete_Error(t *testing.T) {
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

func NewSUT(t *testing.T) (*mocks.MockICardService, *CardHandler, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	service := mocks.NewMockICardService(ctrl)
	sut := NewCardHandler(service)

	return service, sut, context.Background()
}
