package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ApiPort string

	Postgres struct {
		Host         string
		Port         int
		User         string
		Password     string
		Database     string
		DbUrl        string
		MaxOpenConns int
		MaxIdleConns int
	}

	Redis struct {
		Address      string
		Password     string
		Db           int
		PoolSize     int
		MinIdleConns int
	}
}

func LoadConfig(fileName string) (*Config, error) {
	viper.SetConfigFile(fileName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	cfg.ApiPort = viper.GetString("API_PORT")

	cfg.Postgres.Host = viper.GetString("DB_HOST")
	cfg.Postgres.Port = viper.GetInt("DB_PORT")
	cfg.Postgres.User = viper.GetString("DB_USER")
	cfg.Postgres.Password = viper.GetString("DB_PASSWORD")
	cfg.Postgres.Database = viper.GetString("DB_NAME")
	cfg.Postgres.MaxOpenConns = viper.GetInt("DB_MAX_OPEN_CONNS")
	cfg.Postgres.MaxIdleConns = viper.GetInt("DB_MAX_IDLE_CONNS")

	cfg.Postgres.DbUrl = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	cfg.Redis.Address = viper.GetString("REDIS_ADDR")
	cfg.Redis.Password = viper.GetString("REDIS_PASSWORD")
	cfg.Redis.Db = viper.GetInt("REDIS_DB")
	cfg.Redis.PoolSize = viper.GetInt("REDIS_POOL_SIZE")
	cfg.Redis.MinIdleConns = viper.GetInt("REDIS_MIN_IDLE")

	return cfg, nil
}
