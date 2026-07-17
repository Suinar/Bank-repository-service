package currency

import (
	fixture "github.com/Suinar/Bank-repository-service/internal/test/fixture"
	test "github.com/Suinar/Bank-repository-service/internal/test/repository"
	errors "github.com/Suinar/Bank-repository-service/pkg"
	"github.com/Suinar/Bank-repository-service/pkg/core"
	"strconv"

	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestCurrencyCache_Set_GetAll(t *testing.T) {
	redis := test.NewTestRedisDB(t)
	redis.Cleanup(t)

	cache := NewCurrencyCache(redis.DB)

	expected := []core.Currency{
		fixture.NewCurrencyCoreInputIdAndIsoAndName(1, "USD", "Us dollar"),
		fixture.NewCurrencyCoreInputIdAndIsoAndName(2, "EUR", "Euro"),
	}

	ctx := context.Background()

	err := cache.SetAll(ctx, expected)
	require.NoError(t, err)

	actual, err := cache.GetAll(ctx)

	require.NoError(t, err)
	require.Len(t, actual, len(expected))
	require.Equal(t, expected, actual)
}

func TestCurrencyCache_Set_GetById(t *testing.T) {
	redis := test.NewTestRedisDB(t)
	redis.Cleanup(t)

	cache := NewCurrencyCache(redis.DB)

	expected := fixture.NewCurrencyCore()

	ctx := context.Background()

	err := cache.Set(ctx, &expected)
	require.NoError(t, err)

	actual, err := cache.GetById(ctx, expected.Id)

	require.NoError(t, err)
	require.Equal(t, &expected, actual)
}

func TestCurrencyCache_Set_GetByIso(t *testing.T) {
	redis := test.NewTestRedisDB(t)
	redis.Cleanup(t)

	cache := NewCurrencyCache(redis.DB)

	expected := fixture.NewCurrencyCore()

	ctx := context.Background()

	err := cache.Set(ctx, &expected)
	require.NoError(t, err)

	actual, err := cache.GetByIso(ctx, expected.IsoCode)

	require.NoError(t, err)
	require.Equal(t, &expected, actual)
}

func TestCurrencyCache_Set_GetBySymbol(t *testing.T) {
	redis := test.NewTestRedisDB(t)
	redis.Cleanup(t)

	cache := NewCurrencyCache(redis.DB)

	expected := fixture.NewCurrencyCore()

	ctx := context.Background()

	err := cache.Set(ctx, &expected)
	require.NoError(t, err)

	actual, err := cache.GetBySymbol(ctx, expected.Symbol)

	require.NoError(t, err)
	require.Equal(t, &expected, actual)
}

func TestCurrencyCache_Set(t *testing.T) {
	redis := test.NewTestRedisDB(t)
	redis.Cleanup(t)

	cache := NewCurrencyCache(redis.DB)

	ctx := context.Background()

	currency := core.Currency{
		Id:         1,
		Name:       "Dollar",
		Symbol:     '$',
		IsoCode:    "USD",
		MinorUnits: 2,
	}

	err := cache.Set(ctx, &currency)
	require.NoError(t, err)

	result, err := redis.DB.HGetAll(
		ctx,
		cache.GetPrimaryKey(currency.Id),
	).Result()

	require.NoError(t, err)

	require.Equal(t, "1", result["id"])
	require.Equal(t, "Dollar", result["name"])
	require.Equal(t, "$", result["symbol"])
	require.Equal(t, "USD", result["iso_code"])
	require.Equal(t, "2", result["minor_units"])

	isoID, err := redis.DB.Get(
		ctx,
		cache.GetIsoKey(currency.IsoCode),
	).Int64()

	require.NoError(t, err)
	require.Equal(t, currency.Id, isoID)

	symbols, err := redis.DB.SMembers(
		ctx,
		cache.GetSymbolKey(currency.Symbol),
	).Result()

	require.NoError(t, err)

	require.Contains(
		t,
		symbols,
		strconv.FormatInt(currency.Id, 10),
	)
}

func TestCurrencyCache_SetAll(t *testing.T) {
	redis := test.NewTestRedisDB(t)
	redis.Cleanup(t)

	cache := NewCurrencyCache(redis.DB)

	ctx := context.Background()

	currencies := []core.Currency{
		{
			Id:         1,
			Name:       "Dollar",
			Symbol:     '$',
			IsoCode:    "USD",
			MinorUnits: 2,
		},
		{
			Id:         2,
			Name:       "Euro",
			Symbol:     'ˆ',
			IsoCode:    "EUR",
			MinorUnits: 2,
		},
	}

	err := cache.SetAll(ctx, currencies)

	require.NoError(t, err)

	first, err := cache.GetById(ctx, 1)

	require.NoError(t, err)
	require.NotNil(t, first)

	require.Equal(
		t,
		currencies[0],
		*first,
	)

	second, err := cache.GetById(ctx, 2)

	require.NoError(t, err)
	require.NotNil(t, second)

	require.Equal(
		t,
		currencies[1],
		*second,
	)

	all, err := cache.GetAll(ctx)

	require.NoError(t, err)

	require.Len(t, all, 2)

}

func TestCurrencyCache_Update_Existing(t *testing.T) {
	redis := test.NewTestRedisDB(t)
	redis.Cleanup(t)

	cache := NewCurrencyCache(redis.DB)

	ctx := context.Background()

	oldCurrency := core.Currency{
		Id:         1,
		Name:       "USD",
		Symbol:     '$',
		IsoCode:    "USD",
		MinorUnits: 2,
	}

	err := cache.Set(ctx, &oldCurrency)
	require.NoError(t, err)

	updatedCurrency := core.Currency{
		Id:         1,
		Name:       "US Dollar",
		Symbol:     '$',
		IsoCode:    "USD",
		MinorUnits: 2,
	}

	err = cache.Update(ctx, &updatedCurrency)
	require.NoError(t, err)

	result, err := cache.GetById(ctx, 1)

	require.NoError(t, err)
	require.NotNil(t, result)

	require.Equal(t, updatedCurrency.Name, result.Name)
}

func TestCurrencyCache_Delete_Success(t *testing.T) {
	redis := test.NewTestRedisDB(t)
	redis.Cleanup(t)

	cache := NewCurrencyCache(redis.DB)

	ctx := context.Background()

	currency := core.Currency{
		Id:         1,
		Name:       "USD",
		Symbol:     '$',
		IsoCode:    "USD",
		MinorUnits: 2,
	}

	err := cache.Set(ctx, &currency)
	require.NoError(t, err)

	err = cache.Delete(ctx, currency.Id)
	require.NoError(t, err)

	result, err := cache.GetById(ctx, currency.Id)
	require.NoError(t, err)

	require.Nil(t, result)

	isoValue, err := redis.DB.Get(
		ctx,
		cache.GetIsoKey(currency.IsoCode),
	).Result()

	require.Error(t, err)
	require.Empty(t, isoValue)

	symbolCount, err := redis.DB.SCard(
		ctx,
		cache.GetSymbolKey(currency.Symbol),
	).Result()

	require.NoError(t, err)
	require.Equal(t, int64(0), symbolCount)
}

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

func (h failingRedisHook) DialHook(next redis.DialHook) redis.DialHook {
	return next
}

func (h failingRedisHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		if !h.pipeline && cmd.Name() == h.command {
			return context.Canceled
		}
		return next(ctx, cmd)
	}
}

func (h failingRedisHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		if h.pipeline {
			for _, cmd := range cmds {
				if cmd.Name() == h.command {
					return context.Canceled
				}
			}
		}
		return next(ctx, cmds)
	}
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
