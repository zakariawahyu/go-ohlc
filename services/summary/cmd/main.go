package main

import (
	"fmt"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/pb"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"github.com/zakariawahyu/go-ohlc/pkg/redis"
	"github.com/zakariawahyu/go-ohlc/services/summary/server/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.NewConfig()
	appLogger := logger.NewLogger(cfg)
	appLogger.InitLogger()

	redisClient, err := redis.InitRedis(cfg)
	if err != nil {
		appLogger.Fatalf("Can't connect to redis %v", err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.App.SummaryServicePort))
	if err != nil {
		appLogger.Fatalf("Service could not listen %v", err)
	}

	grpcServer := grpc.NewServer()
	summaryServer := service.SummaryService{
		RedisClient: redisClient,
		Log:         appLogger,
	}

	pb.RegisterSummaryServiceServer(grpcServer, &summaryServer)
	appLogger.Fatal(grpcServer.Serve(listen))
}
