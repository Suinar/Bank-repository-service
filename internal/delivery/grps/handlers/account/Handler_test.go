package account

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

func TestAccountHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmpty()

	expected := fixture.NewAccountList(
		fixture.TestName,
		fixture.TestName,
	)

	service.
		EXPECT().
		GetAll(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_GetAll_Error(t *testing.T) {
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

func TestAccountHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequest()

	expected := fixture.NewAccountList(
		fixture.TestName,
		fixture.TestName,
	)

	service.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_GetByUser_Error(t *testing.T) {
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

func TestAccountHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	expected := fixture.NewAccount(fixture.TestName)

	service.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_GetById_Error(t *testing.T) {
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

func TestAccountHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewAccount(fixture.TestName)

	expected := fixture.NewAccount(fixture.TestName)

	service.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Create(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Create_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewAccount(fixture.TestName)

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

func TestAccountHandler_Blocking_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	expected := fixture.NewAccount(fixture.TestName)

	service.
		EXPECT().
		Blocking(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Blocking(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Blocking_Error(t *testing.T) {
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

func TestAccountHandler_Close_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	expected := fixture.NewAccount(fixture.TestName)

	service.
		EXPECT().
		Close(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Close(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Close_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequest()

	service.
		EXPECT().
		Close(gomock.Any(), gomock.Eq(req)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Close(ctx, req)

	require.Error(t, err)

	assert.Nil(t, result)

	assert.ErrorIs(t, err, errors.TestError)
}

func TestAccountHandler_Update_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateAccountRequest(fixture.TestUpdateName)

	expected := fixture.NewAccount(fixture.TestUpdateName)

	service.
		EXPECT().
		Update(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Update_Error(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateAccountRequest(fixture.TestUpdateName)

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

func TestAccountHandler_Delete_Success(t *testing.T) {
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

func TestAccountHandler_Delete_Error(t *testing.T) {
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

func NewSUT(t *testing.T) (*mocks.MockIAccountService, *AccountHandler, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	service := mocks.NewMockIAccountService(ctrl)
	sut := NewAccountHandler(service)

	return service, sut, context.Background()
}
