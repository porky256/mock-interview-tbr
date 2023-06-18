package dal

import (
	"database/sql"
	"fmt"
	"time"

	// driver for postgresql
	_ "github.com/lib/pq"
)

// DBConfig config to connect to database
type DBConfig struct {
	DriverName    string
	MaxOpenDBConn int
	MaxIdleDBConn int
	MaxDBLifetime time.Duration
	MaxDBIdletime time.Duration
	Host          string
	Port          string
	Name          string
	User          string
	Password      string
	SSLMode       string
}

// DB wrapper for sql.DB
type DB struct {
	DB *sql.DB
}

// ConnectSQL establishes connection to DB
func ConnectSQL(config DBConfig) (*DB, error) {
	db, err := sql.Open(config.DriverName, buildConnString(config))

	if err != nil {
		return nil, fmt.Errorf("error occurred while connecting to DB: %w", err)
	}

	return &DB{DB: db}, nil
}

// buildConnString forms conn string from config
func buildConnString(c DBConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		c.User, c.Password, c.Name, c.Host, c.Port, c.SSLMode)
}
