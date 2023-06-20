package userrepo

import (
	"context"
	"fmt"

	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
)

// InsertUser inserts a user into the database
func (db *PGUserProvider) InsertUser(user repomodels.UserRepo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `INSERT INTO  users (username, first_name, last_name, email, password, phone, user_status, description) 
			 VALUES
			 ($1, $2, $3, $4, $5, $6, $7,$8)`

	res, err := db.DB.ExecContext(ctx, stmt,
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
		return 0, fmt.Errorf("error with InsertUser query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error with InsertUser query: %w", err)
	}

	if rowsAffected == 0 {
		return 0, database.ErrNoRowsInserted
	}

	newID, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error with InsertUser query: %w", err)
	}

	return int(newID), nil
}

// GetUserByID scans for user with selected id
func (db *PGUserProvider) GetUserByID(id int) (*repomodels.UserRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, username, first_name, last_name, email, password, phone, 
        user_status, description, created_at, updated_at FROM users WHERE id=$1`
	user := new(repomodels.UserRepo)
	row := db.DB.QueryRowContext(ctx, stmt, id)
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

// GetUserByUsername scans for user with selected username
func (db *PGUserProvider) GetUserByUsername(username string) (*repomodels.UserRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	user := new(repomodels.UserRepo)
	stmt := `SELECT id, username, first_name, last_name, email, password, phone, 
        user_status, description, created_at, updated_at FROM users WHERE username=$1`
	row := db.DB.QueryRowContext(ctx, stmt, username)
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
		return nil, fmt.Errorf("error with GetUserByUsername query: %w", err)
	}

	return user, nil
}

// UpdateUser updates a user with selected id
func (db *PGUserProvider) UpdateUser(user repomodels.UserRepo) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE users SET username = $1, first_name = $2 last_name = $3, email = $4, 
                              password = $5, phone = $6, user_status=$7, description = $8 WHERE id=$9
			`
	res, err := db.DB.ExecContext(ctx, stmt,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Phone,
		user.UserStatus,
		user.Description,
		user.ID,
	)

	if err != nil {
		return fmt.Errorf("error with UpdateUser query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with UpdateUser query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

// DeleteUserByID deletes user with selected id
func (db *PGUserProvider) DeleteUserByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteUserByID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteUserByID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
