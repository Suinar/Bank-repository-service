package configs

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config contains validated runtime settings for the service and its dependencies.
type Config struct {
	Environment     string
	GRPCHost        string
	GRPCPort        string
	Network         string
	ShutdownTimeout time.Duration

	// Postgres contains connection and pool settings for the primary database.
	Postgres struct {
		Host            string
		Port            int
		User            string
		Password        string
		DBName          string
		SSLMode         string
		DBUrl           string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime time.Duration
		ConnMaxIdleTime time.Duration
	}

	// Redis contains connection and pool settings for the currency cache.
	Redis struct {
		Addr         string
		Password     string
		DB           int
		PoolSize     int
		MinIdleConns int
	}

	// Kafka reserves broker and consumer identity settings for event integration.
	Kafka struct {
		Enabled       bool
		Brokers       []string
		ClientID      string
		ConsumerGroup string
	}
}

// LoadConfig reads, validates, and returns the application configuration.
func LoadConfig() (*Config, error) {
	v := newViper()
	if err := readOptionalEnvFile(v, ".env"); err != nil {
		return nil, err
	}

	cfg := &Config{}
	cfg.Environment = v.GetString("APP_ENV")
	cfg.GRPCHost = v.GetString("GRPC_HOST")
	cfg.GRPCPort = v.GetString("GRPC_PORT")
	cfg.Network = v.GetString("NETWORK")
	cfg.ShutdownTimeout = v.GetDuration("SHUTDOWN_TIMEOUT")

	fillPostgres(cfg, v, "DB_")
	fillRedis(cfg, v, "REDIS_")
	fillKafka(cfg, v)

	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

// LoadTestConfig returns integration-test configuration with environment overrides.
func LoadTestConfig() (*Config, error) {
	v := newViper()
	if err := readOptionalEnvFile(v, ".env.test"); err != nil {
		return nil, err
	}

	cfg := &Config{}
	fillPostgres(cfg, v, "TEST_DB_")
	fillRedis(cfg, v, "TEST_REDIS_")

	if cfg.Postgres.Host == "" || cfg.Postgres.Port == 0 || cfg.Postgres.User == "" || cfg.Postgres.DBName == "" {
		return nil, errors.New("test PostgreSQL configuration is incomplete")
	}
	if cfg.Redis.Addr == "" {
		return nil, errors.New("test Redis configuration is incomplete")
	}
	return cfg, nil
}

// Validate checks that all required runtime settings are present and consistent.
func (cfg *Config) Validate() error {
	if cfg.GRPCPort == "" {
		return errors.New("GRPC_PORT is required")
	}
	if cfg.Network == "" {
		return errors.New("NETWORK is required")
	}
	if cfg.Postgres.Host == "" || cfg.Postgres.Port == 0 || cfg.Postgres.User == "" || cfg.Postgres.DBName == "" {
		return errors.New("PostgreSQL configuration is incomplete")
	}
	if cfg.Redis.Addr == "" {
		return errors.New("REDIS_ADDR is required")
	}
	if cfg.Kafka.Enabled && len(cfg.Kafka.Brokers) == 0 {
		return errors.New("KAFKA_BROKERS is required when Kafka is enabled")
	}
	return nil
}

// newViper creates an isolated Viper instance with service defaults.
func newViper() *viper.Viper {
	v := viper.New()
	v.SetConfigType("env")
	v.AutomaticEnv()
	v.SetDefault("APP_ENV", "local")
	v.SetDefault("GRPC_HOST", "0.0.0.0")
	v.SetDefault("GRPC_PORT", "50052")
	v.SetDefault("NETWORK", "tcp")
	v.SetDefault("SHUTDOWN_TIMEOUT", "10s")
	v.SetDefault("DB_SSL_MODE", "disable")
	v.SetDefault("DB_MAX_OPEN_CONNS", 25)
	v.SetDefault("DB_MAX_IDLE_CONNS", 5)
	v.SetDefault("DB_CONN_MAX_LIFETIME", "30m")
	v.SetDefault("DB_CONN_MAX_IDLE_TIME", "5m")
	v.SetDefault("REDIS_POOL_SIZE", 10)
	v.SetDefault("REDIS_MIN_IDLE", 2)
	v.SetDefault("KAFKA_ENABLED", false)
	v.SetDefault("KAFKA_CLIENT_ID", "bank-repository-service")
	v.SetDefault("KAFKA_CONSUMER_GROUP", "bank-repository-service")
	v.SetDefault("TEST_DB_HOST", "localhost")
	v.SetDefault("TEST_DB_PORT", 5432)
	v.SetDefault("TEST_DB_USER", "postgres")
	v.SetDefault("TEST_DB_PASSWORD", "postgres")
	v.SetDefault("TEST_DB_NAME", "bank_test")
	v.SetDefault("TEST_DB_SSL_MODE", "disable")
	v.SetDefault("TEST_REDIS_ADDR", "localhost:6379")
	v.SetDefault("TEST_REDIS_DB", 0)
	return v
}

// readOptionalEnvFile loads a local env file when one is available.
func readOptionalEnvFile(v *viper.Viper, name string) error {
	path, err := findFile(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read %s: %w", name, err)
	}
	return nil
}

// findFile searches the current directory and its parents for a named file.
func findFile(name string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}

// fillPostgres populates the related configuration section from Viper.
func fillPostgres(cfg *Config, v *viper.Viper, prefix string) {
	cfg.Postgres.Host = v.GetString(prefix + "HOST")
	cfg.Postgres.Port = v.GetInt(prefix + "PORT")
	cfg.Postgres.User = v.GetString(prefix + "USER")
	cfg.Postgres.Password = v.GetString(prefix + "PASSWORD")
	cfg.Postgres.DBName = v.GetString(prefix + "NAME")
	cfg.Postgres.SSLMode = v.GetString(prefix + "SSL_MODE")
	if cfg.Postgres.SSLMode == "" {
		cfg.Postgres.SSLMode = "disable"
	}
	cfg.Postgres.MaxOpenConns = v.GetInt(prefix + "MAX_OPEN_CONNS")
	cfg.Postgres.MaxIdleConns = v.GetInt(prefix + "MAX_IDLE_CONNS")
	cfg.Postgres.ConnMaxLifetime = v.GetDuration(prefix + "CONN_MAX_LIFETIME")
	cfg.Postgres.ConnMaxIdleTime = v.GetDuration(prefix + "CONN_MAX_IDLE_TIME")

	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.Postgres.User, cfg.Postgres.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Postgres.Host, cfg.Postgres.Port),
		Path:   cfg.Postgres.DBName,
	}
	query := u.Query()
	query.Set("sslmode", cfg.Postgres.SSLMode)
	u.RawQuery = query.Encode()
	cfg.Postgres.DBUrl = u.String()
}

// fillRedis populates the related configuration section from Viper.
func fillRedis(cfg *Config, v *viper.Viper, prefix string) {
	cfg.Redis.Addr = v.GetString(prefix + "ADDR")
	cfg.Redis.Password = v.GetString(prefix + "PASSWORD")
	cfg.Redis.DB = v.GetInt(prefix + "DB")
	cfg.Redis.PoolSize = v.GetInt(prefix + "POOL_SIZE")
	cfg.Redis.MinIdleConns = v.GetInt(prefix + "MIN_IDLE")
}

// fillKafka populates the related configuration section from Viper.
func fillKafka(cfg *Config, v *viper.Viper) {
	cfg.Kafka.Enabled = v.GetBool("KAFKA_ENABLED")
	cfg.Kafka.ClientID = v.GetString("KAFKA_CLIENT_ID")
	cfg.Kafka.ConsumerGroup = v.GetString("KAFKA_CONSUMER_GROUP")
	for _, broker := range strings.Split(v.GetString("KAFKA_BROKERS"), ",") {
		if broker = strings.TrimSpace(broker); broker != "" {
			cfg.Kafka.Brokers = append(cfg.Kafka.Brokers, broker)
		}
	}
}
