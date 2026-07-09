package credit

import (
	fixture "github.com/Suinar/Bank-exhange-rate-service/internal/test/fixture"
	connectToDB "github.com/Suinar/Bank-exhange-rate-service/internal/test/repository"
	core "github.com/Suinar/Bank-exhange-rate-service/pkg/core"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreditRepository_GetAll(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	reqId1 := fixture.NewCreditCoreInputId(1)
	reqId2 := fixture.NewCreditCoreInputId(2)

	expected := []core.Credit{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertCredit(t, &reqId1)
	db.InsertCredit(t, &reqId2)

	credits, err := repo.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, credits, 2)

	require.Equal(t, expected[0], credits[0])
	require.Equal(t, expected[1], credits[1])
}

func TestCreditRepository_GetByUser(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	reqId1 := fixture.NewCreditCoreInputId(1)
	reqId2 := fixture.NewCreditCoreInputId(2)

	expected := []core.Credit{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertCredit(t, &reqId1)
	db.InsertCredit(t, &reqId2)

	credits, err := repo.GetByUser(ctx, user.Id)

	require.NoError(t, err)
	require.Len(t, credits, 2)

	require.Equal(t, expected[0], credits[0])
	require.Equal(t, expected[1], credits[1])
}

func TestCreditRepository_GetById(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewCreditCore()

	expected := fixture.NewCreditCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertCredit(t, &req)

	credit, err := repo.GetById(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, credit)
}

func TestCreditRepository_Create(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewCreditCore()

	expected := fixture.NewCreditCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)

	credit, err := repo.Create(ctx, &req)

	require.NoError(t, err)

	require.Equal(t, &expected, credit)
}

func TestCreditRepository_Repay(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewCreditCore()

	expected := fixture.NewCreditCoreInputAmount(0)

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertCredit(t, &req)

	credit, err := repo.Repay(ctx, req.Id, fixture.TestAmount)

	require.NoError(t, err)

	require.Equal(t, &expected, credit)
}

func TestCreditRepository_Delete(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	req := fixture.NewCreditCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertCredit(t, &req)

	err := repo.Delete(ctx, req.Id)

	require.NoError(t, err)
}

func SetupRepositoryTest(t *testing.T) (*CreditRepository, *connectToDB.TestDB, context.Context) {
	db := connectToDB.NewTestPostgresDb(t)

	repo := NewCreditRepository(db.DB)

	ctx := context.Background()

	db.Cleanup(t)

	return repo, db, ctx
}


