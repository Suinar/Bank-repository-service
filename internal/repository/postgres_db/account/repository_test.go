package account

import (
	fixture "Bank-repository-service/internal/test/fixture"
	connectToDB "Bank-repository-service/internal/test/repository"
	core "Bank-repository-service/pkg/core"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountRepository_GetAll(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	reqId1 := fixture.NewAccountCoreInputId(1)
	reqId2 := fixture.NewAccountCoreInputId(2)

	expected := []core.Account{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &reqId1)
	db.InsertAccount(t, &reqId2)

	accounts, err := repo.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, accounts, 2)

	require.Equal(t, expected[0], accounts[0])
	require.Equal(t, expected[1], accounts[1])
}

func TestAccountRepository_GetByUser(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	reqId1 := fixture.NewAccountCoreInputId(1)
	reqId2 := fixture.NewAccountCoreInputId(2)

	expected := []core.Account{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &reqId1)
	db.InsertAccount(t, &reqId2)

	accounts, err := repo.GetByUser(ctx, user.Id)

	require.NoError(t, err)
	require.Len(t, accounts, 2)

	require.Equal(t, expected[0], accounts[0])
	require.Equal(t, expected[1], accounts[1])
}

func TestAccountRepository_GetById(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewAccountCore()

	expected := fixture.NewAccountCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &req)

	account, err := repo.GetById(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, account)
}

func TestAccountRepository_Create(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewAccountCore()

	expected := fixture.NewAccountCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)

	account, err := repo.Create(ctx, &req)

	require.NoError(t, err)

	require.Equal(t, &expected, account)
}

func TestAccountRepository_Blocking(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewAccountCore()

	expected := fixture.NewAccountCoreInputStatus(core.AccountStatusBlocked)

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &req)

	account, err := repo.Blocking(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, account)
}

func TestAccountRepository_Close(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewAccountCore()

	expected := fixture.NewAccountCoreInputStatus(core.AccountStatusClosed)

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &req)

	account, err := repo.Close(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, account)
}

func TestAccountRepository_Update(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewAccountUpdateInputCore()

	expected := fixture.NewAccountCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &expected)

	account, err := repo.Update(ctx, expected.Id, req)

	require.NoError(t, err)

	require.Equal(t, &expected, account)
}

func TestAccountRepository_Delete(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewAccountCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &req)

	err := repo.Delete(ctx, req.Id)

	require.NoError(t, err)
}

func SetupRepositoryTest(t *testing.T) (*AccountRepository, *connectToDB.TestDB, context.Context) {
	db := connectToDB.NewTestPostgresDb(t)

	repo := NewAccountRepository(db.DB)

	ctx := context.Background()

	db.Cleanup(t)

	return repo, db, ctx
}

