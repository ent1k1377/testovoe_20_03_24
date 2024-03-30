package postgres

import (
	"fmt"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/model"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/postgres/util"
)

type ProductRepository struct {
	repo *Repository
}

func GetOrderItemsIds(orderItems []model.OrderItem) []int64 {
	orderItemsIds := make([]int64, len(orderItems))
	for i, v := range orderItems {
		orderItemsIds[i] = v.ProductID
	}

	return orderItemsIds
}

func (r *ProductRepository) GetProducts(orderItemsIds []int64) ([]model.Product, error) {
	query := "select * from products where id in (%s)"

	rows, err := r.repo.db.Query(util.GetQuery(query, len(orderItemsIds)), util.ConvertInt64SliceToInterfaceSlice(orderItemsIds)...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.StockQuantity)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		products = append(products, product)
	}

	return products, nil
}
