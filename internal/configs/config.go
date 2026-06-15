package configs

import "github.com/spf13/viper"

type Config struct {
	Redis struct {
		Address      string
		Password     string
		Db           int
		PoolSize     int
		MinIdleConns int
	}

	Postgres struct {
		DbUrl        string
		MaxOpenConns int
		MaxIdleConns int
	}

	Grpc struct {
		MainServerPort string
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	cfg.Redis.Address = viper.GetString("REDIS_ADDR")
	cfg.Redis.Password = viper.GetString("REDIS_PASSWORD")
	cfg.Redis.Db = viper.GetInt("REDIS_DB")
	cfg.Redis.PoolSize = viper.GetInt("REDIS_POOL_SIZE")
	cfg.Redis.MinIdleConns = viper.GetInt("REDIS_MIN_IDLE")

	cfg.Postgres.DbUrl = viper.GetString("DB_PORT")
	cfg.Postgres.MaxOpenConns = viper.GetInt("DB_MAX_OPEN_CONNS")
	cfg.Postgres.MaxIdleConns = viper.GetInt("DB_MAX_IDLE_CONNS")

	cfg.Grpc.MainServerPort = viper.GetString("MAIN_GRPC_PORT")

	return cfg, nil
}
