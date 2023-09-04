package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-ohlc/pb"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	redis2 "github.com/zakariawahyu/go-ohlc/pkg/redis"
)

type SummaryService struct {
	RedisClient *redis.Client
	Log         logger.Logger
	pb.UnimplementedSummaryServiceServer
}

func (s *SummaryService) Summary(ctx context.Context, req *pb.SummaryRequest) (*pb.SummaryResponse, error) {
	stockCode := req.GetStockCode()

	result, err := redis2.GetRedis(s.RedisClient, ctx, stockCode)
	if err != nil {
		s.Log.Errorf("Error get redis %v", err)
		return nil, err
	}

	return &pb.SummaryResponse{
		StockCode:     result.StockCode,
		PreviousPrice: result.PreviousPrice,
		OpenPrice:     result.OpenPrice,
		HighestPrice:  result.HighestPrice,
		LowestPrice:   result.LowestPrice,
		ClosePrice:    result.ClosePrice,
		AveragePrice:  result.AveragePrice,
		Volume:        result.Volume,
		Value:         result.Value,
	}, nil
}
