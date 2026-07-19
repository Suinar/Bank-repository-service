package deposit

import (
	"context"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"
	connectToDB "github.com/kVinsom/Bank-repository-service/internal/test/repository"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDepositRepository_GetAll(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	reqId1 := fixture.NewDepositCoreInputId(1)
	reqId2 := fixture.NewDepositCoreInputId(2)

	expected := []core.Deposit{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertDeposit(t, &reqId1)
	db.InsertDeposit(t, &reqId2)

	deposits, err := repo.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, deposits, 2)

	require.Equal(t, expected[0], deposits[0])
	require.Equal(t, expected[1], deposits[1])
}

func TestDepositRepository_GetByUser(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	reqId1 := fixture.NewDepositCoreInputId(1)
	reqId2 := fixture.NewDepositCoreInputId(2)

	expected := []core.Deposit{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertDeposit(t, &reqId1)
	db.InsertDeposit(t, &reqId2)

	deposits, err := repo.GetByUser(ctx, user.Id)

	require.NoError(t, err)
	require.Len(t, deposits, 2)

	require.Equal(t, expected[0], deposits[0])
	require.Equal(t, expected[1], deposits[1])
}

func TestDepositRepository_GetById(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewDepositCore()

	expected := fixture.NewDepositCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertDeposit(t, &req)

	deposit, err := repo.GetById(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, deposit)
}

func TestDepositRepository_Create(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewDepositCore()

	expected := fixture.NewDepositCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)

	deposit, err := repo.Create(ctx, &req)

	require.NoError(t, err)

	require.Equal(t, &expected, deposit)
}

func TestDepositRepository_Replenish(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewDepositCore()

	expected := fixture.NewDepositCoreInputAmount(2 * fixture.TestAmount)

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertDeposit(t, &req)

	deposit, err := repo.Replenish(ctx, req.Id, fixture.TestAmount)

	require.NoError(t, err)

	require.Equal(t, &expected, deposit)
}

func TestDepositRepository_Delete(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewDepositCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertDeposit(t, &req)

	err := repo.Delete(ctx, req.Id)

	require.NoError(t, err)
}

func SetupRepositoryTest(t *testing.T) (*DepositRepository, *connectToDB.TestDB, context.Context) {
	db := connectToDB.NewTestPostgresDb(t)

	repo := NewDepositRepository(db.DB)

	ctx := context.Background()

	db.Cleanup(t)

	return repo, db, ctx
}
