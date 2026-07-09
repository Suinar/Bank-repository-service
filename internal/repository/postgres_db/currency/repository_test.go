package currency

import (
	fixture "github.com/Suinar/Bank-repository-service/internal/test/fixture"
	connectToDB "github.com/Suinar/Bank-repository-service/internal/test/repository"
	core "github.com/Suinar/Bank-repository-service/pkg/core"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurrencyRepository_GetAll(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	reqId1 := fixture.NewCurrencyCoreInputIdAndIsoAndName(1, "USD", "Us dollar")
	reqId2 := fixture.NewCurrencyCoreInputIdAndIsoAndName(2, "EUR", "Euro")

	expected := []core.Currency{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &reqId1)
	db.InsertCurrency(t, &reqId2)

	currencies, err := repo.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, currencies, 2)

	require.Equal(t, expected[0], currencies[0])
	require.Equal(t, expected[1], currencies[1])
}

func TestCurrencyRepository_GetById(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewCurrencyCore()

	expected := fixture.NewCurrencyCore()

	db.InsertCurrency(t, &req)

	currency, err := repo.GetById(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, currency)
}

func TestCurrencyRepository_GetByIso(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewCurrencyCore()

	expected := fixture.NewCurrencyCore()

	db.InsertCurrency(t, &req)

	currency, err := repo.GetByIso(ctx, req.IsoCode)

	require.NoError(t, err)

	require.Equal(t, &expected, currency)
}

func TestCurrencyRepository_GetBySymbol(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewCurrencyCore()

	expected := fixture.NewCurrencyCore()

	db.InsertCurrency(t, &req)

	currency, err := repo.GetBySymbol(ctx, req.Symbol)

	require.NoError(t, err)

	require.Equal(t, &expected, currency)
}

func TestCurrencyRepository_Create(t *testing.T) {
	repo, _, ctx := SetupRepositoryTest(t)

	req := fixture.NewCurrencyCore()

	expected := fixture.NewCurrencyCore()

	currency, err := repo.Create(ctx, &req)

	require.NoError(t, err)

	require.Equal(t, &expected, currency)
}

func TestCurrencyRepository_Update(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewCurrencyUpdateInputCore()

	expected := fixture.NewCurrencyCore()

	db.InsertCurrency(t, &expected)

	currency, err := repo.Update(ctx, expected.Id, req)

	require.NoError(t, err)

	require.Equal(t, &expected, currency)
}

func TestCurrencyRepository_Delete(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	req := fixture.NewCurrencyCore()

	db.InsertCurrency(t, &req)

	err := repo.Delete(ctx, req.Id)

	require.NoError(t, err)
}

func SetupRepositoryTest(t *testing.T) (*CurrencyRepository, *connectToDB.TestDB, context.Context) {
	db := connectToDB.NewTestPostgresDb(t)

	repo := NewCurrencyRepository(db.DB)

	ctx := context.Background()

	db.Cleanup(t)

	return repo, db, ctx
}



