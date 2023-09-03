package service

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/entity"
	"github.com/zakariawahyu/go-ohlc/pb"
	kafka_client "github.com/zakariawahyu/go-ohlc/pkg/kafka-client"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"net/http"
	"time"
)

type OrderService struct {
	Producer    kafka_client.Producer
	KafkaConfig config.KafkaConfig
	Log         logger.Logger
	pb.UnimplementedOrderServiceServer
}

func (o *OrderService) Order(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	order := entity.Order{
		StockCode:   req.GetStockCode(),
		OrderNumber: req.GetOrderNumber(),
		Type:        req.GetType(),
		Quantity:    req.GetQuantity(),
		Price:       req.GetPrice(),
	}

	dataInBytes, _ := json.Marshal(order)

	if err := o.Producer.PublishMessage(ctx, kafka.Message{
		Topic: o.KafkaConfig.KafkaTopic,
		Value: dataInBytes,
		Time:  time.Now().UTC(),
	}); err != nil {
		o.Log.Errorf("Error publish data to kafka. err : %v", err)
		return nil, err
	}
	o.Log.Infof("Push data to kafka %v", order)

	return &pb.OrderResponse{
		Status:  http.StatusCreated,
		Error:   "false",
		Message: "Success created order",
	}, nil
}
