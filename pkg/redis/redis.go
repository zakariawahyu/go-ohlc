package redis

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-ohlc/config"
)

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.RedisHost,
		Password: cfg.Redis.RedisPassword,
	})
	ctx := context.Background()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis.InitRedis.Ping")
	}

	return client, nil
}
