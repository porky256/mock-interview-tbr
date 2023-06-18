package dal

import "time"

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
