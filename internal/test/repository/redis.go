package repository

import (
	"context"
	config "github.com/kVinsom/Bank-repository-service/internal/configs"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

// TestRedisDB wraps a Redis client with integration-test helpers.
type TestRedisDB struct {
	DB *redis.Client
}

// NewTestRedisDB creates a ready-to-use test redis d b.
func NewTestRedisDB(t *testing.T) *TestRedisDB {
	t.Helper()

	cfg, err := config.LoadTestConfig()
	require.NoError(t, err)

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = rdb.Ping(ctx).Err()
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = rdb.Close()
	})

	return &TestRedisDB{
		DB: rdb,
	}
}

// Cleanup removes data created by the current integration test from TestRedisDB.
func (t *TestRedisDB) Cleanup(tb testing.TB, keys ...string) {
	tb.Helper()

	if len(keys) == 0 {
		err := t.DB.FlushDB(context.Background()).Err()
		require.NoError(tb, err)
		return
	}

	err := t.DB.Del(context.Background(), keys...).Err()
	require.NoError(tb, err)
}

// Seed inserts deterministic data into TestRedisDB for an integration test.
func (t *TestRedisDB) Seed(tb testing.TB, key string, value any, expiration time.Duration) {
	tb.Helper()

	err := t.DB.Set(context.Background(), key, value, expiration).Err()
	require.NoError(tb, err)
}

// Exec executes a test operation against TestRedisDB and asserts success.
func (t *TestRedisDB) Exec(tb testing.TB, fn func(rdb *redis.Client) error) {
	tb.Helper()

	err := fn(t.DB)
	require.NoError(tb, err)
}
