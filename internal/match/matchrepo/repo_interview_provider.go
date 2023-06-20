package matchrepo

import (
	"context"
	"fmt"

	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
)

// InsertInterview inserts a interview into the database
func (db *PGMatchProvider) InsertInterview(interview repomodels.InterviewRepo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `INSERT INTO  interviews (match_id, status, interview_date) 
			 VALUES ($1, $2, $3)`

	res, err := db.DB.ExecContext(ctx, stmt,
		interview.MatchID,
		interview.Status,
		interview.InterviewDate,
	)

	if err != nil {
		return 0, fmt.Errorf("error with InsertInterview query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error with InsertInterview query: %w", err)
	}

	if rowsAffected == 0 {
		return 0, database.ErrNoRowsInserted
	}

	newID, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error with InsertInterview query: %w", err)
	}

	return int(newID), nil
}

// GetInterviewByID scans for interview by selected id
func (db *PGMatchProvider) GetInterviewByID(id int) (*repomodels.InterviewRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, match_id, status, interview_date, created_at, updated_at 
				FROM interviews WHERE id=$1`
	interview := new(repomodels.InterviewRepo)
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&interview.ID,
		&interview.MatchID,
		&interview.Status,
		&interview.InterviewDate,
		&interview.CreatedAt,
		&interview.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetInterviewByID query: %w", err)
	}

	return interview, nil
}

// UpdateInterview updates interview with selected id
func (db *PGMatchProvider) UpdateInterview(interview repomodels.InterviewRepo) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE interviews SET match_id = $1, status = $2, interview_date = $3 WHERE id=$4
			`
	res, err := db.DB.ExecContext(ctx, stmt,
		interview.MatchID,
		interview.Status,
		interview.InterviewDate,
		interview.ID,
	)

	if err != nil {
		return fmt.Errorf("error with UpdateInterview query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with UpdateInterview query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

// DeleteInterviewByID deletes interview with selected id
func (db *PGMatchProvider) DeleteInterviewByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM interviews WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteInterviewByID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteInterviewByID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
