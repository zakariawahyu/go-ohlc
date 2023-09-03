package helpers

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetRedis(redis *redis.Client, ctx context.Context, key string) (OHLC, error) {
	ohlcBytes, _ := redis.Get(ctx, key).Bytes()

	ohlc := OHLC{}
	if err := json.Unmarshal(ohlcBytes, &ohlc); err != nil {
		return ohlc, err
	}

	return ohlc, nil
}

func SetRedis(redis *redis.Client, ctx context.Context, key string, ttl int, ohlc *OHLC) error {
	ohlcBytes, err := json.Marshal(ohlc)
	if err != nil {
		return err
	}

	if err = redis.Set(ctx, key, ohlcBytes, time.Second*time.Duration(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
