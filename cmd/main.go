package main

import (
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/internal/app"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
)

func main() {
	cfg := config.NewConfig()

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()

	app.NewApp(cfg, appLogger)
}
