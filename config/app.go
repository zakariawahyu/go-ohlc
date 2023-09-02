package config

import "github.com/spf13/viper"

type AppConfig struct {
	ClientPort         string
	SummaryServicePort string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		ClientPort:         viper.GetString("APP_CLIENT_PORT"),
		SummaryServicePort: viper.GetString("APP_SERVICE_SUMMARY_PORT"),
	}
}
