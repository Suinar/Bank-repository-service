package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	configs "Bank-repository-service/internal/configs"

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

func (t *TestDB) Cleanup(tb testing.TB) {
	tb.Helper()

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

	_, err := t.DB.ExecContext(context.Background(), query)
	require.NoError(tb, err)
}

func (t *TestDB) Seed(tb testing.TB, query string, args ...any) {
	tb.Helper()

	_, err := t.DB.ExecContext(context.Background(), query, args...)
	require.NoError(tb, err)
}

func (t *TestDB) Exec(tb testing.TB, query string, args ...any) sql.Result {
	tb.Helper()

	result, err := t.DB.ExecContext(context.Background(), query, args...)
	require.NoError(tb, err)

	return result
}
