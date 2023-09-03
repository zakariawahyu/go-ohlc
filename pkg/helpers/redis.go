package helpers

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-ohlc/entity"
	"time"
)

func GetRedis(redis *redis.Client, ctx context.Context, key string) (entity.OHLC, error) {
	ohlcBytes, _ := redis.Get(ctx, key).Bytes()

	ohlc := entity.OHLC{}
	if err := json.Unmarshal(ohlcBytes, &ohlc); err != nil {
		return ohlc, err
	}

	return ohlc, nil
}

func SetRedis(redis *redis.Client, ctx context.Context, key string, ttl int, ohlc *entity.OHLC) error {
	ohlcBytes, err := json.Marshal(ohlc)
	if err != nil {
		return err
	}

	if err = redis.Set(ctx, key, ohlcBytes, time.Second*time.Duration(ttl)).Err(); err != nil {
		return err
	}

	return nil
}
