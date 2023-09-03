package helpers

import (
	"strconv"
)

type Order struct {
	StockCode   string `json:"stock_code"`
	OrderNumber string `json:"order_number"`
	Type        string `json:"type"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}

type OHLC struct {
	StockCode     string
	PreviousPrice int64
	OpenPrice     int64
	HighestPrice  int64
	LowestPrice   int64
	ClosePrice    int64
	AveragePrice  int64
	Volume        int64
	Value         int64
}

func CalculateOHLC(ohlc OHLC, data Order) OHLC {
	price, _ := strconv.ParseInt(data.Price, 10, 64)
	quantity, _ := strconv.ParseInt(data.Quantity, 10, 64)

	// Check if no record before
	if ohlc.StockCode == "" && data.Type != "A" {
		ohlc = OHLC{
			StockCode:     data.StockCode,
			PreviousPrice: price,
			OpenPrice:     price,
		}
	}

	if data.Type == "E" || data.Type == "P" {
		ohlc.Volume += quantity
		ohlc.Value += price * quantity
		ohlc.AveragePrice = ohlc.Value / ohlc.Volume

		if price > ohlc.HighestPrice {
			ohlc.HighestPrice = price
		}

		if price < ohlc.LowestPrice || ohlc.LowestPrice == 0 {
			ohlc.LowestPrice = price
		}

		if quantity > 0 {
			ohlc.ClosePrice = price
		}
	}

	if data.Type == "A" && quantity <= 1 {
		ohlc.PreviousPrice = price
	}

	return ohlc
}
