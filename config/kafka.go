package config

import "github.com/spf13/viper"

type KafkaConfig struct {
	KafkaBroker    string
	KafkaGroupID   string
	KafkaTopic     string
	KafkaPartition int
	KafkaReplica   int
}

func LoadKafkaConfig() KafkaConfig {
	return KafkaConfig{
		KafkaBroker:    viper.GetString("KAFKA_BROKER"),
		KafkaGroupID:   viper.GetString("KAFKA_GROUP_ID"),
		KafkaTopic:     viper.GetString("KAFKA_TOPIC"),
		KafkaPartition: viper.GetInt("KAFKA_PARTITION"),
		KafkaReplica:   viper.GetInt("KAFKA_REPLICA"),
	}
}
