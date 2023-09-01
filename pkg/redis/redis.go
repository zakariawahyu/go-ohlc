package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-ohlc/config"
	"log"
)

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.RedisHost,
		Password: cfg.Redis.RedisPassword,
	})
	ctx := context.Background()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed connect to Redis")
		return nil, err
	}

	return client, nil
}
