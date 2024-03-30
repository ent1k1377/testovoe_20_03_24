package postgres

import (
	"fmt"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/model"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/postgres/util"
)

type ProductsShelve struct {
	repo *Repository
}

func GetProductsShelvesIds(products []model.Product) []int64 {
	productsIds := make([]int64, len(products))
	for i, v := range products {
		productsIds[i] = v.ID
	}

	return productsIds
}

func (r *ProductsShelve) GetProductsShelves(productsIds []int64) ([]model.ProductsShelf, error) {
	query := "select * from products_shelves where product_id in (%s)"

	rows, err := r.repo.db.Query(util.GetQuery(query, len(productsIds)), util.ConvertInt64SliceToInterfaceSlice(productsIds)...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	productsShelves := make([]model.ProductsShelf, 0)
	for rows.Next() {
		var productsShelf model.ProductsShelf

		err := rows.Scan(&productsShelf.ID, &productsShelf.ProductID, &productsShelf.ShelvesID, &productsShelf.IsPrimary)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		productsShelves = append(productsShelves, productsShelf)
	}

	return productsShelves, nil
}
