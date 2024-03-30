package model

type OrderItem struct {
	ID        int64 `json:"id"`
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}
