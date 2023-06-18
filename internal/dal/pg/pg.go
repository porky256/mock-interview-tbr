package pg

import (
	"database/sql"
	"github.com/porky256/mock-interview-tbr/internal/dal"
)

type postgresDB struct {
	DB *sql.DB
}

// NewPostgresDB creates a new postgres DB entity
func NewPostgresDB(db *sql.DB) dal.GlobalDatabaseProvider {
	return &postgresDB{
		DB: db,
	}
}
