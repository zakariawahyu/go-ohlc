package controller

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"github.com/zakariawahyu/go-ohlc/entity"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
)

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
	ctx := context.Background()
	for {
		m, err := c.kafkaReader.FetchMessage(ctx)
		if err != nil {
			c.log.Errorf("error while receiving message : %s", err.Error())
		}

		data := entity.Order{}
		if err = json.Unmarshal(m.Value, &data); err != nil {
			c.log.Errorf("error while receiving message : %s", err.Error())
		}

		ohlc, _ := helpers.GetRedis(c.redisClient, ctx, data.StockCode)

		calculateOHLC := helpers.CalculateOHLC(ohlc, data)
		c.log.Infof("Calculate from kafka %v", calculateOHLC)

		if err = helpers.SetRedis(c.redisClient, ctx, data.StockCode, 1800, &calculateOHLC); err != nil {
			c.log.Errorf("Error set redis %v", err.Error())
		}
	}
}
