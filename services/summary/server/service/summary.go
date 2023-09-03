package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zakariawahyu/go-ohlc/pb"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
)

type SummaryService struct {
	RedisClient *redis.Client
	Log         logger.Logger
	pb.UnimplementedSummaryServiceServer
}

func (s *SummaryService) Summary(ctx context.Context, req *pb.SummaryRequest) (*pb.SummaryResponse, error) {
	stockCode := req.GetStockCode()

	result, err := helpers.GetRedis(s.RedisClient, ctx, stockCode)
	if err != nil {
		s.Log.Errorf("Error get redis %v", err)
		return nil, errors.Wrap(err, "Error get redis data")
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
