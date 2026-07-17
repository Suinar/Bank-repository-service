package currency

import (
	"context"
	"errors"
	"strconv"
	"unicode/utf8"

	errror "github.com/Suinar/Bank-repository-service/pkg"
	core "github.com/Suinar/Bank-repository-service/pkg/core"

	"github.com/redis/go-redis/v9"
)

const (
	currenciesSetKey  = "currencies"
	currencyKeyPrefix = "currency:"
)

// CurrencyCache maintains Redis-backed lookup data for its domain.
type CurrencyCache struct {
	rdb *redis.Client
}

// NewCurrencyCache creates a ready-to-use currency cache.
func NewCurrencyCache(rdb *redis.Client) *CurrencyCache {
	return &CurrencyCache{rdb: rdb}
}

// GetAll returns all records available through CurrencyCache.
func (c *CurrencyCache) GetAll(ctx context.Context) ([]core.Currency, error) {
	ids, err := c.rdb.SMembers(ctx, currenciesSetKey).Result()
	if err != nil {
		return nil, errror.CacheGetError
	}

	if len(ids) == 0 {
		return []core.Currency{}, nil
	}

	pipe := c.rdb.Pipeline()
	cmds := make([]*redis.MapStringStringCmd, 0, len(ids))

	for _, id := range ids {
		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, errror.CacheGetError
		}

		key := c.GetPrimaryKey(parsedId)
		cmds = append(cmds, pipe.HGetAll(ctx, key))
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, errror.CacheGetError
	}

	currencies := make([]core.Currency, 0, len(ids))

	for _, cmd := range cmds {
		data := cmd.Val()

		if len(data) == 0 {
			continue
		}

		currency, err := c.MapToCurrency(data)
		if err != nil {
			continue
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

// GetById returns records matching the requested id lookup.
func (c *CurrencyCache) GetById(ctx context.Context, id int64) (*core.Currency, error) {
	primaryKey := c.GetPrimaryKey(id)

	data, err := c.rdb.HGetAll(ctx, primaryKey).Result()
	if err != nil {
		return nil, errror.CacheGetError
	}

	if len(data) == 0 {
		return nil, nil
	}

	currency, err := c.MapToCurrency(data)
	if err != nil {
		return nil, errror.InternalServerError
	}

	return &currency, nil
}

// GetByIso returns records matching the requested iso lookup.
func (c *CurrencyCache) GetByIso(ctx context.Context, iso string) (*core.Currency, error) {
	idStr, err := c.rdb.Get(ctx, c.GetIsoKey(iso)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errror.CacheGetError
		}

		return nil, errror.InternalServerError
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, errror.BadRequest
	}

	return c.GetById(ctx, id)
}

// GetBySymbol returns records matching the requested symbol lookup.
func (c *CurrencyCache) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	ids, err := c.rdb.SMembers(ctx, c.GetSymbolKey(symbol)).Result()
	if err != nil {
		return nil, errror.CacheGetError
	}
	if len(ids) == 0 {
		return nil, errror.NotFound
	}

	id, err := strconv.ParseInt(ids[0], 10, 64)
	if err != nil {
		return nil, errror.BadRequest
	}

	return c.GetById(ctx, id)
}

// Set stores one currency and its lookup indexes in Redis.
func (c *CurrencyCache) Set(ctx context.Context, currency *core.Currency) error {
	primaryKey := c.GetPrimaryKey(currency.Id)
	pipe := c.rdb.Pipeline()
	pipe.HSet(ctx, primaryKey,
		"id", currency.Id,
		"name", currency.Name,
		"symbol", string(currency.Symbol),
		"iso_code", currency.IsoCode,
		"minor_units", currency.MinorUnits,
	)
	pipe.SAdd(ctx, currenciesSetKey, strconv.FormatInt(currency.Id, 10))
	pipe.Set(ctx, c.GetIsoKey(currency.IsoCode), currency.Id, 0)
	pipe.SAdd(ctx, c.GetSymbolKey(currency.Symbol), currency.Id)
	pipe.SAdd(ctx, c.GetMinorUnitsKey(currency.MinorUnits), currency.Id)

	if _, err := pipe.Exec(ctx); err != nil {
		return errror.CacheSetError
	}
	return nil
}

// SetAll replaces the cached currency collection and its lookup indexes.
func (c *CurrencyCache) SetAll(ctx context.Context, currencies []core.Currency) error {
	if len(currencies) == 0 {
		return nil
	}
	if err := c.clear(ctx); err != nil {
		return errror.CacheDeleteError
	}

	pipe := c.rdb.Pipeline()
	for i := range currencies {
		currency := &currencies[i]
		primaryKey := c.GetPrimaryKey(currency.Id)
		pipe.HSet(ctx, primaryKey,
			"id", currency.Id,
			"name", currency.Name,
			"symbol", string(currency.Symbol),
			"iso_code", currency.IsoCode,
			"minor_units", currency.MinorUnits,
		)
		pipe.SAdd(ctx, currenciesSetKey, strconv.FormatInt(currency.Id, 10))
		pipe.Set(ctx, c.GetIsoKey(currency.IsoCode), currency.Id, 0)
		pipe.SAdd(ctx, c.GetSymbolKey(currency.Symbol), currency.Id)
		pipe.SAdd(ctx, c.GetMinorUnitsKey(currency.MinorUnits), currency.Id)
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return errror.CacheSetError
	}
	return nil
}

// Update applies the requested changes through CurrencyCache.
func (c *CurrencyCache) Update(ctx context.Context, currency *core.Currency) error {
	existing, err := c.GetById(ctx, currency.Id)
	if err != nil {
		if errors.Is(err, errror.CacheGetError) {
			return errror.CacheGetError
		}

		return errror.InternalServerError
	}

	if existing == nil {
		return errror.NotFound
	}

	if err := c.Delete(ctx, currency.Id); err != nil {
		return err
	}

	if err := c.Set(ctx, currency); err != nil {
		return err
	}

	return nil
}

// Delete removes the requested record through CurrencyCache.
func (c *CurrencyCache) Delete(ctx context.Context, id int64) error {
	currency, err := c.GetById(ctx, id)
	if err != nil {
		return errror.CacheDeleteError
	}

	if currency == nil {
		return nil
	}

	primaryKey := c.GetPrimaryKey(id)
	isoKey := c.GetIsoKey(currency.IsoCode)
	symbolKey := c.GetSymbolKey(currency.Symbol)
	minorUnitsKey := c.GetMinorUnitsKey(currency.MinorUnits)

	pipe := c.rdb.Pipeline()

	pipe.Del(ctx, primaryKey)
	pipe.Del(ctx, isoKey)
	pipe.SRem(ctx, symbolKey, id)
	pipe.SRem(ctx, minorUnitsKey, id)
	pipe.SRem(ctx, currenciesSetKey, strconv.FormatInt(id, 10))

	_, err = pipe.Exec(ctx)
	if err != nil {
		return errror.CacheDeleteError
	}

	return nil
}

// clear incrementally unlinks currency keys without blocking Redis with KEYS.
func (c *CurrencyCache) clear(ctx context.Context) error {
	var cursor uint64
	for {
		keys, nextCursor, err := c.rdb.Scan(ctx, cursor, currencyKeyPrefix+"*", 128).Result()
		if err != nil {
			return err
		}
		if len(keys) > 0 {
			if err := c.rdb.Unlink(ctx, keys...).Err(); err != nil {
				return err
			}
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return c.rdb.Unlink(ctx, currenciesSetKey).Err()
}

// MapToCurrency converts a Redis hash into a currency domain model.
func (c *CurrencyCache) MapToCurrency(data map[string]string) (core.Currency, error) {
	id, err := strconv.ParseInt(data["id"], 10, 64)
	if err != nil {
		return core.Currency{}, errror.BadRequest
	}

	minorUnits, err := strconv.ParseInt(data["minor_units"], 10, 8)
	if err != nil {
		return core.Currency{}, errror.BadRequest
	}

	var symbol rune
	if value := data["symbol"]; value != "" {
		symbol, _ = utf8.DecodeRuneInString(value)
	}

	return core.Currency{
		Id:         id,
		Name:       data["name"],
		Symbol:     symbol,
		IsoCode:    data["iso_code"],
		MinorUnits: int8(minorUnits),
	}, nil
}

// GetPrimaryKey builds the Redis key for the corresponding currency index.
func (c *CurrencyCache) GetPrimaryKey(id int64) string {
	return currencyKeyPrefix + strconv.FormatInt(id, 10)
}

// GetIsoKey builds the Redis key for the corresponding currency index.
func (c *CurrencyCache) GetIsoKey(iso string) string {
	return currencyKeyPrefix + "iso:" + iso
}

// GetSymbolKey builds the Redis key for the corresponding currency index.
func (c *CurrencyCache) GetSymbolKey(symbol rune) string {
	return currencyKeyPrefix + "symbol:" + string(symbol)
}

// GetMinorUnitsKey builds the Redis key for the corresponding currency index.
func (c *CurrencyCache) GetMinorUnitsKey(minorUnits int8) string {
	return currencyKeyPrefix + "minor_units:" + strconv.Itoa(int(minorUnits))
}
