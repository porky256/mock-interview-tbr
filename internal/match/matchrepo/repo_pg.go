package matchrepo

import (
	"database/sql"
	"time"
)

// PGMatchProvider implements GlobalDatabaseProvider
type PGMatchProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGMatchProvider creates a new postgres DB entity
func NewPGMatchProvider(db *sql.DB, timeout time.Duration) *PGMatchProvider {
	return &PGMatchProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}
