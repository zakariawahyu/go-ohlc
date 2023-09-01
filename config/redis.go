package config

import (
	"github.com/spf13/viper"
)

type RedisConfig struct {
	RedisHost     string
	RedisPassword string
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		RedisHost:     viper.GetString("REDIS_HOST"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
	}
}
