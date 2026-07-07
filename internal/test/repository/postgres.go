package repository

import (
	configs "Bank-repository-service/internal/configs"
	"Bank-repository-service/pkg/core"
	"context"
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/stretchr/testify/require"
)

type TestDB struct {
	DB *sqlx.DB
}

func NewTestPostgresDb(t *testing.T) *TestDB {
	t.Helper()

	cfg := configs.LoadTestConfig()

	db, err := sqlx.Connect("postgres", cfg.Postgres.DBUrl)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = db.Close()
	})

	return &TestDB{
		DB: db,
	}
}

func (tdb *TestDB) Cleanup(t testing.TB) {
	t.Helper()

	var query = `
TRUNCATE TABLE
	users,
	currencies,
	accounts,
	cards,
	credits,
	deposits
RESTART IDENTITY CASCADE;
`

	_, err := tdb.DB.ExecContext(context.Background(), query)
	require.NoError(t, err)
}

func (tdb *TestDB) Seed(t testing.TB, query string, args ...any) {
	t.Helper()

	_, err := tdb.DB.ExecContext(context.Background(), query, args...)
	require.NoError(t, err)
}

func (tdb *TestDB) Exec(t testing.TB, query string, args ...any) sql.Result {
	t.Helper()

	result, err := tdb.DB.ExecContext(context.Background(), query, args...)
	require.NoError(t, err)

	return result
}

func (tdb *TestDB) InsertUser(t testing.TB, user *core.User) {
	query := `
	INSERT INTO users  (id, first_name, middle_name, last_name, email, phone_number) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	tdb.Exec(t, query, user.Id, user.FirstName, user.MiddleName, user.LastName, user.Email, user.PhoneNumber)
}

func (tdb *TestDB) InsertCurrency(t testing.TB, currency *core.Currency) {
	query := `
	INSERT INTO currencies  (id, name, symbol, iso_code, minor_units) 
	VALUES ($1, $2, $3, $4, $5)`

	tdb.Exec(t, query, currency.Id, currency.Name, currency.Symbol, currency.IsoCode, currency.MinorUnits)
}

func (tdb *TestDB) InsertAccount(t testing.TB, account *core.Account) {
	query := `
	INSERT INTO accounts  (id, user_id, currency_id, name, balance, status) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	tdb.Exec(t, query, account.Id, account.UserId, account.CurrencyId, account.Name, account.Balance, account.Status)
}

func (tdb *TestDB) InsertCard(t testing.TB, card *core.Card) {
	query := `
	INSERT INTO cards  (id, user_id, account_id, number, expiry_month, expiry_year, status) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	tdb.Exec(t, query, card.Id, card.UserId, card.AccountId, card.Number, card.ExpiryMonth, card.ExpiryYear, card.Status)
}

func (tdb *TestDB) InsertCredit(t testing.TB, credit *core.Credit) {
	query := `
	INSERT INTO credits  (id, user_id, currency_id, amount, monthly_payment, status) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	tdb.Exec(t, query, credit.Id, credit.UserId, credit.CurrencyId, credit.Amount, credit.MonthlyPayment, credit.Status)
}

func (tdb *TestDB) InsertDeposit(t testing.TB, deposit *core.Deposit) {
	query := `
	INSERT INTO deposits  (id, user_id, currency_id, amount, interest_rate, term_months, status) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	tdb.Exec(t, query, deposit.Id, deposit.UserId, deposit.CurrencyId, deposit.Amount, deposit.InterestRate, deposit.TermMonths, deposit.Status)
}
