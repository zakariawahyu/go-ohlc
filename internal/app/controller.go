package app

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers/response"
	kafka_client "github.com/zakariawahyu/go-ohlc/pkg/kafka-client"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"net/http"
	"time"
)

type Controller struct {
	producer    kafka_client.Producer
	kafkaConfig config.KafkaConfig
	log         logger.Logger
}

type OrderData struct {
	StockCode   string `json:"stock_code"`
	OrderNumber string `json:"order_number"`
	Type        string `json:"type"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}

func NewController(producer kafka_client.Producer, kafkaConfig config.KafkaConfig, log logger.Logger) Controller {
	return Controller{
		producer:    producer,
		kafkaConfig: kafkaConfig,
		log:         log,
	}
}

func (ctrl *Controller) Order(c echo.Context) error {
	ctx := c.Request().Context()

	directory := "./subsetdata/"
	ndjsonData, err := helpers.LoadFiles(directory)
	if err != nil {
		ctrl.log.Errorf("Error load files %v", err)
	}

	for _, data := range ndjsonData {
		orderData := OrderData{
			StockCode:   data.StockCode,
			OrderNumber: data.OrderNumber,
			Type:        data.Type,
			Quantity:    data.Quantity,
			Price:       data.Price,
		}

		dataInBytes, _ := json.Marshal(orderData)

		ctrl.producer.PublishMessage(ctx, kafka.Message{
			Topic: ctrl.kafkaConfig.KafkaTopic,
			Value: dataInBytes,
			Time:  time.Now().UTC(),
		})
		ctrl.log.Infof("Push data to kafka %v", orderData)
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, "Success push to kafka"))
}
