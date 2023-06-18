package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	repomodel "github.com/porky256/mock-interview-tbr/internal/models"
)

// PGUserProvider implements GlobalDatabaseProvider
type PGUserProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGUserProvider creates a new postgres DB entity
func NewPGUserProvider(db *sql.DB, timeout time.Duration) PGUserProvider {
	return PGUserProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}

// InsertUser inserts a user into the database
func (db *PGUserProvider) InsertUser(user repomodel.UserRepo) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `insert into users (username, first_name, last_name, email, password, phone, user_status, description) 
			values
			($1, $2, $3, $4, $5, $6, $7,$8)`

	_, err := db.DB.ExecContext(ctx, stmt,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Phone,
		user.UserStatus,
		user.Description,
	)

	if err != nil {
		return fmt.Errorf("error with InsertUser query: %w", err)
	}

	return nil
}

// GetUserByID scans for user with selected id
func (db *PGUserProvider) GetUserByID(id int) (*repomodel.UserRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	user := new(repomodel.UserRepo)
	row := db.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE id=$1", id)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.UserStatus,
		&user.Description,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetUserByID query: %w", err)
	}

	return user, nil
}
