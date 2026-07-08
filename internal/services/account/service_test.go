package account

import (
	mocks "Bank-repository-service/internal/mocks/repository"
	fixture "Bank-repository-service/internal/test/fixture"
	errors "Bank-repository-service/pkg"
	core "Bank-repository-service/pkg/core"
	account "github.com/Suinar/Bank-proto/repository/account"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountService_GetAll_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := []core.Account{
		fixture.NewAccountCore(),
		fixture.NewAccountCore(),
	}

	repository.
		EXPECT().
		GetAll(gomock.Any()).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	require.Len(t, result.Accounts, len(expected))

	for i := range expected {
		AssertAccountEqual(t, &expected[i], result.Accounts[i])
	}
}

func TestAccountService_GetAll_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

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

func TestAccountService_GetByUser_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequestProto()

	expected := []core.Account{
		fixture.NewAccountCore(),
		fixture.NewAccountCore(),
	}

	repository.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req.UserId)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.NoError(t, err)

	assert.Len(t, result.Accounts, len(expected))

	for i := range expected {
		AssertAccountEqual(t, &expected[i], result.Accounts[i])
	}
}

func TestAccountService_GetByUser_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequestProto()

	repository.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req.UserId)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestAccountService_GetById_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewAccountCore()

	repository.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req.Id)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	AssertAccountEqual(t, &expected, result)
}

func TestAccountService_GetById_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

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

func TestAccountService_Create_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	reqProto := fixture.NewAccountProto()
	reqCore := fixture.NewAccountCore()

	expected := fixture.NewAccountCore()

	repository.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(&reqCore)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.Create(ctx, reqProto)

	require.NoError(t, err)

	AssertAccountEqual(t, &expected, result)
}

func TestAccountService_Create_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	reqProto := fixture.NewAccountProto()
	reqCore := fixture.NewAccountCore()

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

func TestAccountService_Blocking_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewAccountCore()

	repository.
		EXPECT().
		Blocking(gomock.Any(), gomock.Eq(req.Id)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.Blocking(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)

	AssertAccountEqual(t, &expected, result)
}

func TestAccountService_Blocking_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	repository.
		EXPECT().
		Blocking(gomock.Any(), gomock.Eq(req.Id)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Blocking(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestAccountService_Close_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewAccountCore()

	repository.
		EXPECT().
		Close(gomock.Any(), gomock.Eq(req.Id)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.Close(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)

	AssertAccountEqual(t, &expected, result)
}

func TestAccountService_Close_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	repository.
		EXPECT().
		Close(gomock.Any(), gomock.Eq(req.Id)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.Close(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestAccountService_Update_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateAccountRequestProto()

	expected := fixture.NewAccountCore()

	expectedInput := fixture.NewAccountUpdateInputCore()

	repository.
		EXPECT().
		Update(
			gomock.Any(),
			gomock.Eq(req.Id),
			gomock.AssignableToTypeOf(&core.AccountUpdateInput{}),
		).
		DoAndReturn(func(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error) {
			assert.Equal(t, expectedInput, input)

			return &expected, nil
		}).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)

	AssertAccountEqual(t, &expected, result)
}

func TestAccountService_Update_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateAccountRequestProto()

	expectedInput := fixture.NewAccountUpdateInputCore()

	repository.
		EXPECT().
		Update(
			gomock.Any(),
			gomock.Eq(req.Id),
			gomock.AssignableToTypeOf(&core.AccountUpdateInput{}),
		).
		DoAndReturn(func(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error) {
			assert.Equal(t, expectedInput, input)

			return nil, errors.TestError
		}).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.Error(t, err)
	require.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestAccountService_Delete_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewEmptyProto()

	repository.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req.Id)).
		Return(req.Id, nil).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, expected, result)
}

func TestAccountService_Delete_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	repository.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req.Id)).
		Return(int64(0), errors.TestError).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.Error(t, err)
	require.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestAccountService_toProto_Success(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	req := fixture.NewAccountCore()

	expected := fixture.NewAccountProto()

	result := sut.toProto(&req)

	assert.Equal(t, expected, result)
}

func TestAccountService_toCore_Success(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	req := fixture.NewAccountProto()

	expected := fixture.NewAccountCore()

	result := sut.fromProto(req)

	assert.Equal(t, &expected, result)
}

func NewSUT(t *testing.T) (*mocks.MockIAccountRepository, *AccountService, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := mocks.NewMockIAccountRepository(ctrl)

	sut := NewAccountService(repository)

	return repository, sut, context.Background()

}

func AssertAccountEqual(t *testing.T, expected *core.Account, actual *account.Account) {
	t.Helper()

	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.UserId, actual.UserId)
	assert.Equal(t, expected.CurrencyId, actual.CurrencyId)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Balance, actual.Balance)
	assert.Equal(t, int32(expected.Status), int32(actual.Status))
}

