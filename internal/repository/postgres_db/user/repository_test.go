package user

import (
	"context"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"
	connectToDB "github.com/kVinsom/Bank-repository-service/internal/test/repository"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserRepository_GetAll(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	reqId1 := fixture.NewUserCoreInputIdAndEmailAndPhoneNumber(1, "Mykola@gmail.com", "+38000000000")
	reqId2 := fixture.NewUserCoreInputIdAndEmailAndPhoneNumber(2, "Mykola2@gmail.com", "+38000000001")

	expected := []core.User{
		reqId1,
		reqId2,
	}

	db.InsertUser(t, &reqId1)
	db.InsertUser(t, &reqId2)

	Users, err := repo.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, Users, 2)

	require.Equal(t, expected[0], Users[0])
	require.Equal(t, expected[1], Users[1])
}

func TestUserRepository_GetById(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewUserCore()

	expected := fixture.NewUserCore()

	db.InsertUser(t, &req)

	user, err := repo.GetById(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, user)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewUserCore()

	expected := fixture.NewUserCore()

	db.InsertUser(t, &req)

	user, err := repo.GetByEmail(ctx, req.Email)

	require.NoError(t, err)

	require.Equal(t, &expected, user)
}

func TestUserRepository_GetByPhoneNumber(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewUserCore()

	expected := fixture.NewUserCore()

	db.InsertUser(t, &req)

	user, err := repo.GetByPhoneNumber(ctx, req.PhoneNumber)

	require.NoError(t, err)

	require.Equal(t, &expected, user)
}

func TestUserRepository_Create(t *testing.T) {
	repo, _, ctx := SetupRepositoryTest(t)

	req := fixture.NewUserCore()

	expected := fixture.NewUserCore()

	user, err := repo.Create(ctx, &req)

	require.NoError(t, err)

	require.Equal(t, &expected, user)
}

func TestUserRepository_Update(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewUserUpdateInputCore()

	expected := fixture.NewUserCore()

	db.InsertUser(t, &expected)

	user, err := repo.Update(ctx, expected.Id, req)

	require.NoError(t, err)

	require.Equal(t, &expected, user)
}

func TestUserRepository_Delete(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewUserCore()

	db.InsertUser(t, &req)

	err := repo.Delete(ctx, req.Id)

	require.NoError(t, err)
}

func SetupRepositoryTest(t *testing.T) (*UserRepository, *connectToDB.TestDB, context.Context) {
	db := connectToDB.NewTestPostgresDb(t)

	repo := NewUserRepository(db.DB)

	ctx := context.Background()

	db.Cleanup(t)

	return repo, db, ctx
}
