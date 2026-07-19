package card

import (
	"context"
	card "github.com/Suinar/Bank-proto/repository/card"
	mocks "github.com/kVinsom/Bank-repository-service/internal/mocks/repository"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCardService_GetAll_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewEmptyProto()

	expected := []core.Card{
		fixture.NewCardCore(),
		fixture.NewCardCore(),
	}

	repository.
		EXPECT().
		GetAll(gomock.Any()).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetAll(ctx, req)

	require.NoError(t, err)

	require.Len(t, result.Cards, len(expected))

	for i := range expected {
		AssertCardEqual(t, &expected[i], result.Cards[i])
	}
}

func TestCardService_GetByUser_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewUserIdRequestProto()

	expected := []core.Card{
		fixture.NewCardCore(),
		fixture.NewCardCore(),
	}

	repository.
		EXPECT().
		GetByUser(gomock.Any(), gomock.Eq(req.UserId)).
		Return(expected, nil).
		Times(1)

	result, err := sut.GetByUser(ctx, req)

	require.NoError(t, err)

	assert.Len(t, result.Cards, len(expected))

	for i := range expected {
		AssertCardEqual(t, &expected[i], result.Cards[i])
	}
}

func TestCardService_GetById_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewCardCore()

	repository.
		EXPECT().
		GetById(gomock.Any(), gomock.Eq(req.Id)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.GetById(ctx, req)

	require.NoError(t, err)

	AssertCardEqual(t, &expected, result)
}

func TestCardService_GetByNumber_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewCardNumberRequestProto()

	expected := fixture.NewCardCore()

	repository.
		EXPECT().
		GetByNumber(gomock.Any(), gomock.Eq(req.Number)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.GetByNumber(ctx, req)

	require.NoError(t, err)

	AssertCardEqual(t, &expected, result)
}

func TestCardService_Create_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	reqProto := fixture.NewCardProto()
	reqCore := fixture.NewCardCore()

	expected := fixture.NewCardCore()

	repository.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(&reqCore)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.Create(ctx, reqProto)

	require.NoError(t, err)

	AssertCardEqual(t, &expected, result)
}

func TestCardService_Blocking_Success(t *testing.T) {
	t.Parallel()

	repository, sut, ctx := NewSUT(t)

	req := fixture.NewIdRequestProto()

	expected := fixture.NewCardCore()

	repository.
		EXPECT().
		Blocking(gomock.Any(), gomock.Eq(req.Id)).
		Return(&expected, nil).
		Times(1)

	result, err := sut.Blocking(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)

	AssertCardEqual(t, &expected, result)
}

func TestCardService_Delete_Success(t *testing.T) {
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

func TestCardService_toProto_Success(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	req := fixture.NewCardCore()

	expected := fixture.NewCardProto()

	result := sut.toProto(&req)

	assert.Equal(t, expected, result)
}

func TestCardService_toCore_Success(t *testing.T) {
	t.Parallel()

	_, sut, _ := NewSUT(t)

	req := fixture.NewCardProto()

	expected := fixture.NewCardCore()

	result := sut.fromProto(req)

	assert.Equal(t, &expected, result)
}

func NewSUT(t *testing.T) (*mocks.MockICardRepository, *CardService, context.Context) {
	t.Helper()

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := mocks.NewMockICardRepository(ctrl)

	sut := NewCardService(repository)

	return repository, sut, context.Background()

}

func AssertCardEqual(t *testing.T, expected *core.Card, actual *card.Card) {
	t.Helper()

	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.UserId, actual.UserId)
	assert.Equal(t, expected.AccountId, actual.AccountId)
	assert.Equal(t, expected.Number, actual.Number)
	assert.Equal(t, int32(expected.ExpiryMonth), int32(actual.ExpiryMonth))
	assert.Equal(t, int32(expected.ExpiryYear), int32(actual.ExpiryYear))
	assert.Equal(t, int32(expected.Status), int32(actual.Status))

}
