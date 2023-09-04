package redis

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-ohlc/entity"
	"time"
)

func GetRedis(redis *redis.Client, ctx context.Context, key string) (entity.OHLC, error) {
	ohlcBytes, _ := redis.Get(ctx, key).Bytes()

	ohlc := entity.OHLC{}
	if err := json.Unmarshal(ohlcBytes, &ohlc); err != nil {
		return ohlc, errors.Wrap(err, "redis.GetRedis.Unmarshal")
	}

	return ohlc, nil
}

func SetRedis(redis *redis.Client, ctx context.Context, key string, ttl int, ohlc *entity.OHLC) error {
	ohlcBytes, err := json.Marshal(ohlc)
	if err != nil {
		return errors.Wrap(err, "redis.SetRedis.Marshal")
	}

	if err = redis.Set(ctx, key, ohlcBytes, time.Second*time.Duration(ttl)).Err(); err != nil {
		return errors.Wrap(err, "redis.SetRedis.Set")
	}

	return nil
}
