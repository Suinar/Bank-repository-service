package card

import (
	"context"
	"database/sql"
	"errors"
	connectToDB "github.com/Suinar/Bank-repository-service/internal/test/repository"
	pkgerrors "github.com/Suinar/Bank-repository-service/pkg"
	core "github.com/Suinar/Bank-repository-service/pkg/core"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestCardRepository_ErrorPaths(t *testing.T) {
	dbError := errors.New("database error")
	newRepo := func(t *testing.T) (*CardRepository, sqlmock.Sqlmock) {
		db, mock := connectToDB.NewMockPostgresDb(t)
		return NewCardRepository(db), mock
	}

	reads := []struct {
		name        string
		args        []interface{}
		dbErr, want error
		call        func(*CardRepository) error
	}{
		{"all internal", nil, dbError, pkgerrors.InternalServerError, func(r *CardRepository) error { _, err := r.GetAll(context.Background()); return err }},
		{"user missing", []interface{}{int64(1)}, sql.ErrNoRows, pkgerrors.NotFound, func(r *CardRepository) error { _, err := r.GetByUser(context.Background(), 1); return err }},
		{"user internal", []interface{}{int64(1)}, dbError, pkgerrors.InternalServerError, func(r *CardRepository) error { _, err := r.GetByUser(context.Background(), 1); return err }},
		{"id missing", []interface{}{int64(1)}, sql.ErrNoRows, pkgerrors.NotFound, func(r *CardRepository) error { _, err := r.GetById(context.Background(), 1); return err }},
		{"id internal", []interface{}{int64(1)}, dbError, pkgerrors.InternalServerError, func(r *CardRepository) error { _, err := r.GetById(context.Background(), 1); return err }},
		{"number missing", []interface{}{"1"}, sql.ErrNoRows, pkgerrors.NotFound, func(r *CardRepository) error { _, err := r.GetByNumber(context.Background(), "1"); return err }},
		{"number internal", []interface{}{"1"}, dbError, pkgerrors.InternalServerError, func(r *CardRepository) error { _, err := r.GetByNumber(context.Background(), "1"); return err }},
	}
	for _, tt := range reads {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newRepo(t)
			expectation := mock.ExpectQuery("SELECT id")
			if tt.args != nil {
				for _, arg := range tt.args {
					expectation.WithArgs(arg)
				}
			}
			expectation.WillReturnError(tt.dbErr)
			require.ErrorIs(t, tt.call(repo), tt.want)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}

	t.Run("create bind", func(t *testing.T) {
		repo, _ := newRepo(t)
		original := bindNamed
		bindNamed = func(string, interface{}) (string, []interface{}, error) { return "", nil, dbError }
		t.Cleanup(func() { bindNamed = original })
		result, err := repo.Create(context.Background(), &core.Card{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
	})
	t.Run("create query", func(t *testing.T) {
		repo, mock := newRepo(t)
		mock.ExpectQuery("INSERT INTO cards").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(dbError)
		result, err := repo.Create(context.Background(), &core.Card{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	for _, tt := range []struct {
		name        string
		dbErr, want error
	}{{"blocking missing", sql.ErrNoRows, pkgerrors.NotFound}, {"blocking internal", dbError, pkgerrors.InternalServerError}} {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newRepo(t)
			mock.ExpectQuery("UPDATE cards").WithArgs(core.CardStatusBlocked, sqlmock.AnyArg(), int64(1)).WillReturnError(tt.dbErr)
			result, err := repo.Blocking(context.Background(), 1)
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
			e := mock.ExpectExec("DELETE FROM cards").WithArgs(int64(1))
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
