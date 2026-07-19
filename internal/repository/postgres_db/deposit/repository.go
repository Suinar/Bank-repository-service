package deposit

import (
	"context"
	"database/sql"
	"errors"
	errror "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"

	"github.com/jmoiron/sqlx"
)

// DepositRepository persists and retrieves its domain model in PostgreSQL.
type DepositRepository struct {
	db *sqlx.DB
}

var bindNamed = sqlx.Named

// NewDepositRepository creates a ready-to-use deposit repository.
func NewDepositRepository(db *sqlx.DB) *DepositRepository {
	return &DepositRepository{db: db}
}

// GetAll returns all records available through DepositRepository.
func (r *DepositRepository) GetAll(ctx context.Context) ([]core.Deposit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_months, status
FROM deposits`

	var deposits []core.Deposit

	if err := r.db.SelectContext(ctx, &deposits, query); err != nil {
		return nil, errror.InternalServerError
	}

	return deposits, nil
}

// GetByUser returns records matching the requested user lookup.
func (r *DepositRepository) GetByUser(ctx context.Context, userId int64) ([]core.Deposit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_months, status
FROM deposits
WHERE user_id = $1`

	var deposits []core.Deposit

	if err := r.db.SelectContext(ctx, &deposits, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return deposits, nil
}

// GetById returns records matching the requested id lookup.
func (r *DepositRepository) GetById(ctx context.Context, id int64) (*core.Deposit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_months, status
FROM deposits
WHERE id = $1`

	var deposit core.Deposit

	if err := r.db.GetContext(ctx, &deposit, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &deposit, nil
}

// Create persists a new record through DepositRepository.
func (r *DepositRepository) Create(ctx context.Context, input *core.Deposit) (*core.Deposit, error) {
	query := `
INSERT INTO deposits (user_id, currency_id, amount, interest_rate, term_months, status)
VALUES (:user_id, :currency_id, :amount, :interest_rate, :term_months, :status)
RETURNING id, user_id, currency_id, amount, interest_rate, term_months, status;`

	namedQuery, args, err := bindNamed(query, input)
	if err != nil {
		return nil, errror.InternalServerError
	}

	var created core.Deposit
	if err := r.db.QueryRowxContext(ctx, r.db.Rebind(namedQuery), args...).StructScan(&created); err != nil {
		return nil, errror.InternalServerError
	}

	return &created, nil
}

// Replenish adds funds to the requested deposit through DepositRepository.
func (r *DepositRepository) Replenish(ctx context.Context, id int64, amount int64) (*core.Deposit, error) {
	query := `
UPDATE deposits
SET amount = amount + $2
WHERE id = $1
AND status = 1
RETURNING id, user_id, currency_id, amount, interest_rate, term_months, status`

	var deposit core.Deposit

	err := r.db.GetContext(ctx, &deposit, query, id, amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &deposit, nil
}

// Delete removes the requested record through DepositRepository.
func (r *DepositRepository) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM deposits
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
