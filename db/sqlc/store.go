package db

import "database/sql"

type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new SQLStore
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}
