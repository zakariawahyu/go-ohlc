package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/entity"
	"github.com/zakariawahyu/go-ohlc/pb"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers/response"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
)

type SummaryServiceClient struct {
	client pb.SummaryServiceClient
	log    logger.Logger
}

func NewSummaryServiceClient(cfg *config.Config, log logger.Logger) SummaryServiceClient {
	conn, err := grpc.Dial(cfg.App.SummaryServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect port %s : %v", cfg.App.SummaryServicePort, err)
	}

	return SummaryServiceClient{
		client: pb.NewSummaryServiceClient(conn),
		log:    log,
	}
}

func (s *SummaryServiceClient) GetSummary(c echo.Context) error {
	stockCode := c.Param("stock-code")

	res, err := s.client.Summary(context.Background(), &pb.SummaryRequest{
		StockCode: strings.ToUpper(stockCode),
	})

	if err != nil {
		s.log.Errorf("gateway.GetSummary %v", err.Error())
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	ohlc := entity.OHLC{
		StockCode:     res.StockCode,
		PreviousPrice: res.PreviousPrice,
		OpenPrice:     res.OpenPrice,
		HighestPrice:  res.HighestPrice,
		LowestPrice:   res.LowestPrice,
		ClosePrice:    res.ClosePrice,
		AveragePrice:  res.AveragePrice,
		Volume:        res.Volume,
		Value:         res.Value,
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, ohlc))
}
