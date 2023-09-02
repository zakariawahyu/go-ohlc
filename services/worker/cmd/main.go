package main

import (
	"github.com/zakariawahyu/go-ohlc/config"
	kafka_client "github.com/zakariawahyu/go-ohlc/pkg/kafka-client"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"github.com/zakariawahyu/go-ohlc/pkg/redis"
	controller2 "github.com/zakariawahyu/go-ohlc/services/worker/controller"
)

func main() {
	cfg := config.NewConfig()
	appLogger := logger.NewLogger(cfg)
	appLogger.InitLogger()

	redisClient, err := redis.InitRedis(cfg)
	if err != nil {
		appLogger.Errorf("Redis connection err : %v", err)
	}

	kafkaReader := kafka_client.NewKafkaReader(cfg.Kafka.KafkaBroker, cfg.Kafka.KafkaTopic, cfg.Kafka.KafkaGroupID)
	defer kafkaReader.Close()

	controller := controller2.NewWorkerController(kafkaReader, redisClient, appLogger)
	controller.Worker()
}
