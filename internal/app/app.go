package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/internal/gateway"
	middleware2 "github.com/zakariawahyu/go-ohlc/internal/middleware"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
)

func NewApp(cfg *config.Config, log logger.Logger) {
	e := echo.New()
	mw := middleware2.NewMiddlewareLogger(cfg, log)
	e.Use(mw.LoggerMiddleware)
	e.Use(middleware.RequestID())

	orderService := gateway.NewOrderServiceClient(cfg, log)
	e.GET("/bulk-order", orderService.BulkOrder)
	e.POST("/order", orderService.Order)

	summaryService := gateway.NewSummaryServiceClient(cfg, log)
	e.GET("/summary/:stock-code", summaryService.GetSummary)

	log.Fatal(e.Start(fmt.Sprintf(":%v", cfg.App.ClientPort)))
}
