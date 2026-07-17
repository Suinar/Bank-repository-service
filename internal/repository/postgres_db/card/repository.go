package card

import (
	"context"
	"database/sql"
	"errors"
	"time"

	errror "github.com/Suinar/Bank-repository-service/pkg"
	core "github.com/Suinar/Bank-repository-service/pkg/core"

	"github.com/jmoiron/sqlx"
)

// CardRepository persists and retrieves its domain model in PostgreSQL.
type CardRepository struct {
	db *sqlx.DB
}

// NewCardRepository creates a ready-to-use card repository.
func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{db: db}
}

// GetAll returns all records available through CardRepository.
func (r *CardRepository) GetAll(ctx context.Context) ([]core.Card, error) {
	query := `
SELECT id, user_id, account_id, number, expiry_month, expiry_year, status
FROM cards`

	var cards []core.Card

	if err := r.db.SelectContext(ctx, &cards, query); err != nil {
		return nil, errror.InternalServerError
	}

	return cards, nil
}

// GetByUser returns records matching the requested user lookup.
func (r *CardRepository) GetByUser(ctx context.Context, userId int64) ([]core.Card, error) {
	query := `
SELECT id, user_id, account_id, number, expiry_month, expiry_year, status
FROM cards
WHERE user_id = $1`

	var cards []core.Card

	if err := r.db.SelectContext(ctx, &cards, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return cards, nil
}

// GetById returns records matching the requested id lookup.
func (r *CardRepository) GetById(ctx context.Context, id int64) (*core.Card, error) {
	query := `
SELECT id, user_id, account_id, number, expiry_month, expiry_year, status
FROM cards
WHERE id = $1`

	var card core.Card

	if err := r.db.GetContext(ctx, &card, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &card, nil
}

// GetByNumber returns records matching the requested number lookup.
func (r *CardRepository) GetByNumber(ctx context.Context, number string) (*core.Card, error) {
	query := `
SELECT id, user_id, account_id, number, expiry_month, expiry_year, status
FROM cards
WHERE number = $1`

	var card core.Card

	if err := r.db.GetContext(ctx, &card, query, number); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &card, nil
}

// Blocking moves the requested record to its blocked state through CardRepository.
func (r *CardRepository) Blocking(ctx context.Context, id int64) (*core.Card, error) {
	query := `
	UPDATE cards
	SET status = $1, updated_at = $2
	WHERE id = $3
	RETURNING id, user_id, account_id, number, expiry_month, expiry_year, status
	`

	var card core.Card

	err := r.db.QueryRowContext(ctx, query, core.CardStatusBlocked, time.Now(), id).
		Scan(&card.Id, &card.UserId, &card.AccountId, &card.Number, &card.ExpiryMonth, &card.ExpiryYear, &card.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &card, nil
}

// Create persists a new record through CardRepository.
func (r *CardRepository) Create(ctx context.Context, input *core.Card) (*core.Card, error) {
	query := `
INSERT INTO cards (user_id, account_id, number, expiry_month, expiry_year, status)
VALUES (:user_id, :account_id, :number, :expiry_month, :expiry_year, :status)
RETURNING id, user_id, account_id, number, expiry_month, expiry_year, status;`

	namedQuery, args, err := sqlx.Named(query, input)
	if err != nil {
		return nil, errror.InternalServerError
	}

	var created core.Card
	if err := r.db.QueryRowxContext(ctx, r.db.Rebind(namedQuery), args...).StructScan(&created); err != nil {
		return nil, errror.InternalServerError
	}

	return &created, nil
}

// Delete removes the requested record through CardRepository.
func (r *CardRepository) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM cards
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
