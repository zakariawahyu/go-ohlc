package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/internal/gateway"
	kafka_client "github.com/zakariawahyu/go-ohlc/pkg/kafka-client"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
)

func NewApp(cfg *config.Config, log logger.Logger) {
	e := echo.New()
	ctx := context.Background()

	// Connect Kafka Brokers
	if err := kafka_client.ConnectKafkaBrokers(ctx, cfg, log); err != nil {
		log.Fatal(err)
	}

	// Kafka Producer
	kafkaProducer := kafka_client.NewProducer(log, cfg.Kafka.KafkaBroker)
	defer kafkaProducer.Close()

	controller := NewController(kafkaProducer, cfg.Kafka, log)
	e.GET("/order", controller.Order)

	summaryService := gateway.NewServiceClient(cfg, log)
	e.GET("/summary/:stock-code", summaryService.GetSummary)

	log.Fatal(e.Start(fmt.Sprintf(":%v", cfg.App.ClientPort)))
}
