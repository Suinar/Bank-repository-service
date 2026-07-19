package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	configs "github.com/kVinsom/Bank-repository-service/internal/configs"
	_ "github.com/lib/pq"
)

// NewPostgresDB opens and verifies a bounded PostgreSQL connection pool.
func NewPostgresDB(ctx context.Context, cfg *configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.Postgres.DBUrl)
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL: %w", err)
	}

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime)

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	return db, nil
}
