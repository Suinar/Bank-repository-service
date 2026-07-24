package card

import (
	"context"
	fixture "github.com/kVinsom/Bank-repository-service/internal/test/fixture"
	connectToDB "github.com/kVinsom/Bank-repository-service/internal/test/repository"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCardRepository_GetAll(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	account := fixture.NewAccountCore()
	reqId1 := fixture.NewCardCoreInputIdAndNumber(1, "1111111111111111111")
	reqId2 := fixture.NewCardCoreInputIdAndNumber(2, "2222222222222222222")

	expected := []core.Card{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &account)
	db.InsertCard(t, &reqId1)
	db.InsertCard(t, &reqId2)

	cards, err := repo.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, cards, 2)

	require.Equal(t, expected[0], cards[0])
	require.Equal(t, expected[1], cards[1])
}

func TestCardRepository_GetByUser(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	account := fixture.NewAccountCore()
	reqId1 := fixture.NewCardCoreInputIdAndNumber(1, "1111111111111111111")
	reqId2 := fixture.NewCardCoreInputIdAndNumber(2, "2222222222222222222")

	expected := []core.Card{
		reqId1,
		reqId2,
	}

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &account)
	db.InsertCard(t, &reqId1)
	db.InsertCard(t, &reqId2)

	cards, err := repo.GetByUser(ctx, user.Id)

	require.NoError(t, err)
	require.Len(t, cards, 2)

	require.Equal(t, expected[0], cards[0])
	require.Equal(t, expected[1], cards[1])
}

func TestCardRepository_GetById(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	account := fixture.NewAccountCore()
	req := fixture.NewCardCore()

	expected := fixture.NewCardCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &account)
	db.InsertCard(t, &req)

	cards, err := repo.GetById(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, cards)
}

func TestCardRepository_GetByNumber(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	account := fixture.NewAccountCore()
	req := fixture.NewCardCore()

	expected := fixture.NewCardCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &account)
	db.InsertCard(t, &req)

	cards, err := repo.GetByNumber(ctx, req.Number)

	require.NoError(t, err)

	require.Equal(t, &expected, cards)
}

func TestCardRepository_Create(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	account := fixture.NewAccountCore()
	req := fixture.NewCardCore()

	expected := fixture.NewCardCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &account)

	cards, err := repo.Create(ctx, &req)

	require.NoError(t, err)

	require.Equal(t, &expected, cards)
}

func TestCardRepository_Blocking(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	account := fixture.NewAccountCore()
	req := fixture.NewCardCore()

	expected := fixture.NewCardCoreInputStatus(core.CardStatusBlocked)

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &account)
	db.InsertCard(t, &req)

	card, err := repo.Blocking(ctx, req.Id)

	require.NoError(t, err)

	require.Equal(t, &expected, card)
}

func TestCardRepository_Delete(t *testing.T) {
	repo, db, ctx := SetupRepositoryTest(t)

	currency := fixture.NewCurrencyCore()
	user := fixture.NewUserCore()
	account := fixture.NewAccountCore()
	req := fixture.NewCardCore()

	db.InsertCurrency(t, &currency)
	db.InsertUser(t, &user)
	db.InsertAccount(t, &account)
	db.InsertCard(t, &req)

	err := repo.Delete(ctx, req.Id)

	require.NoError(t, err)
}

func SetupRepositoryTest(t *testing.T) (*CardRepository, *connectToDB.TestDB, context.Context) {
	db := connectToDB.NewTestPostgresDb(t)

	repo := NewCardRepository(db.DB)

	ctx := context.Background()

	db.Cleanup(t)

	return repo, db, ctx
}
