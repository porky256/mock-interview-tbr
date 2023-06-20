package matchrepo

import (
	"context"
	"fmt"

	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
)

// InsertUserMatch inserts a user's match into the database
func (db *PGMatchProvider) InsertUserMatch(match repomodels.UserMatchRepo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `INSERT INTO  users_matches (user_asker, user_match, match_score) 
			 VALUES ($1, $2, $3)`

	res, err := db.DB.ExecContext(ctx, stmt,
		match.UserAsker,
		match.UserMatch,
		match.MatchScore,
	)

	if err != nil {
		return 0, fmt.Errorf("error with InsertUserMatch query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error with InsertUserMatch query: %w", err)
	}

	if rowsAffected == 0 {
		return 0, database.ErrNoRowsInserted
	}

	newID, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error with InsertUserMatch query: %w", err)
	}

	return int(newID), nil
}

// GetUserMatchByID scans for user's match by selected id
func (db *PGMatchProvider) GetUserMatchByID(id int) (*repomodels.UserMatchRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, user_asker, user_match, match_score, created_at, updated_at 
				FROM users_matches WHERE id=$1`
	match := new(repomodels.UserMatchRepo)
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&match.ID,
		&match.UserAsker,
		&match.UserMatch,
		&match.MatchScore,
		&match.CreatedAt,
		&match.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetUserMatchByID query: %w", err)
	}

	return match, nil
}

// GetUserMatchByUserAskerID scans for user's match by selected user's ID
func (db *PGMatchProvider) GetUserMatchByUserAskerID(userID int) (*repomodels.UserMatchRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	match := new(repomodels.UserMatchRepo)
	stmt := `SELECT id, user_asker, user_match, match_score, created_at, updated_at 
				FROM users_matches WHERE user_asker=$1`
	row := db.DB.QueryRowContext(ctx, stmt, userID)
	err := row.Scan(
		&match.ID,
		&match.UserAsker,
		&match.UserMatch,
		&match.MatchScore,
		&match.CreatedAt,
		&match.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetUserMatchByUserAskerID query: %w", err)
	}

	return match, nil
}

// UpdateUserMatch updates a user's match with selected id
func (db *PGMatchProvider) UpdateUserMatch(match repomodels.UserMatchRepo) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE users_matches SET user_asker = $1, user_match = $2, match_score = $3 WHERE id=$4
			`
	res, err := db.DB.ExecContext(ctx, stmt,
		match.UserAsker,
		match.UserMatch,
		match.MatchScore,
		match.ID,
	)

	if err != nil {
		return fmt.Errorf("error with UpdateUserMatch query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with UpdateUserMatch query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

// DeleteUserMatchByID deletes user's match with selected id
func (db *PGMatchProvider) DeleteUserMatchByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM users_matches WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteUserMatchByID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteUserMatchByID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
