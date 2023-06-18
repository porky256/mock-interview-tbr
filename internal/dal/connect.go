package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type DBConfig struct {
	DriverName    string
	MaxOpenDbConn int
	MaxIdleDbConn int
	MaxDbLifetime time.Duration
	MaxDbIdletime time.Duration
	Host          string
	Port          string
	Name          string
	User          string
	Password      string
	SSLMode       string
}

func ConnectSQL(config DBConfig) (*sql.DB, error) {
	return sql.Open(config.DriverName, buildConnString(config))
}

func buildConnString(c DBConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		c.User, c.Password, c.Name, c.Host, c.Port, c.SSLMode)
}
