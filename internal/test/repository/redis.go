package repository

import (
	config "Bank-repository-service/internal/configs"
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

type TestRedisDB struct {
	DB *redis.Client
}

func NewTestRedisDB(t *testing.T) *TestRedisDB {
	t.Helper()

	cfg := config.LoadTestConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := rdb.Ping(ctx).Err()
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = rdb.Close()
	})

	return &TestRedisDB{
		DB: rdb,
	}
}

func (t *TestRedisDB) Cleanup(tb testing.TB, keys ...string) {
	tb.Helper()

	if len(keys) == 0 {
		return
	}

	err := t.DB.Del(context.Background(), keys...).Err()
	require.NoError(tb, err)
}

func (t *TestRedisDB) Seed(tb testing.TB, key string, value any, expiration time.Duration) {
	tb.Helper()

	err := t.DB.Set(context.Background(), key, value, expiration).Err()
	require.NoError(tb, err)
}

func (t *TestRedisDB) Exec(tb testing.TB, fn func(rdb *redis.Client) error) {
	tb.Helper()

	err := fn(t.DB)
	require.NoError(tb, err)
}
