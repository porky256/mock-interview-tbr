package userrepo

import (
	"database/sql"
	"time"
)

// PGUserProvider implements GlobalDatabaseProvider
type PGUserProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGUserProvider creates a new postgres DB entity
func NewPGUserProvider(db *sql.DB, timeout time.Duration) *PGUserProvider {
	return &PGUserProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}
