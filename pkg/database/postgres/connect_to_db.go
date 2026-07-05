package postgres

import (
	"Bank-repository-service/internal/configs"
	"log"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg *configs.Config) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Postgres.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}

	return db
}
