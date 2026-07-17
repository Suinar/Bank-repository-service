package currency

import (
	fixture "github.com/Suinar/Bank-repository-service/internal/test/fixture"
	test "github.com/Suinar/Bank-repository-service/internal/test/repository"
	errors "github.com/Suinar/Bank-repository-service/pkg"
	"github.com/Suinar/Bank-repository-service/pkg/core"

	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurrencyCache_GetAll_Empty(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)

	result, err := cache.GetAll(context.Background())

	require.NoError(t, err)
	require.Empty(t, result)
}

func TestCurrencyCache_GetAll_InvalidSetID(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)
	ctx := context.Background()
	require.NoError(t, redisDB.DB.SAdd(ctx, currenciesSetKey, "invalid").Err())

	result, err := cache.GetAll(ctx)

	require.ErrorIs(t, err, errors.CacheGetError)
	require.Nil(t, result)
}

func TestCurrencyCache_GetById_InvalidHash(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)
	ctx := context.Background()
	require.NoError(t, redisDB.DB.HSet(ctx, cache.GetPrimaryKey(1), "id", "invalid", "minor_units", "2").Err())

	result, err := cache.GetById(ctx, 1)

	require.ErrorIs(t, err, errors.InternalServerError)
	require.Nil(t, result)
}

func TestCurrencyCache_GetByIso_NotFound(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)

	result, err := cache.GetByIso(context.Background(), "UNKNOWN")

	require.ErrorIs(t, err, errors.CacheGetError)
	require.Nil(t, result)
}

func TestCurrencyCache_GetByIso_InvalidID(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)
	ctx := context.Background()
	require.NoError(t, redisDB.DB.Set(ctx, cache.GetIsoKey("USD"), "invalid", 0).Err())

	result, err := cache.GetByIso(ctx, "USD")

	require.ErrorIs(t, err, errors.BadRequest)
	require.Nil(t, result)
}

func TestCurrencyCache_GetBySymbol_NotFound(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)

	result, err := cache.GetBySymbol(context.Background(), '$')

	require.ErrorIs(t, err, errors.NotFound)
	require.Nil(t, result)
}

func TestCurrencyCache_GetBySymbol_InvalidID(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)
	ctx := context.Background()
	require.NoError(t, redisDB.DB.SAdd(ctx, cache.GetSymbolKey('$'), "invalid").Err())

	result, err := cache.GetBySymbol(ctx, '$')

	require.ErrorIs(t, err, errors.BadRequest)
	require.Nil(t, result)
}

func TestCurrencyCache_SetAll_Empty(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)

	require.NoError(t, cache.SetAll(context.Background(), nil))
}

func TestCurrencyCache_Update_NotFound(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)
	currency := fixture.NewCurrencyCore()

	err := cache.Update(context.Background(), &currency)

	require.ErrorIs(t, err, errors.NotFound)
}

func TestCurrencyCache_Delete_NotFound(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	cache := NewCurrencyCache(redisDB.DB)

	require.NoError(t, cache.Delete(context.Background(), 999999))
}

func TestCurrencyCache_MapToCurrency_InvalidData(t *testing.T) {
	cache := NewCurrencyCache(nil)

	_, err := cache.MapToCurrency(map[string]string{"id": "invalid", "minor_units": "2"})
	require.ErrorIs(t, err, errors.BadRequest)

	_, err = cache.MapToCurrency(map[string]string{"id": "1", "minor_units": "invalid"})
	require.ErrorIs(t, err, errors.BadRequest)

	result, err := cache.MapToCurrency(map[string]string{"id": "1", "minor_units": "2", "symbol": ""})
	require.NoError(t, err)
	require.Equal(t, rune(0), result.Symbol)
}

type failingRedisHook struct {
	command  string
	pipeline bool
}

func TestCurrencyCache_RedisCommandErrors(t *testing.T) {
	tests := []struct {
		name      string
		command   string
		operation func(context.Context, *CurrencyCache) error
		expected  error
	}{
		{
			name:    "get all",
			command: "smembers",
			operation: func(ctx context.Context, cache *CurrencyCache) error {
				_, err := cache.GetAll(ctx)
				return err
			},
			expected: errors.CacheGetError,
		},
		{
			name:    "get by id",
			command: "hgetall",
			operation: func(ctx context.Context, cache *CurrencyCache) error {
				_, err := cache.GetById(ctx, 1)
				return err
			},
			expected: errors.CacheGetError,
		},
		{
			name:    "get by iso",
			command: "get",
			operation: func(ctx context.Context, cache *CurrencyCache) error {
				_, err := cache.GetByIso(ctx, "USD")
				return err
			},
			expected: errors.InternalServerError,
		},
		{
			name:    "get by symbol",
			command: "smembers",
			operation: func(ctx context.Context, cache *CurrencyCache) error {
				_, err := cache.GetBySymbol(ctx, '$')
				return err
			},
			expected: errors.CacheGetError,
		},
		{
			name:    "set all scan",
			command: "scan",
			operation: func(ctx context.Context, cache *CurrencyCache) error {
				return cache.SetAll(ctx, []core.Currency{fixture.NewCurrencyCore()})
			},
			expected: errors.CacheDeleteError,
		},
		{
			name:    "update get",
			command: "hgetall",
			operation: func(ctx context.Context, cache *CurrencyCache) error {
				currency := fixture.NewCurrencyCore()
				return cache.Update(ctx, &currency)
			},
			expected: errors.CacheGetError,
		},
		{
			name:    "delete get",
			command: "hgetall",
			operation: func(ctx context.Context, cache *CurrencyCache) error {
				return cache.Delete(ctx, 1)
			},
			expected: errors.CacheDeleteError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			redisDB := test.NewTestRedisDB(t)
			redisDB.Cleanup(t)
			redisDB.DB.AddHook(failingRedisHook{command: testCase.command})
			cache := NewCurrencyCache(redisDB.DB)

			err := testCase.operation(context.Background(), cache)

			require.ErrorIs(t, err, testCase.expected)
		})
	}
}

func TestCurrencyCache_GetAll_PipelineError(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	ctx := context.Background()
	require.NoError(t, redisDB.DB.SAdd(ctx, currenciesSetKey, "1").Err())
	redisDB.DB.AddHook(failingRedisHook{command: "hgetall", pipeline: true})
	cache := NewCurrencyCache(redisDB.DB)

	result, err := cache.GetAll(ctx)

	require.ErrorIs(t, err, errors.CacheGetError)
	require.Nil(t, result)
}

func TestCurrencyCache_GetAll_SkipsMissingAndMalformedHashes(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	ctx := context.Background()
	require.NoError(t, redisDB.DB.SAdd(ctx, currenciesSetKey, "1", "2").Err())
	require.NoError(t, redisDB.DB.HSet(ctx, "currency:2", "id", "invalid", "minor_units", "2").Err())
	cache := NewCurrencyCache(redisDB.DB)

	result, err := cache.GetAll(ctx)

	require.NoError(t, err)
	require.Empty(t, result)
}

func TestCurrencyCache_Set_PipelineError(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	redisDB.DB.AddHook(failingRedisHook{command: "hset", pipeline: true})
	cache := NewCurrencyCache(redisDB.DB)
	currency := fixture.NewCurrencyCore()

	err := cache.Set(context.Background(), &currency)

	require.ErrorIs(t, err, errors.CacheSetError)
}

func TestCurrencyCache_SetAll_PipelineError(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	redisDB.DB.AddHook(failingRedisHook{command: "hset", pipeline: true})
	cache := NewCurrencyCache(redisDB.DB)

	err := cache.SetAll(context.Background(), []core.Currency{fixture.NewCurrencyCore()})

	require.ErrorIs(t, err, errors.CacheSetError)
}

func TestCurrencyCache_SetAll_ClearUnlinkErrors(t *testing.T) {
	t.Run("currency keys", func(t *testing.T) {
		redisDB := test.NewTestRedisDB(t)
		redisDB.Cleanup(t)
		ctx := context.Background()
		require.NoError(t, redisDB.DB.HSet(ctx, "currency:stale", "id", "1").Err())
		redisDB.DB.AddHook(failingRedisHook{command: "unlink"})
		cache := NewCurrencyCache(redisDB.DB)

		err := cache.SetAll(ctx, []core.Currency{fixture.NewCurrencyCore()})

		require.ErrorIs(t, err, errors.CacheDeleteError)
	})

	t.Run("currency set", func(t *testing.T) {
		redisDB := test.NewTestRedisDB(t)
		redisDB.Cleanup(t)
		redisDB.DB.AddHook(failingRedisHook{command: "unlink"})
		cache := NewCurrencyCache(redisDB.DB)

		err := cache.SetAll(context.Background(), []core.Currency{fixture.NewCurrencyCore()})

		require.ErrorIs(t, err, errors.CacheDeleteError)
	})
}

func TestCurrencyCache_Update_ErrorPaths(t *testing.T) {
	t.Run("malformed existing currency", func(t *testing.T) {
		redisDB := test.NewTestRedisDB(t)
		redisDB.Cleanup(t)
		ctx := context.Background()
		currency := fixture.NewCurrencyCore()
		require.NoError(t, redisDB.DB.HSet(ctx, "currency:1", "id", "invalid", "minor_units", "2").Err())
		cache := NewCurrencyCache(redisDB.DB)

		err := cache.Update(ctx, &currency)

		require.ErrorIs(t, err, errors.InternalServerError)
	})

	t.Run("delete pipeline", func(t *testing.T) {
		redisDB := test.NewTestRedisDB(t)
		redisDB.Cleanup(t)
		ctx := context.Background()
		currency := fixture.NewCurrencyCore()
		cache := NewCurrencyCache(redisDB.DB)
		require.NoError(t, cache.Set(ctx, &currency))
		redisDB.DB.AddHook(failingRedisHook{command: "del", pipeline: true})

		err := cache.Update(ctx, &currency)

		require.ErrorIs(t, err, errors.CacheDeleteError)
	})

	t.Run("set pipeline", func(t *testing.T) {
		redisDB := test.NewTestRedisDB(t)
		redisDB.Cleanup(t)
		ctx := context.Background()
		currency := fixture.NewCurrencyCore()
		cache := NewCurrencyCache(redisDB.DB)
		require.NoError(t, cache.Set(ctx, &currency))
		redisDB.DB.AddHook(failingRedisHook{command: "hset", pipeline: true})

		err := cache.Update(ctx, &currency)

		require.ErrorIs(t, err, errors.CacheSetError)
	})
}

func TestCurrencyCache_Delete_PipelineError(t *testing.T) {
	redisDB := test.NewTestRedisDB(t)
	redisDB.Cleanup(t)
	ctx := context.Background()
	currency := fixture.NewCurrencyCore()
	cache := NewCurrencyCache(redisDB.DB)
	require.NoError(t, cache.Set(ctx, &currency))
	redisDB.DB.AddHook(failingRedisHook{command: "del", pipeline: true})

	err := cache.Delete(ctx, currency.Id)

	require.ErrorIs(t, err, errors.CacheDeleteError)
}
