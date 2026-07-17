package account

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	testdb "github.com/Suinar/Bank-repository-service/internal/test/repository"
	pkgerrors "github.com/Suinar/Bank-repository-service/pkg"
	"github.com/Suinar/Bank-repository-service/pkg/core"
	"github.com/stretchr/testify/require"
)

var databaseError = errors.New("database error")

func mockRepository(t *testing.T) (*AccountRepository, sqlmock.Sqlmock) {
	db, mock := testdb.NewMockPostgresDb(t)
	return NewAccountRepository(db), mock
}

func TestAccountRepository_ReadErrors(t *testing.T) {
	tests := []struct {
		name, query       string
		dbError, expected error
		args              []driver.Value
		call              func(*AccountRepository) error
	}{
		{"all", "SELECT id", databaseError, pkgerrors.InternalServerError, nil, func(r *AccountRepository) error { _, e := r.GetAll(context.Background()); return e }},
		{"user missing", "SELECT id", sql.ErrNoRows, pkgerrors.NotFound, []driver.Value{int64(1)}, func(r *AccountRepository) error { _, e := r.GetByUser(context.Background(), 1); return e }},
		{"user internal", "SELECT id", databaseError, pkgerrors.InternalServerError, []driver.Value{int64(1)}, func(r *AccountRepository) error { _, e := r.GetByUser(context.Background(), 1); return e }},
		{"id missing", "SELECT id", sql.ErrNoRows, pkgerrors.NotFound, []driver.Value{int64(1)}, func(r *AccountRepository) error { _, e := r.GetById(context.Background(), 1); return e }},
		{"id internal", "SELECT id", databaseError, pkgerrors.InternalServerError, []driver.Value{int64(1)}, func(r *AccountRepository) error { _, e := r.GetById(context.Background(), 1); return e }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := mockRepository(t)
			q := mock.ExpectQuery(tt.query)
			if tt.args != nil {
				q.WithArgs(tt.args...)
			}
			q.WillReturnError(tt.dbError)
			require.ErrorIs(t, tt.call(repo), tt.expected)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAccountRepository_WriteErrors(t *testing.T) {
	t.Run("create bind", func(t *testing.T) {
		repo, _ := mockRepository(t)
		original := bindNamed
		bindNamed = func(string, interface{}) (string, []interface{}, error) { return "", nil, databaseError }
		t.Cleanup(func() { bindNamed = original })
		result, err := repo.Create(context.Background(), &core.Account{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
	})
	t.Run("create query", func(t *testing.T) {
		repo, mock := mockRepository(t)
		mock.ExpectQuery("INSERT INTO accounts").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(databaseError)
		result, err := repo.Create(context.Background(), &core.Account{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	for _, action := range []struct {
		name   string
		status core.AccountStatus
		call   func(*AccountRepository) error
	}{
		{"blocking missing", core.AccountStatusBlocked, func(r *AccountRepository) error { _, e := r.Blocking(context.Background(), 1); return e }},
		{"blocking internal", core.AccountStatusBlocked, func(r *AccountRepository) error { _, e := r.Blocking(context.Background(), 1); return e }},
		{"close missing", core.AccountStatusClosed, func(r *AccountRepository) error { _, e := r.Close(context.Background(), 1); return e }},
		{"close internal", core.AccountStatusClosed, func(r *AccountRepository) error { _, e := r.Close(context.Background(), 1); return e }},
	} {
		t.Run(action.name, func(t *testing.T) {
			repo, mock := mockRepository(t)
			dbErr := databaseError
			expected := pkgerrors.InternalServerError
			if action.name == "blocking missing" || action.name == "close missing" {
				dbErr = sql.ErrNoRows
				expected = pkgerrors.NotFound
			}
			mock.ExpectQuery("UPDATE accounts").WithArgs(action.status, sqlmock.AnyArg(), int64(1)).WillReturnError(dbErr)
			require.ErrorIs(t, action.call(repo), expected)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
	name := "updated"
	for _, tt := range []struct {
		name              string
		dbError, expected error
	}{{"update missing", sql.ErrNoRows, pkgerrors.NotFound}, {"update internal", databaseError, pkgerrors.InternalServerError}} {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := mockRepository(t)
			mock.ExpectQuery("UPDATE accounts").WithArgs(name, int64(1), core.AccountStatusClosed).WillReturnError(tt.dbError)
			result, err := repo.Update(context.Background(), 1, &core.AccountUpdateInput{Name: &name})
			require.Nil(t, result)
			require.ErrorIs(t, err, tt.expected)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
	t.Run("empty update", func(t *testing.T) {
		repo, _ := mockRepository(t)
		result, err := repo.Update(context.Background(), 1, &core.AccountUpdateInput{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.BadRequest)
	})
}

func TestAccountRepository_DeleteErrors(t *testing.T) {
	tests := []struct {
		name              string
		result            sql.Result
		dbError, expected error
	}{{"exec", nil, databaseError, pkgerrors.InternalServerError}, {"rows", sqlmock.NewErrorResult(databaseError), nil, pkgerrors.InternalServerError}, {"missing", sqlmock.NewResult(0, 0), nil, pkgerrors.NotFound}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := mockRepository(t)
			e := mock.ExpectExec("DELETE FROM accounts").WithArgs(int64(1))
			if tt.dbError != nil {
				e.WillReturnError(tt.dbError)
			} else {
				e.WillReturnResult(tt.result)
			}
			require.ErrorIs(t, repo.Delete(context.Background(), 1), tt.expected)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
