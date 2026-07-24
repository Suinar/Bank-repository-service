package currency

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	errror "github.com/kVinsom/Bank-repository-service/pkg"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"

	"github.com/jmoiron/sqlx"
)

// CurrencyRepository persists and retrieves its domain model in PostgreSQL.
type CurrencyRepository struct {
	db *sqlx.DB
}

var bindNamed = sqlx.Named

// NewCurrencyRepository creates a ready-to-use currency repository.
func NewCurrencyRepository(db *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

// GetAll returns all records available through CurrencyRepository.
func (r *CurrencyRepository) GetAll(ctx context.Context) ([]core.Currency, error) {
	query := `
SELECT id, name, symbol, iso_code, minor_units
FROM currencies`

	var currencies []core.Currency

	if err := r.db.SelectContext(ctx, &currencies, query); err != nil {
		return nil, errror.InternalServerError
	}

	return currencies, nil
}

// GetById returns records matching the requested id lookup.
func (r *CurrencyRepository) GetById(ctx context.Context, id int64) (*core.Currency, error) {
	query := `
SELECT id, name, symbol, iso_code, minor_units
FROM currencies
WHERE id = $1`

	var currency core.Currency

	if err := r.db.GetContext(ctx, &currency, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &currency, nil
}

// GetByIso returns records matching the requested iso lookup.
func (r *CurrencyRepository) GetByIso(ctx context.Context, isoCode string) (*core.Currency, error) {
	query := `
SELECT id, name, symbol, iso_code, minor_units
FROM currencies
WHERE iso_code = $1`

	var currency core.Currency

	if err := r.db.GetContext(ctx, &currency, query, isoCode); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &currency, nil
}

// GetBySymbol returns records matching the requested symbol lookup.
func (r *CurrencyRepository) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	query := `
SELECT id, name, symbol, iso_code, minor_units
FROM currencies
WHERE symbol = $1`

	var currency core.Currency

	if err := r.db.GetContext(ctx, &currency, query, symbol); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &currency, nil
}

// Create persists a new record through CurrencyRepository.
func (r *CurrencyRepository) Create(ctx context.Context, input *core.Currency) (*core.Currency, error) {
	query := `
INSERT INTO currencies (name, symbol, iso_code, minor_units)
VALUES (:name, :symbol, :iso_code, :minor_units)
RETURNING id, name, symbol, iso_code, minor_units`

	namedQuery, args, err := bindNamed(query, input)
	if err != nil {
		return nil, errror.InternalServerError
	}

	var created core.Currency
	if err := r.db.QueryRowxContext(ctx, r.db.Rebind(namedQuery), args...).StructScan(&created); err != nil {
		return nil, errror.InternalServerError
	}

	return &created, nil
}

// Update applies the requested changes through CurrencyRepository.
func (r *CurrencyRepository) Update(ctx context.Context, id int64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
	setParts := make([]string, 0, 5)
	args := make([]interface{}, 0, 5)
	argId := 1

	if input.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Symbol != nil {
		setParts = append(setParts, fmt.Sprintf("symbol = $%d", argId))
		args = append(args, *input.Symbol)
		argId++
	}

	if input.IsoCode != nil {
		setParts = append(setParts, fmt.Sprintf("iso_code = $%d", argId))
		args = append(args, *input.IsoCode)
		argId++
	}

	if input.MinorUnits != nil {
		setParts = append(setParts, fmt.Sprintf("minor_units = $%d", argId))
		args = append(args, *input.MinorUnits)
		argId++
	}

	if len(setParts) == 0 {
		return nil, errror.BadRequest
	}

	setParts = append(setParts, "updated_at = NOW()")

	args = append(args, id)

	query := fmt.Sprintf(`
UPDATE currencies
SET %s
WHERE id = $%d
RETURNING id, name, symbol, iso_code, minor_units
`, strings.Join(setParts, ", "), argId)

	var updated core.Currency

	err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errror.NotFound
		}

		return nil, errror.InternalServerError
	}

	return &updated, nil
}

// Delete removes the requested record through CurrencyRepository.
func (r *CurrencyRepository) Delete(ctx context.Context, id int64) error {
	query := `
DELETE FROM currencies
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
