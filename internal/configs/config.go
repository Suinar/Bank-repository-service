package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GrpsPort string
	Network  string

	Postgres struct {
		Host         string
		Port         int
		User         string
		Password     string
		DBName       string
		DBUrl        string
		MaxOpenConns int
		MaxIdleConns int
	}

	Redis struct {
		Addr         string
		Password     string
		DB           int
		PoolSize     int
		MinIdleConns int
	}
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	cfg := &Config{}

	cfg.GrpsPort = viper.GetString("GRPS_PORT")
	cfg.Network = viper.GetString("NETWORK")

	cfg.Postgres.Host = viper.GetString("DB_HOST")
	cfg.Postgres.Port = viper.GetInt("DB_PORT")
	cfg.Postgres.User = viper.GetString("DB_USER")
	cfg.Postgres.Password = viper.GetString("DB_PASSWORD")
	cfg.Postgres.DBName = viper.GetString("DB_NAME")
	cfg.Postgres.MaxOpenConns = viper.GetInt("DB_MAX_OPEN_CONNS")
	cfg.Postgres.MaxIdleConns = viper.GetInt("DB_MAX_IDLE_CONNS")

	cfg.Postgres.DBUrl = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)

	cfg.Redis.Addr = viper.GetString("REDIS_ADDR")
	cfg.Redis.Password = viper.GetString("REDIS_PASSWORD")
	cfg.Redis.DB = viper.GetInt("REDIS_DB")
	cfg.Redis.PoolSize = viper.GetInt("REDIS_POOL_SIZE")
	cfg.Redis.MinIdleConns = viper.GetInt("REDIS_MIN_IDLE")

	return cfg
}

func LoadTestConfig() *Config {
	viper.SetConfigFile("../../../../.env.test")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	cfg := &Config{}

	cfg.Postgres.Host = viper.GetString("TEST_DB_HOST")
	cfg.Postgres.Port = viper.GetInt("TEST_DB_PORT")
	cfg.Postgres.User = viper.GetString("TEST_DB_USER")
	cfg.Postgres.Password = viper.GetString("TEST_DB_PASSWORD")
	cfg.Postgres.DBName = viper.GetString("TEST_DB_NAME")

	cfg.Postgres.DBUrl = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)

	cfg.Redis.Addr = viper.GetString("TEST_REDIS_ADDR")
	cfg.Redis.Password = viper.GetString("TEST_REDIS_PASSWORD")
	cfg.Redis.DB = viper.GetInt("TEST_REDIS_DB")

	return cfg
}


