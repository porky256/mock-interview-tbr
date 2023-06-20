package skillrepo

import (
	"database/sql"
	"time"
)

// PGSkillProvider implements GlobalDatabaseProvider
type PGSkillProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGSkillProvider creates a new postgres DB entity
func NewPGSkillProvider(db *sql.DB, timeout time.Duration) *PGSkillProvider {
	return &PGSkillProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}
