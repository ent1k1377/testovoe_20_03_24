package postgres

import (
	"database/sql"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository"
)

type Repository struct {
	db                  *sql.DB
	orderItemRepository repository.OrderItemRepository
	productRepository   repository.ProductRepository
	productShelf        repository.ProductsShelfRepository
	shelf               repository.ShelfRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) OrderItem() repository.OrderItemRepository {
	if r.orderItemRepository != nil {
		return r.orderItemRepository
	}

	r.orderItemRepository = &OrderItemRepository{repo: r}

	return r.orderItemRepository
}

func (r *Repository) Product() repository.ProductRepository {
	if r.productRepository != nil {
		return r.productRepository
	}

	r.productRepository = &ProductRepository{repo: r}

	return r.productRepository
}

func (r *Repository) ProductShelf() repository.ProductsShelfRepository {
	if r.productShelf != nil {
		return r.productShelf
	}

	r.productShelf = &ProductsShelveRepository{repo: r}

	return r.productShelf
}

func (r *Repository) Shelf() repository.ShelfRepository {
	if r.productShelf != nil {
		return r.shelf
	}

	r.shelf = &ShelfRepository{repo: r}

	return r.shelf
}
