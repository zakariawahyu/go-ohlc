package helpers

import (
	"github.com/zakariawahyu/go-ohlc/entity"
	"strconv"
)

func CalculateOHLC(ohlc entity.OHLC, data entity.Order) entity.OHLC {
	price, _ := strconv.ParseInt(data.Price, 10, 64)
	quantity, _ := strconv.ParseInt(data.Quantity, 10, 64)

	if ohlc.StockCode == "" && data.Type != "A" {
		ohlc = entity.OHLC{
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
