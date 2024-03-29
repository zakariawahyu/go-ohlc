package kafka_client

import (
	"context"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
)

func ConnectKafkaBrokers(ctx context.Context, cfg *config.Config, log logger.Logger) error {
	kafkaConn, err := kafka.DialContext(ctx, "tcp", cfg.Kafka.KafkaBroker[0])
	if err != nil {
		return errors.Wrap(err, "kafka-client.ConnectKafkaBrokers.DialContext")
	}

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafka-client.ConnectKafkaBrokers.kafkaConn.Brokers")
	}
	defer kafkaConn.Close()

	InitKafkaTopics(ctx, kafkaConn, log, cfg.Kafka)

	log.Infof("(kafka-client connected) brokers: %+v", brokers)
	return nil
}
