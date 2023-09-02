package kafka_client

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"net"
	"strconv"
)

func InitKafkaTopics(ctx context.Context, kafkaConn *kafka.Conn, log logger.Logger, cfg config.KafkaConfig) {
	controller, err := kafkaConn.Controller()
	if err != nil {
		log.Error("kafka-client.InitKafkaTopics.kafkaConn.Controller err: %v", err)
		return
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	log.Infof("kafka controller uri : %s", controllerURI)

	conn, err := kafka.DialContext(ctx, "tcp", controllerURI)
	if err != nil {
		log.Errorf("kafka-client.InitKafkaTopics.DialContext err: %v", err)
		return
	}
	defer conn.Close()

	kafkaTopic := getKafkaTopicConfig(cfg)

	if err := conn.CreateTopics(kafkaTopic); err != nil {
		log.Warn("kafka-client.InitKafkaTopics.CreateTopics", err)
		return
	}

	log.Infof("kafka topics created or already exists : %+v", []kafka.TopicConfig{kafkaTopic})
}

func getKafkaTopicConfig(cfg config.KafkaConfig) kafka.TopicConfig {
	return kafka.TopicConfig{
		Topic:             cfg.KafkaTopic,
		NumPartitions:     cfg.KafkaPartition,
		ReplicationFactor: cfg.KafkaReplica,
	}
}
