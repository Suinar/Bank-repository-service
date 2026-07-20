package user

import (
	"context"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"
	errors "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_GetAll_Error(t *testing.T) {
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

func TestUserService_GetById_Error(t *testing.T) {
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

func TestUserService_GetByEmail_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewEmailRequestProto()

	repository.
		EXPECT().
		GetByEmail(gomock.Any(), gomock.Eq(req.Email)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetByEmail(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestUserService_GetByPhoneNumber_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewPhoneNumberRequestProto()

	repository.
		EXPECT().
		GetByPhoneNumber(gomock.Any(), gomock.Eq(req.PhoneNumber)).
		Return(nil, errors.TestError).
		Times(1)

	result, err := sut.GetByPhoneNumber(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestUserService_Create_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	reqProto := fixture.NewUserProto()
	reqCore := fixture.NewUserCore()

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

func TestUserService_Update_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewUpdateUserRequestProto()

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

			return nil, errors.TestError
		}).
		Times(1)

	result, err := sut.Update(ctx, req)

	require.Error(t, err)
	require.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestUserService_Delete_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	repository.
		EXPECT().
		Delete(gomock.Any(), gomock.Eq(req.Id)).
		Return(errors.TestError).
		Times(1)

	result, err := sut.Delete(ctx, req)

	require.Error(t, err)
	require.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestUserService_ChangePassword_Error(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)
	req := fixture.NewChangePasswordRequestProto()

	repository.EXPECT().
		ChangePassword(gomock.Any(), req.Id, req.NewPassword).
		Return(errors.TestError).
		Times(1)

	result, err := sut.ChangePassword(ctx, req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.TestError)
}

func TestUserService_Mappers_Nil(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	assert.Nil(t, sut.toProto(nil))
	assert.Nil(t, sut.fromProto(nil))
}
