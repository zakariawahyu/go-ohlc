package kafka_client

import (
	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(kafkaURL []string, groupTopics string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		GroupTopics:            []string{groupTopics},
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		MaxWait:                maxWait,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	})
}
