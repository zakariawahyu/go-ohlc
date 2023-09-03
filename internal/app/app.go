package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/internal/gateway"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
)

func NewApp(cfg *config.Config, log logger.Logger) {
	e := echo.New()

	controller := gateway.NewOrderServiceClient(cfg, log)
	e.GET("/order", controller.Order)

	summaryService := gateway.NewSummaryServiceClient(cfg, log)
	e.GET("/summary/:stock-code", summaryService.GetSummary)

	log.Fatal(e.Start(fmt.Sprintf(":%v", cfg.App.ClientPort)))
}
