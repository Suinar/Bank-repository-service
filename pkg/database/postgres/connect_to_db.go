package postgres

import (
	"Bank-repository-service/internal/configs"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewPostgres(cfg *configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.Postgres.DbUrl)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return db, nil
}
