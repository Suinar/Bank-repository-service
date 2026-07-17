package user

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pkgerrors "github.com/Suinar/Bank-repository-service/pkg"
	"github.com/Suinar/Bank-repository-service/pkg/core"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

var databaseError = errors.New("database error")

func newMockRepository(t *testing.T) (*UserRepository, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })
	return NewUserRepository(sqlx.NewDb(db, "sqlmock")), mock
}

func TestUserRepository_GetAll_Error(t *testing.T) {
	repo, mock := newMockRepository(t)
	mock.ExpectQuery("SELECT id, first_name, middle_name, last_name, email, phone_number").WillReturnError(databaseError)
	result, err := repo.GetAll(context.Background())
	require.Nil(t, result)
	require.ErrorIs(t, err, pkgerrors.InternalServerError)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Lookups_Errors(t *testing.T) {
	tests := []struct {
		name              string
		dbError, expected error
		call              func(*UserRepository) (*core.User, error)
	}{
		{"id not found", sql.ErrNoRows, pkgerrors.NotFound, func(r *UserRepository) (*core.User, error) { return r.GetById(context.Background(), 1) }},
		{"id internal", databaseError, pkgerrors.InternalServerError, func(r *UserRepository) (*core.User, error) { return r.GetById(context.Background(), 1) }},
		{"email not found", sql.ErrNoRows, pkgerrors.NotFound, func(r *UserRepository) (*core.User, error) {
			return r.GetByEmail(context.Background(), "missing@example.com")
		}},
		{"email internal", databaseError, pkgerrors.InternalServerError, func(r *UserRepository) (*core.User, error) {
			return r.GetByEmail(context.Background(), "missing@example.com")
		}},
		{"phone not found", sql.ErrNoRows, pkgerrors.NotFound, func(r *UserRepository) (*core.User, error) {
			return r.GetByPhoneNumber(context.Background(), "+38000000000")
		}},
		{"phone internal", databaseError, pkgerrors.InternalServerError, func(r *UserRepository) (*core.User, error) {
			return r.GetByPhoneNumber(context.Background(), "+38000000000")
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newMockRepository(t)
			mock.ExpectQuery("SELECT id, first_name, middle_name, last_name, email, phone_number").WithArgs(sqlmock.AnyArg()).WillReturnError(tt.dbError)
			result, err := tt.call(repo)
			require.Nil(t, result)
			require.ErrorIs(t, err, tt.expected)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUserRepository_Create_Errors(t *testing.T) {
	t.Run("named query", func(t *testing.T) {
		repo, _ := newMockRepository(t)
		original := bindNamed
		bindNamed = func(string, interface{}) (string, []interface{}, error) { return "", nil, databaseError }
		t.Cleanup(func() { bindNamed = original })
		result, err := repo.Create(context.Background(), &core.User{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
	})
	t.Run("query", func(t *testing.T) {
		repo, mock := newMockRepository(t)
		mock.ExpectQuery("INSERT INTO users").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(databaseError)
		result, err := repo.Create(context.Background(), &core.User{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.InternalServerError)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepository_Update_Errors(t *testing.T) {
	t.Run("empty input", func(t *testing.T) {
		repo, _ := newMockRepository(t)
		result, err := repo.Update(context.Background(), 1, &core.UserUpdateInput{})
		require.Nil(t, result)
		require.ErrorIs(t, err, pkgerrors.BadRequest)
	})
	first, middle, last := "first", "middle", "last"
	input := &core.UserUpdateInput{FirstName: &first, MiddleName: &middle, LastName: &last}
	for _, tt := range []struct {
		name              string
		dbError, expected error
	}{
		{"not found", sql.ErrNoRows, pkgerrors.NotFound},
		{"internal", databaseError, pkgerrors.InternalServerError},
	} {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newMockRepository(t)
			mock.ExpectQuery("UPDATE users").WithArgs(first, middle, last, int64(1)).WillReturnError(tt.dbError)
			result, err := repo.Update(context.Background(), 1, input)
			require.Nil(t, result)
			require.ErrorIs(t, err, tt.expected)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUserRepository_Delete_Errors(t *testing.T) {
	tests := []struct {
		name              string
		result            sql.Result
		dbError, expected error
	}{
		{"exec", nil, databaseError, pkgerrors.InternalServerError},
		{"rows affected", sqlmock.NewErrorResult(databaseError), nil, pkgerrors.InternalServerError},
		{"not found", sqlmock.NewResult(0, 0), nil, pkgerrors.NotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, mock := newMockRepository(t)
			expectation := mock.ExpectExec("DELETE FROM users").WithArgs(int64(1))
			if tt.dbError != nil {
				expectation.WillReturnError(tt.dbError)
			} else {
				expectation.WillReturnResult(tt.result)
			}
			require.ErrorIs(t, repo.Delete(context.Background(), 1), tt.expected)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
