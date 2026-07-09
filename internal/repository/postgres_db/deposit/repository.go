package deposit

import (
	errror "github.com/Suinar/Bank-repository-service/pkg"
	core "github.com/Suinar/Bank-repository-service/pkg/core"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type DepositRepository struct {
	db *sqlx.DB
}

func NewDepositRepository(db *sqlx.DB) *DepositRepository {
	return &DepositRepository{db: db}
}

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

func (r *DepositRepository) Create(ctx context.Context, input *core.Deposit) (*core.Deposit, error) {
	query := `
INSERT INTO deposits (user_id, currency_id, amount, interest_rate, term_months, status)
VALUES (:user_id, :currency_id, :amount, :interest_rate, :term_months, :status)
RETURNING id, user_id, currency_id, amount, interest_rate, term_months, status;`

	rows, err := r.db.NamedQueryContext(ctx, query, input)
	if err != nil {
		return nil, errror.InternalServerError
	}
	defer rows.Close()

	if rows.Next() {
		var created core.Deposit
		if err := rows.StructScan(&created); err != nil {
			return nil, errror.InternalServerError
		}
		return &created, nil
	}
	return nil, errror.InternalServerError
}

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



