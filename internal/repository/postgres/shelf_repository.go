package postgres

import (
	"fmt"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/model"
)

type ShelfRepository struct {
	repo *Repository
}

func (r *ShelfRepository) GetAllShelves() ([]model.Shelf, error) {
	query := "select * from shelves"

	rows, err := r.repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var shelves []model.Shelf
	for rows.Next() {
		var shelf model.Shelf
		err := rows.Scan(&shelf.ID, &shelf.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		shelves = append(shelves, shelf)
	}

	return shelves, nil
}
