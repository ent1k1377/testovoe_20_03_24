package model

type Product struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Price         int32  `json:"price"`
	StockQuantity int32  `json:"stock_quantity"`
}
