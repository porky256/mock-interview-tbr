package pg

import (
	"database/sql"
	"time"
)

// PostgresDB implements GlobalDatabaseProvider
type PostgresDB struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPostgresDB creates a new postgres DB entity
func NewPostgresDB(db *sql.DB, timeout time.Duration) PostgresDB {
	return PostgresDB{
		DB:           db,
		QueryTimeout: timeout,
	}
}
