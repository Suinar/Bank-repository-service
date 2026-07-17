package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	errror "github.com/Suinar/Bank-repository-service/pkg"
	core "github.com/Suinar/Bank-repository-service/pkg/core"

	"github.com/jmoiron/sqlx"
)

// AccountRepository persists and retrieves its domain model in PostgreSQL.
type AccountRepository struct {
	db *sqlx.DB
}

var bindNamed = sqlx.Named

// NewAccountRepository creates a ready-to-use account repository.
func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// GetAll returns all records available through AccountRepository.
func (r *AccountRepository) GetAll(ctx context.Context) ([]core.Account, error) {
	query := `
SELECT id, user_id, currency_id, name, balance, status
FROM accounts
    `

	var accounts []core.Account

	if err := r.db.SelectContext(ctx, &accounts, query); err != nil {
		return nil, errror.InternalServerError
	}

	return accounts, nil
}

// GetByUser returns records matching the requested user lookup.
func (r *AccountRepository) GetByUser(ctx context.Context, userId int64) ([]core.Account, error) {
	query := `
SELECT id, user_id, currency_id, name, balance, status
FROM accounts
WHERE user_id = $1`

	var accounts []core.Account

	if err := r.db.SelectContext(ctx, &accounts, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return accounts, nil
}

// GetById returns records matching the requested id lookup.
func (r *AccountRepository) GetById(ctx context.Context, id int64) (*core.Account, error) {
	query := `
SELECT id, user_id, currency_id, name, balance, status
FROM accounts
WHERE id = $1`

	var account core.Account

	if err := r.db.GetContext(ctx, &account, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &account, nil
}

// Create persists a new record through AccountRepository.
func (r *AccountRepository) Create(ctx context.Context, input *core.Account) (*core.Account, error) {
	query := `
INSERT INTO accounts (user_id, currency_id, name, balance, status)
VALUES (:user_id, :currency_id, :name, :balance, :status)
RETURNING id, user_id, currency_id, name, balance, status;`

	namedQuery, args, err := bindNamed(query, input)
	if err != nil {
		return nil, errror.InternalServerError
	}

	var created core.Account
	if err := r.db.QueryRowxContext(ctx, r.db.Rebind(namedQuery), args...).StructScan(&created); err != nil {
		return nil, errror.InternalServerError
	}

	return &created, nil
}

// Blocking moves the requested record to its blocked state through AccountRepository.
func (r *AccountRepository) Blocking(ctx context.Context, id int64) (*core.Account, error) {
	query := `
	UPDATE accounts
	SET status = $1, updated_at = $2
	WHERE id = $3
	RETURNING id, user_id, currency_id, name, balance, status
	`

	var account core.Account

	err := r.db.QueryRowContext(ctx, query, core.AccountStatusBlocked, time.Now(), id).
		Scan(&account.Id, &account.UserId, &account.CurrencyId, &account.Name, &account.Balance, &account.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &account, nil
}

// Close moves the requested account to its closed state through AccountRepository.
func (r *AccountRepository) Close(ctx context.Context, id int64) (*core.Account, error) {
	query := `
	UPDATE accounts
	SET status = $1, updated_at = $2
	WHERE id = $3
	RETURNING id, user_id, currency_id, name, balance, status
	`

	var account core.Account

	err := r.db.QueryRowContext(ctx, query, core.AccountStatusClosed, time.Now(), id).
		Scan(&account.Id, &account.UserId, &account.CurrencyId, &account.Name, &account.Balance, &account.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &account, nil
}

// Update applies the requested changes through AccountRepository.
func (r *AccountRepository) Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error) {
	setParts := make([]string, 0, 2)
	args := make([]interface{}, 0, 3)
	argId := 1

	if input.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if len(setParts) == 0 {
		return nil, errror.BadRequest
	}

	setParts = append(setParts, "updated_at = NOW()")

	args = append(args, id)
	args = append(args, core.AccountStatusClosed)

	query := fmt.Sprintf(`
UPDATE accounts
SET %s
WHERE id = $%d AND status != $%d
RETURNING id, user_id, currency_id, name, balance, status
`, strings.Join(setParts, ", "), argId, argId+1)

	var updated core.Account

	err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &updated, nil
}

// Delete removes the requested record through AccountRepository.
func (r *AccountRepository) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM accounts
	WHERE id = $1
	`

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
