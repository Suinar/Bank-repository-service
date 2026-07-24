package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	errror "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
	"strings"

	"github.com/jmoiron/sqlx"
)

// UserRepository persists and retrieves its domain model in PostgreSQL.
type UserRepository struct {
	db *sqlx.DB
}

var bindNamed = sqlx.Named

// NewUserRepository creates a ready-to-use user repository.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAll returns all records available through UserRepository.
func (r *UserRepository) GetAll(ctx context.Context) ([]core.User, error) {
	query := `
SELECT id, first_name, middle_name, last_name, email, phone_number
FROM users`

	var users []core.User

	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return nil, errror.InternalServerError
	}

	return users, nil
}

// GetById returns records matching the requested id lookup.
func (r *UserRepository) GetById(ctx context.Context, id int64) (*core.User, error) {
	query := `
SELECT id, first_name, middle_name, last_name, email, phone_number
FROM users
WHERE id = $1`

	var user core.User

	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &user, nil
}

// GetByEmail returns records matching the requested email lookup.
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*core.User, error) {
	query := `
SELECT id, first_name, middle_name, last_name, email, phone_number
FROM users
WHERE email = $1`

	var user core.User

	if err := r.db.GetContext(ctx, &user, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &user, nil
}

// GetByPhoneNumber returns records matching the requested phone number lookup.
func (r *UserRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error) {
	query := `
SELECT id, first_name, middle_name, last_name, email, phone_number
FROM users
WHERE phone_number = $1`

	var user core.User

	if err := r.db.GetContext(ctx, &user, query, phoneNumber); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &user, nil
}

// Create persists a new record through UserRepository.
func (r *UserRepository) Create(ctx context.Context, input *core.User) (*core.User, error) {
	query := `
INSERT INTO users (first_name, middle_name, last_name, email, phone_number)
VALUES (:first_name, :middle_name, :last_name, :email, :phone_number)
RETURNING id, first_name, middle_name, last_name, email, phone_number;`

	namedQuery, args, err := bindNamed(query, input)
	if err != nil {
		return nil, errror.InternalServerError
	}

	var created core.User
	if err := r.db.QueryRowxContext(ctx, r.db.Rebind(namedQuery), args...).StructScan(&created); err != nil {
		return nil, errror.InternalServerError
	}

	return &created, nil
}

// Update applies the requested changes through UserRepository.
func (r *UserRepository) Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error) {
	setParts := make([]string, 0, 4)
	args := make([]interface{}, 0, 4)
	argId := 1

	if input.FirstName != nil {
		setParts = append(setParts, fmt.Sprintf("first_name = $%d", argId))
		args = append(args, *input.FirstName)
		argId++
	}
	if input.MiddleName != nil {
		setParts = append(setParts, fmt.Sprintf("middle_name = $%d", argId))
		args = append(args, *input.MiddleName)
		argId++
	}
	if input.LastName != nil {
		setParts = append(setParts, fmt.Sprintf("last_name = $%d", argId))
		args = append(args, *input.LastName)
		argId++
	}

	if len(setParts) == 0 {
		return nil, errror.BadRequest
	}

	setParts = append(setParts, "updated_at = NOW()")

	args = append(args, id)

	query := fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id = $%d
		RETURNING id, first_name, middle_name, last_name, email, phone_number
	`, strings.Join(setParts, ", "), argId)

	var updated core.User

	err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}
		return nil, errror.InternalServerError
	}

	return &updated, nil
}

// ChangePassword updates the password hash for the requested user.
func (r *UserRepository) ChangePassword(ctx context.Context, id int64, password string) error {
	query := `
UPDATE users
SET password_hash = $1, updated_at = NOW()
WHERE id = $2`

	result, err := r.db.ExecContext(ctx, query, password, id)
	if err != nil {
		return errror.InternalServerError
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errror.InternalServerError
	}

	if rowsAffected == 0 {
		return errror.NotFound
	}

	return nil
}

// Delete removes the requested record through UserRepository.
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `
DELETE FROM users 
WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errror.InternalServerError
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errror.InternalServerError
	}

	if rowsAffected == 0 {
		return errror.NotFound
	}

	return nil
}
