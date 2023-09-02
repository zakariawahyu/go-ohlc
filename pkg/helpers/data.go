package helpers

type NdjsonData struct {
	Type             string `json:"type"`
	OrderNumber      string `json:"order_number"`
	OrderVerb        string `json:"order_verb"`
	Quantity         string `json:"quantity"`
	ExecutedQuantity string `json:"executed_quantity"`
	OrderBook        string `json:"order_book"`
	Price            string `json:"price"`
	ExecutionPrice   string `json:"execution_price"`
	StockCode        string `json:"stock_code"`
}
