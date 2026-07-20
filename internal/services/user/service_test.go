package user

import (
	"context"
	user "github.com/Suinar/Bank-proto/repository/user"
	mocks "github.com/kVinsom/Bank-repository-service/internal/mocks/repository"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_GetAll_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := []core.User{
		fixture.NewUserCore(),
		fixture.NewUserCore(),
	}

	repository.
		EXPECT().
		GetAll(gomock.Any()).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	require.Len(t, result.Users, len(expected))

	for i := range expected {
		AssertUserEqual(t, &expected[i], result.Users[i])
	}
}

func TestUserService_GetById_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewUserCore()

	repository.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req.Id)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	AssertUserEqual(t, &expected, result)
}

func TestUserService_GetByEmail_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewEmailRequestProto()

	expected := fixture.NewUserCore()

	repository.
		EXPECT().
		GetByEmail(gomock.Any(), gomock.Eq(req.Email)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.GetByEmail(ctx, req)

	require.NoError(t, err)

	AssertUserEqual(t, &expected, result)
}

func TestUserService_GetByPhoneNumber_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewPhoneNumberRequestProto()

	expected := fixture.NewUserCore()

	repository.
		EXPECT().
		GetByPhoneNumber(gomock.Any(), gomock.Eq(req.PhoneNumber)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.GetByPhoneNumber(ctx, req)

	require.NoError(t, err)

	AssertUserEqual(t, &expected, result)
}

func TestUserService_Create_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	reqProto := fixture.NewUserProto()
	reqCore := fixture.NewUserCore()

	expected := fixture.NewUserCore()

	repository.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(&reqCore)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.Create(ctx, reqProto)

	require.NoError(t, err)

	AssertUserEqual(t, &expected, result)
}

func TestUserService_Update_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateUserRequestProto()

	expected := fixture.NewUserCore()

	expectedInput := fixture.NewUserUpdateInputCore()

	repository.
		EXPECT().
		Update(
			gomock.Any(),
			gomock.Eq(req.Id),
			gomock.AssignableToTypeOf(&core.UserUpdateInput{}),
		).
		DoAndReturn(func(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error) {
			assert.Equal(t, expectedInput, input)

			return &expected, nil
		}).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)

	AssertUserEqual(t, &expected, result)
}

func TestUserService_Delete_Success(t *testing.T) {
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

func TestUserService_ChangePassword_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)
	req := fixture.NewChangePasswordRequestProto()

	repository.EXPECT().
		ChangePassword(gomock.Any(), req.Id, req.NewPassword).
		Return(nil).
		Times(1)

	result, err := sut.ChangePassword(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestUserService_toProto_Success(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	req := fixture.NewUserCore()

	expected := fixture.NewUserProto()

	result := sut.toProto(&req)

	assert.Equal(t, expected, result)
}

func TestUserService_toCore_Success(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	req := fixture.NewUserProto()

	expected := fixture.NewUserCore()

	result := sut.fromProto(req)

	assert.Equal(t, &expected, result)
}

func NewSUT(t *testing.T) (*mocks.MockIUserRepository, *UserService, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := mocks.NewMockIUserRepository(ctrl)

	sut := NewUserService(repository)

	return repository, sut, context.Background()

}

func AssertUserEqual(t *testing.T, expected *core.User, actual *user.User) {
	t.Helper()

	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.FirstName, actual.FirstName)
	assert.Equal(t, *expected.MiddleName, *actual.MiddleName)
	assert.Equal(t, expected.LastName, actual.LastName)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.PhoneNumber, actual.PhoneNumber)
	assert.Equal(t, expected.PasswordHash, actual.PasswordHash)
}
