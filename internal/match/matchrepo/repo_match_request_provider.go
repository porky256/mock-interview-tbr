package matchrepo

import (
	"context"
	"fmt"

	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
)

// InsertMatchRequest inserts a match request into the database
func (db *PGMatchProvider) InsertMatchRequest(req repomodels.MatchRequestRepo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `INSERT INTO  match_requests user_id VALUES $1`

	res, err := db.DB.ExecContext(ctx, stmt,
		req.UserID,
	)

	if err != nil {
		return 0, fmt.Errorf("error with InsertMatchRequest query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error with InsertMatchRequest query: %w", err)
	}

	if rowsAffected == 0 {
		return 0, database.ErrNoRowsInserted
	}

	newID, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error with InsertMatchRequest query: %w", err)
	}

	return int(newID), nil
}

// GetMatchRequestByID scans for match request with selected id
func (db *PGMatchProvider) GetMatchRequestByID(id int) (*repomodels.MatchRequestRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, user_id, created_at, updated_at FROM match_requests WHERE id=$1`
	req := new(repomodels.MatchRequestRepo)
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&req.ID,
		&req.UserID,
		&req.CreatedAt,
		&req.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetMatchRequestByID query: %w", err)
	}

	return req, nil
}

// GetMatchRequestByUserID scans for match requests with selected user's ID
func (db *PGMatchProvider) GetMatchRequestByUserID(userID int) (*repomodels.MatchRequestRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	req := new(repomodels.MatchRequestRepo)
	stmt := `SELECT id, user_id, created_at, updated_at FROM match_requests WHERE user_id=$1`
	row := db.DB.QueryRowContext(ctx, stmt, userID)
	err := row.Scan(
		&req.ID,
		&req.UserID,
		&req.CreatedAt,
		&req.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetMatchRequestByUserID query: %w", err)
	}

	return req, nil
}

// UpdateMatchRequest updates match request with selected id
func (db *PGMatchProvider) UpdateMatchRequest(req repomodels.MatchRequestRepo) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE match_requests SET user_id = $1 WHERE id=$2
			`
	res, err := db.DB.ExecContext(ctx, stmt,
		req.UserID,
		req.ID,
	)

	if err != nil {
		return fmt.Errorf("error with UpdateMatchRequest query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with UpdateMatchRequest query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

// DeleteMatchRequestByID deletes match request with selected id
func (db *PGMatchProvider) DeleteMatchRequestByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM match_requests WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteMatchRequestByID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteMatchRequestByID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
