package postgres

import (
	"fmt"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/model"
)

var _ repository.OrderItemRepository = (*OrderItemRepository)(nil)

type OrderItemRepository struct {
	repo *Repository
}

func (r *OrderItemRepository) GetOrderItems(id int64) ([]model.OrderItem, error) {
	query := "SELECT * FROM order_items WHERE order_id = $1"

	rows, err := r.repo.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// NOTE: кол-во записей по хорошему нужно как емкаость записать
	orderItems := make([]model.OrderItem, 0)
	for rows.Next() {
		var orderItem model.OrderItem

		if err := rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}

func (r *OrderItemRepository) GetOrderCount(id int64) (int, error) {
	query := "SELECT count(id) FROM order_items WHERE order_id = ?"

	var count int
	err := r.repo.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("")
	}

	return count, nil
}
