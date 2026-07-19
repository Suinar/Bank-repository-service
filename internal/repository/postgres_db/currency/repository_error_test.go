package currency

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	testdb "github.com/kVinsom/Bank-repository-service/internal/test/repository"
	pkgerrors "github.com/kVinsom/Bank-repository-service/pkg"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestCurrencyRepository_ErrorPaths(t *testing.T) {
	dbError := errors.New("database error")
	newRepo := func(t *testing.T) (*CurrencyRepository, sqlmock.Sqlmock) {
		db, mock := testdb.NewMockPostgresDb(t)
		return NewCurrencyRepository(db), mock
	}

	reads := []struct {
		name        string
		arg         interface{}
		dbErr, want error
		call        func(*CurrencyRepository) error
	}{
		{"all internal", nil, dbError, pkgerrors.InternalServerError, func(r *CurrencyRepository) error { _, e := r.GetAll(context.Background()); return e }},
		{"id missing", int64(1), sql.ErrNoRows, pkgerrors.NotFound, func(r *CurrencyRepository) error { _, e := r.GetById(context.Background(), 1); return e }},
		{"id internal", int64(1), dbError, pkgerrors.InternalServerError, func(r *CurrencyRepository) error { _, e := r.GetById(context.Background(), 1); return e }},
		{"iso missing", "USD", sql.ErrNoRows, pkgerrors.NotFound, func(r *CurrencyRepository) error { _, e := r.GetByIso(context.Background(), "USD"); return e }},
		{"iso internal", "USD", dbError, pkgerrors.InternalServerError, func(r *CurrencyRepository) error { _, e := r.GetByIso(context.Background(), "USD"); return e }},
		{"symbol missing", '$', sql.ErrNoRows, pkgerrors.NotFound, func(r *CurrencyRepository) error { _, e := r.GetBySymbol(context.Background(), '$'); return e }},
		{"symbol internal", '$', dbError, pkgerrors.InternalServerError, func(r *CurrencyRepository) error { _, e := r.GetBySymbol(context.Background(), '$'); return e }},
	}
	for _, tt := range reads {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newRepo(t)
			q := mock.ExpectQuery("SELECT id")
			if tt.arg != nil {
				q.WithArgs(tt.arg)
			}
			q.WillReturnError(tt.dbErr)
			require.ErrorIs(t, tt.call(repo), tt.want)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}

	t.Run("create bind", func(t *testing.T) {
		repo, _ := newRepo(t)
		original := bindNamed
		bindNamed = func(string, interface{}) (string, []interface{}, error) { return "", nil, dbError }
		t.Cleanup(func() { bindNamed = original })
		result, err := repo.Create(context.Background(), &core.Currency{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
	})
	t.Run("create query", func(t *testing.T) {
		repo, mock := newRepo(t)
		mock.ExpectQuery("INSERT INTO currencies").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(dbError)
		result, err := repo.Create(context.Background(), &core.Currency{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("empty update", func(t *testing.T) {
		repo, _ := newRepo(t)
		result, err := repo.Update(context.Background(), 1, &core.CurrencyUpdateInput{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.BadRequest)
	})
	name := "Dollar"
	for _, tt := range []struct {
		name        string
		dbErr, want error
	}{{"update missing", sql.ErrNoRows, pkgerrors.NotFound}, {"update internal", dbError, pkgerrors.InternalServerError}} {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newRepo(t)
			mock.ExpectQuery("UPDATE currencies").WithArgs(name, int64(1)).WillReturnError(tt.dbErr)
			result, err := repo.Update(context.Background(), 1, &core.CurrencyUpdateInput{Name: &name})
			require.Nil(t, result)
			require.ErrorIs(t, err, tt.want)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
	for _, tt := range []struct {
		name        string
		result      sql.Result
		dbErr, want error
	}{
		{"delete exec", nil, dbError, pkgerrors.InternalServerError},
		{"delete rows", sqlmock.NewErrorResult(dbError), nil, pkgerrors.InternalServerError},
		{"delete missing", sqlmock.NewResult(0, 0), nil, pkgerrors.NotFound},
	} {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newRepo(t)
			e := mock.ExpectExec("DELETE FROM currencies").WithArgs(int64(1))
			if tt.dbErr != nil {
				e.WillReturnError(tt.dbErr)
			} else {
				e.WillReturnResult(tt.result)
			}
			require.ErrorIs(t, repo.Delete(context.Background(), 1), tt.want)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
