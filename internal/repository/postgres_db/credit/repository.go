package credit

import (
	"context"
	"database/sql"
	"errors"
	errror "github.com/Suinar/Bank-repository-service/pkg"
	core "github.com/Suinar/Bank-repository-service/pkg/core"

	"github.com/jmoiron/sqlx"
)

// CreditRepository persists and retrieves its domain model in PostgreSQL.
type CreditRepository struct {
	db *sqlx.DB
}

var bindNamed = sqlx.Named

// NewCreditRepository creates a ready-to-use credit repository.
func NewCreditRepository(db *sqlx.DB) *CreditRepository {
	return &CreditRepository{db: db}
}

// GetAll returns all records available through CreditRepository.
func (r *CreditRepository) GetAll(ctx context.Context) ([]core.Credit, error) {
	query := `
SELECT id, user_id, currency_id, amount, monthly_payment, status
FROM credits`

	var credits []core.Credit

	if err := r.db.SelectContext(ctx, &credits, query); err != nil {
		return nil, errror.InternalServerError
	}

	return credits, nil
}

// GetByUser returns records matching the requested user lookup.
func (r *CreditRepository) GetByUser(ctx context.Context, userId int64) ([]core.Credit, error) {
	query := `
SELECT id, user_id, currency_id, amount, monthly_payment, status
FROM credits
WHERE user_id = $1`

	var credits []core.Credit

	if err := r.db.SelectContext(ctx, &credits, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return credits, nil
}

// GetById returns records matching the requested id lookup.
func (r *CreditRepository) GetById(ctx context.Context, id int64) (*core.Credit, error) {
	query := `
SELECT id, user_id, currency_id, amount, monthly_payment, status
FROM credits
WHERE id = $1`

	var credits core.Credit

	if err := r.db.GetContext(ctx, &credits, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &credits, nil
}

// Create persists a new record through CreditRepository.
func (r *CreditRepository) Create(ctx context.Context, input *core.Credit) (*core.Credit, error) {
	query := `
INSERT INTO credits (user_id, currency_id, amount, monthly_payment, status)
VALUES (:user_id, :currency_id, :amount, :monthly_payment, :status)
Returning id, user_id, currency_id, amount, monthly_payment, status;`

	namedQuery, args, err := bindNamed(query, input)
	if err != nil {
		return nil, errror.InternalServerError
	}

	var created core.Credit
	if err := r.db.QueryRowxContext(ctx, r.db.Rebind(namedQuery), args...).StructScan(&created); err != nil {
		return nil, errror.InternalServerError
	}

	return &created, nil
}

// Repay applies a repayment to the requested credit through CreditRepository.
func (r *CreditRepository) Repay(ctx context.Context, id int64, amount int64) (*core.Credit, error) {
	query := `
UPDATE credits
SET amount = amount - $2
WHERE id = $1
AND amount >= $2
RETURNING id, user_id, currency_id, amount, monthly_payment, status`

	var credit core.Credit

	err := r.db.GetContext(ctx, &credit, query, id, amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &credit, nil
}

// Delete removes the requested record through CreditRepository.
func (r *CreditRepository) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM credits
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
