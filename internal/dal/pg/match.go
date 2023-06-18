package pg

import (
	dalmodel "github.com/porky256/mock-interview-tbr/internal/dal/models"

	"context"
	"fmt"
)

// GetUserMatchByID scans for user's match with selected id
func (db *PostgresDB) GetUserMatchByID(id int) (*dalmodel.UserMatch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	match := new(dalmodel.UserMatch)
	row := db.DB.QueryRowContext(ctx, "SELECT * FROM users_matches WHERE id=$1", id)
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
