package postgres

import (
	"database/sql"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository"
)

type Repository struct {
	db                  *sql.DB
	orderItemRepository repository.OrderItemRepository
	productRepository   repository.ProductRepository
	productShelf        repository.ProductsShelf
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

func (r *Repository) ProductShelf() repository.ProductsShelf {
	if r.productShelf != nil {
		return r.productShelf
	}

	r.productShelf = &ProductsShelve{repo: r}

	return r.productShelf
}
