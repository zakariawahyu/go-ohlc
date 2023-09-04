package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	ClientPort         string
	OrderServicePort   string
	SummaryServicePort string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		ClientPort:         fmt.Sprintf(":%v", viper.GetString("APP_CLIENT_PORT")),
		OrderServicePort:   fmt.Sprintf(":%v", viper.GetString("APP_SERVICE_ORDER_PORT")),
		SummaryServicePort: fmt.Sprintf(":%v", viper.GetString("APP_SERVICE_SUMMARY_PORT")),
	}
}
