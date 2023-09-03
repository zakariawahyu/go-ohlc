package gateway

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/pb"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers/response"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type OrderServiceClient struct {
	client pb.OrderServiceClient
	log    logger.Logger
}

func NewOrderServiceClient(cfg *config.Config, log logger.Logger) OrderServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf(":%v", cfg.App.OrderServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect port %s : %v", cfg.App.OrderServicePort, err)
	}

	return OrderServiceClient{
		client: pb.NewOrderServiceClient(conn),
		log:    log,
	}
}

func (o *OrderServiceClient) Order(c echo.Context) error {
	directory := "./subsetdata/"
	ndjsonData, err := helpers.LoadFiles(directory)
	if err != nil {
		o.log.Errorf("Error load files %v", err)
	}

	for _, data := range ndjsonData {
		order := &pb.OrderRequest{
			StockCode:   data.StockCode,
			OrderNumber: data.OrderNumber,
			Type:        data.Type,
			Quantity:    data.Quantity,
			Price:       data.Price,
		}

		o.client.Order(context.Background(), order)
	}

	return c.JSON(http.StatusOK, response.NewSuccessResponse(http.StatusOK, "Success push to kafka"))
}
