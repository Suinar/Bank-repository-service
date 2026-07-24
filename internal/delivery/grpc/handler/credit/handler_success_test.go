package credit

import (
	"context"
	mocks "github.com/kVinsom/Bank-repository-service/internal/mocks/services"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreditHandler_GetAll_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := fixture.NewCreditListProto()

	service.
		EXPECT().
		GetAll(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCreditHandler_GetByUser_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequestProto()

	expected := fixture.NewCreditListProto()

	service.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCreditHandler_GetById_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewCreditProto()

	service.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCreditHandler_Create_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewCreditProto()

	expected := fixture.NewCreditProto()

	service.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Create(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCreditHandler_Repay_Success(t *testing.T) {
	t.Parallel()

	service, sut, ctx := NewSUT(t)

	req := fixture.NewAmountRequestProto(fixture.TestAmount)

	expected := fixture.NewCreditProto()

	service.
		EXPECT().
		Repay(gomock.Any(), gomock.Eq(req)).
		Return(expected, nil).
		Times(1)

	result, err := sut.Repay(ctx, req)

	require.NoError(t, err)

	assert.Equal(t, expected, result)
}

func TestCreditHandler_Delete_Success(t *testing.T) {
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

func NewSUT(t *testing.T) (*mocks.MockICreditService, *CreditHandler, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	service := mocks.NewMockICreditService(ctrl)
	sut := NewCreditHandler(service)

	return service, sut, context.Background()
}
