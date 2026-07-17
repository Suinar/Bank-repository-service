package deposit

import (
	"context"
	deposit "github.com/Suinar/Bank-proto/repository/deposit"
	mocks "github.com/Suinar/Bank-repository-service/internal/mocks/repository"
	fixture "github.com/Suinar/Bank-repository-service/internal/test/fixture"
	core "github.com/Suinar/Bank-repository-service/pkg/core"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDepositService_GetAll_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := []core.Deposit{
		fixture.NewDepositCore(),
		fixture.NewDepositCore(),
	}

	repository.
		EXPECT().
		GetAll(gomock.Any()).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	require.Len(t, result.Deposits, len(expected))

	for i := range expected {
		AssertDepositEqual(t, &expected[i], result.Deposits[i])
	}
}

func TestDepositService_GetByUser_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequestProto()

	expected := []core.Deposit{
		fixture.NewDepositCore(),
		fixture.NewDepositCore(),
	}

	repository.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req.UserId)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.NoError(t, err)

	assert.Len(t, result.Deposits, len(expected))

	for i := range expected {
		AssertDepositEqual(t, &expected[i], result.Deposits[i])
	}
}

func TestDepositService_GetById_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewDepositCore()

	repository.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req.Id)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	AssertDepositEqual(t, &expected, result)
}

func TestDepositService_Create_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	reqProto := fixture.NewDepositProto()
	reqCore := fixture.NewDepositCore()

	expected := fixture.NewDepositCore()

	repository.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(&reqCore)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.Create(ctx, reqProto)

	require.NoError(t, err)

	AssertDepositEqual(t, &expected, result)
}

func TestDepositService_Delete_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewEmptyProto()

	repository.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req.Id)).
		Return(nil).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, expected, result)
}

func TestDepositService_Replenish_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewAmountRequestProto(fixture.TestAmount)

	expected := fixture.NewDepositCore()

	repository.
		EXPECT().
		Replenish(gomock.Any(), gomock.Eq(req.Id), gomock.Eq(req.Amount)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.Replenish(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)

	AssertDepositEqual(t, &expected, result)
}

func TestDepositService_toProto_Success(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	req := fixture.NewDepositCore()

	expected := fixture.NewDepositProto()

	result := sut.toProto(&req)

	assert.Equal(t, expected, result)
}

func TestDepositService_toCore_Success(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	req := fixture.NewDepositProto()

	expected := fixture.NewDepositCore()

	result := sut.fromProto(req)

	assert.Equal(t, &expected, result)
}

func NewSUT(t *testing.T) (*mocks.MockIDepositRepository, *DepositService, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := mocks.NewMockIDepositRepository(ctrl)

	sut := NewDepositService(repository)

	return repository, sut, context.Background()

}

func AssertDepositEqual(t *testing.T, expected *core.Deposit, actual *deposit.Deposit) {
	t.Helper()

	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.UserId, actual.UserId)
	assert.Equal(t, expected.CurrencyId, actual.CurrencyId)
	assert.Equal(t, expected.Amount, actual.Amount)
	assert.Equal(t, float32(expected.InterestRate), actual.InterestRate)
	assert.Equal(t, int32(expected.TermMonths), actual.TermMonths)
	assert.Equal(t, int32(expected.Status), int32(actual.Status))
}
