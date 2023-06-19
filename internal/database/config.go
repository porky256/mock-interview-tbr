package database

import (
	"errors"
	"time"
)

var ErrNoRowsDeleted = errors.New("no rows was deleted")

var ErrNoRowsUpdated = errors.New("no rows was updated")

var ErrNoRowsInserted = errors.New("no rows was inserted")

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
