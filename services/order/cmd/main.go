package main

import (
	"context"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/pb"
	kafka_client "github.com/zakariawahyu/go-ohlc/pkg/kafka-client"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"github.com/zakariawahyu/go-ohlc/services/order/server/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.NewConfig()
	appLogger := logger.NewLogger(cfg)
	appLogger.InitLogger()
	ctx := context.Background()

	if err := kafka_client.ConnectKafkaBrokers(ctx, cfg, appLogger); err != nil {
		appLogger.Fatalf("Can't connect kafka broker %v", err)
	}

	kafkaProducer := kafka_client.NewProducer(appLogger, cfg.Kafka.KafkaBroker)
	defer kafkaProducer.Close()

	listen, err := net.Listen("tcp", cfg.App.OrderServicePort)
	if err != nil {
		appLogger.Fatalf("Service could not listen %v", err)
	}

	grpcServer := grpc.NewServer()
	orderServer := service.OrderService{
		Producer:    kafkaProducer,
		KafkaConfig: cfg.Kafka,
		Log:         appLogger,
	}

	pb.RegisterOrderServiceServer(grpcServer, &orderServer)
	appLogger.Fatal(grpcServer.Serve(listen))
}
