package credit

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

func TestCreditRepository_ErrorPaths(t *testing.T) {
	dbError := errors.New("database error")
	newRepo := func(t *testing.T) (*CreditRepository, sqlmock.Sqlmock) {
		db, mock := connectToDB.NewMockPostgresDb(t)
		return NewCreditRepository(db), mock
	}
	reads := []struct {
		name        string
		args        []interface{}
		dbErr, want error
		call        func(*CreditRepository) error
	}{
		{"all internal", nil, dbError, pkgerrors.InternalServerError, func(r *CreditRepository) error { _, e := r.GetAll(context.Background()); return e }},
		{"user missing", []interface{}{int64(1)}, sql.ErrNoRows, pkgerrors.NotFound, func(r *CreditRepository) error { _, e := r.GetByUser(context.Background(), 1); return e }},
		{"user internal", []interface{}{int64(1)}, dbError, pkgerrors.InternalServerError, func(r *CreditRepository) error { _, e := r.GetByUser(context.Background(), 1); return e }},
		{"id missing", []interface{}{int64(1)}, sql.ErrNoRows, pkgerrors.NotFound, func(r *CreditRepository) error { _, e := r.GetById(context.Background(), 1); return e }},
		{"id internal", []interface{}{int64(1)}, dbError, pkgerrors.InternalServerError, func(r *CreditRepository) error { _, e := r.GetById(context.Background(), 1); return e }},
		{"repay missing", []interface{}{int64(1), int64(2)}, sql.ErrNoRows, pkgerrors.NotFound, func(r *CreditRepository) error { _, e := r.Repay(context.Background(), 1, 2); return e }},
		{"repay internal", []interface{}{int64(1), int64(2)}, dbError, pkgerrors.InternalServerError, func(r *CreditRepository) error { _, e := r.Repay(context.Background(), 1, 2); return e }},
	}
	for _, tt := range reads {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newRepo(t)
			q := mock.ExpectQuery("(SELECT id|UPDATE credits)")
			if tt.args != nil {
				if len(tt.args) == 1 {
					q.WithArgs(tt.args[0])
				} else {
					q.WithArgs(tt.args[0], tt.args[1])
				}
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
		result, err := repo.Create(context.Background(), &core.Credit{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
	})
	t.Run("create query", func(t *testing.T) {
		repo, mock := newRepo(t)
		mock.ExpectQuery("INSERT INTO credits").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(dbError)
		result, err := repo.Create(context.Background(), &core.Credit{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
		require.NoError(t, mock.ExpectationsWereMet())
	})
	for _, tt := range []struct {
		name        string
		result      sql.Result
		dbErr, want error
	}{{"delete exec", nil, dbError, pkgerrors.InternalServerError}, {"delete rows", sqlmock.NewErrorResult(dbError), nil, pkgerrors.InternalServerError}, {"delete missing", sqlmock.NewResult(0, 0), nil, pkgerrors.NotFound}} {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newRepo(t)
			e := mock.ExpectExec("DELETE FROM credits").WithArgs(int64(1))
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
