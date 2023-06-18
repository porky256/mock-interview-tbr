package pg

import (
	"database/sql"
	"fmt"
	"github.com/porky256/mock-interview-tbr/internal/database"

	// driver for postgresql
	_ "github.com/lib/pq"
)

// DB wrapper for sql.DB
type DB struct {
	DB *sql.DB
}

// ConnectPGSQL establishes connection to DB
func ConnectPGSQL(config database.DBConfig) (*DB, error) {
	db, err := sql.Open(config.DriverName, buildPGConnString(config))

	if err != nil {
		return nil, fmt.Errorf("error occurred while connecting to DB: %w", err)
	}

	return &DB{DB: db}, nil
}

// buildPGConnString forms conn string from config
func buildPGConnString(c database.DBConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		c.User, c.Password, c.Name, c.Host, c.Port, c.SSLMode)
}
