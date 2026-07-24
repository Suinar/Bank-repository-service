package account

import (
	mocks "github.com/kVinsom/Bank-repository-service/internal/mocks/services"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"

	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := fixture.NewAccountListProto()

	service.
		EXPECT().
		GetAll(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequestProto()

	expected := fixture.NewAccountListProto()

	service.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewAccountProto()

	service.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewAccountProto()

	expected := fixture.NewAccountProto()

	service.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Create(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Blocking_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewAccountProto()

	service.
		EXPECT().
		Blocking(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Blocking(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Close_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewAccountProto()

	service.
		EXPECT().
		Close(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Close(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Update_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateAccountRequestProto()

	expected := fixture.NewAccountProto()

	service.
		EXPECT().
		Update(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestAccountHandler_Delete_Success(t *testing.T) {
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

func NewSUT(t *testing.T) (*mocks.MockIAccountService, *AccountHandler, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	service := mocks.NewMockIAccountService(ctrl)
	sut := NewAccountHandler(service)

	return service, sut, context.Background()
}
