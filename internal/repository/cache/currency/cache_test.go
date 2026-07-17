package currency

import (
	fixture "github.com/Suinar/Bank-repository-service/internal/test/fixture"
	test "github.com/Suinar/Bank-repository-service/internal/test/repository"
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
