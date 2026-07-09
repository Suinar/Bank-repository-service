package postgres

import (
	configs "github.com/Suinar/Bank-exhange-rate-service/internal/configs"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *configs.Config) *sqlx.DB {
	db, err := sqlx.Connect("postgres", cfg.Postgres.DBUrl)
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


