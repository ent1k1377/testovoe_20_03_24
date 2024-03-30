package model

type ProductsShelf struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	ShelvesID int64 `json:"shelves_id"`
	IsPrimary bool  `json:"is_primary"`
}
