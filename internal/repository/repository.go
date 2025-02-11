package repository

import "github.com/ent1k1377/testovoe_20_03_24/internal/repository/model"

type Storage interface {
	OrderItem() OrderItemRepository
	Product() ProductRepository
	ProductShelf() ProductsShelfRepository
	Shelf() ShelfRepository
}

type OrderItemRepository interface {
	GetOrderItems(id int64) ([]model.OrderItem, error)
	GetOrderCount(id int64) (int, error)
}

type ProductRepository interface {
	GetProducts(orderItemsIds []int64) ([]model.Product, error)
}

type ProductsShelfRepository interface {
	GetProductsShelves(products []int64) ([]model.ProductsShelf, error)
}

type ShelfRepository interface {
	GetAllShelves() ([]model.Shelf, error)
}
