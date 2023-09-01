package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Redis RedisConfig
}

func NewConfig() *Config {
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error can't load config")
	}

	return &Config{
		Redis: LoadRedisConfig(),
	}
}
