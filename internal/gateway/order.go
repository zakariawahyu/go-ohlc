package gateway

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-ohlc/config"
	"github.com/zakariawahyu/go-ohlc/entity"
	"github.com/zakariawahyu/go-ohlc/pb"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers"
	"github.com/zakariawahyu/go-ohlc/pkg/helpers/response"
	"github.com/zakariawahyu/go-ohlc/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

var validate = validator.New()

type OrderServiceClient struct {
	client pb.OrderServiceClient
	log    logger.Logger
}

func NewOrderServiceClient(cfg *config.Config, log logger.Logger) OrderServiceClient {
	conn, err := grpc.Dial(cfg.App.OrderServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect port %s : %v", cfg.App.OrderServicePort, err)
	}

	return OrderServiceClient{
		client: pb.NewOrderServiceClient(conn),
		log:    log,
	}
}

func (o *OrderServiceClient) Order(c echo.Context) error {
	order := entity.Order{}

	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err))
	}

	if err := validate.Struct(&order); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
	}

	res, err := o.client.Order(context.Background(), &pb.OrderRequest{
		StockCode:   order.StockCode,
		OrderNumber: order.OrderNumber,
		Type:        order.Type,
		Quantity:    order.Quantity,
		Price:       order.Price,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, err))
	}

	return c.JSON(http.StatusCreated, response.NewSuccessResponse(http.StatusCreated, res))
}

func (o *OrderServiceClient) BulkOrder(c echo.Context) error {
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
