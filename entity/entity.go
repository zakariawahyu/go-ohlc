package entity

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

type Order struct {
	StockCode   string `json:"stock_code"`
	OrderNumber string `json:"order_number"`
	Type        string `json:"type"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}

type OHLC struct {
	StockCode     string `json:"stock_code"`
	PreviousPrice int64  `json:"previous_price"`
	OpenPrice     int64  `json:"open_price"`
	HighestPrice  int64  `json:"highest_price"`
	LowestPrice   int64  `json:"lowest_price"`
	ClosePrice    int64  `json:"close_price"`
	AveragePrice  int64  `json:"average_price"`
	Volume        int64  `json:"volume"`
	Value         int64  `json:"value"`
}
