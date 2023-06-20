package matchrepo

import (
	"context"
	"fmt"

	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
)

// InsertSkillInRequest inserts a skill related to request into the database
func (db *PGMatchProvider) InsertSkillInRequest(skill repomodels.SkillInRequestRepo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `INSERT INTO  skills_in_requests (request_id, skill_id) 
			 VALUES ($1, $2)`

	res, err := db.DB.ExecContext(ctx, stmt,
		skill.RequestID,
		skill.SkillID,
	)

	if err != nil {
		return 0, fmt.Errorf("error with InsertSkillInRequest query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error with InsertSkillInRequest query: %w", err)
	}

	if rowsAffected == 0 {
		return 0, database.ErrNoRowsInserted
	}

	newID, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error with InsertSkillInRequest query: %w", err)
	}

	return int(newID), nil
}

// GetSkillInRequestByID scans for a skill related to request by selected id
func (db *PGMatchProvider) GetSkillInRequestByID(id int) (*repomodels.SkillInRequestRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, request_id, skill_id, created_at, updated_at 
				FROM skills_in_requests WHERE id=$1`
	skill := new(repomodels.SkillInRequestRepo)
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&skill.ID,
		&skill.RequestID,
		&skill.SkillID,
		&skill.CreatedAt,
		&skill.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetSkillInRequestByID query: %w", err)
	}

	return skill, nil
}

// GetSkillInRequestByRequestID scans for a skills related to request by selected request's ID
func (db *PGMatchProvider) GetSkillInRequestByRequestID(requestID int) (*repomodels.SkillInRequestRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	match := new(repomodels.SkillInRequestRepo)
	stmt := `SELECT id, request_id, skill_id, created_at, updated_at 
				FROM skills_in_requests WHERE request_id=$1`
	row := db.DB.QueryRowContext(ctx, stmt, requestID)
	err := row.Scan(
		&match.ID,
		&match.RequestID,
		&match.SkillID,
		&match.CreatedAt,
		&match.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetSkillInRequestByRequestID query: %w", err)
	}

	return match, nil
}

// DeleteSkillInRequestByID deletes skill related to request with selected id
func (db *PGMatchProvider) DeleteSkillInRequestByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM skills_in_requests WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteSkillInRequestByID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteSkillInRequestByID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}

// DeleteSkillInRequestByRequestID deletes skills related to request with selected request id
func (db *PGMatchProvider) DeleteSkillInRequestByRequestID(requestID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM skills_in_requests WHERE request_id=$1", requestID)

	if err != nil {
		return fmt.Errorf("error with DeleteSkillInRequestByRequestID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteSkillInRequestByRequestID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
