package model

// Order представляет информацию о заказе.
type Order struct {
	OrderID     int64
	ProductID   int64
	ProductName string
	Quantity    int32
	ShelveNames []string
}

type Product struct {
	Name     string
	Quantity int32
}

type GetOrderInfoRow struct {
	ProductName     string `json:"product_name"`
	ProductID       int64  `json:"product_id"`
	OrderID         int64  `json:"order_id"`
	Quantity        int32  `json:"quantity"`
	ShelveName      string `json:"shelve_name"`
	ShelveIsPrimary bool   `json:"shelve_is_primary"`
}
