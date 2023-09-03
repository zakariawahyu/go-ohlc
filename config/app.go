package config

import "github.com/spf13/viper"

type AppConfig struct {
	ClientPort         string
	OrderServicePort   string
	SummaryServicePort string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		ClientPort:         viper.GetString("APP_CLIENT_PORT"),
		OrderServicePort:   viper.GetString("APP_SERVICE_ORDER_PORT"),
		SummaryServicePort: viper.GetString("APP_SERVICE_SUMMARY_PORT"),
	}
}
