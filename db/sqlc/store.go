package db

import "database/sql"

type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore создает новый экземпляр хранилища SQL.
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}
