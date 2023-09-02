package controller

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
)

type OrderData struct {
	StockCode   string `json:"stock_code"`
	OrderNumber string `json:"order_number"`
	Type        string `json:"type"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}

type WorkerController struct {
	kafkaReader *kafka.Reader
	redisClient *redis.Client
	log         logger.Logger
}

func NewWorkerController(kafkaReader *kafka.Reader, redisClient *redis.Client, log logger.Logger) WorkerController {
	return WorkerController{
		kafkaReader: kafkaReader,
		redisClient: redisClient,
		log:         log,
	}
}

func (c *WorkerController) Worker() {
	for {
		m, err := c.kafkaReader.ReadMessage(context.Background())
		if err != nil {
			c.log.Errorf("error while receiving message : %s", err.Error())
		}

		data := OrderData{}
		if err = json.Unmarshal(m.Value, &data); err != nil {
			c.log.Errorf("error while receiving message : %s", err.Error())
		}

		c.log.Info(data)
	}
}
