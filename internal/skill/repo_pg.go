package skill

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
	"time"
)

// PGSkillProvider implements GlobalDatabaseProvider
type PGSkillProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGSkillProvider creates a new postgres DB entity
func NewPGSkillProvider(db *sql.DB, timeout time.Duration) *PGSkillProvider {
	return &PGSkillProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}

// InsertSkill inserts a skill into the database
func (db *PGSkillProvider) InsertSkill(skill repomodels.SkillRepo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `INSERT INTO  skills (name, description) VALUES ($1, $2)`

	res, err := db.DB.ExecContext(ctx, stmt,
		skill.Name,
		skill.Description,
	)

	if err != nil {
		return 0, fmt.Errorf("error with InsertSkill query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error with InsertSkill query: %w", err)
	}

	if rowsAffected == 0 {
		return 0, database.ErrNoRowsInserted
	}

	newID, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error with InsertSkill query: %w", err)
	}

	return int(newID), nil
}

// GetSkillByID scans for skill with selected id
func (db *PGSkillProvider) GetSkillByID(id int) (*repomodels.SkillRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, name, description, created_at, updated_at FROM skills WHERE id=$1`
	skill := new(repomodels.SkillRepo)
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&skill.ID,
		&skill.Name,
		&skill.Description,
		&skill.CreatedAt,
		&skill.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetSkillByID query: %w", err)
	}

	return skill, nil
}

// GetSkillByName scans for skill with selected name
func (db *PGSkillProvider) GetSkillByName(name string) (*repomodels.SkillRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	skill := new(repomodels.SkillRepo)
	stmt := `SELECT id, name, description, created_at, updated_at FROM skills WHERE name=$1`
	row := db.DB.QueryRowContext(ctx, stmt, name)
	err := row.Scan(
		&skill.ID,
		&skill.Name,
		&skill.Description,
		&skill.CreatedAt,
		&skill.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetSkillByName query: %w", err)
	}

	return skill, nil
}

// UpdateSkill updates a skill with selected id
func (db *PGSkillProvider) UpdateSkill(skill repomodels.SkillRepo) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE skills SET name = $1, description = $8 WHERE id=$9
			`
	res, err := db.DB.ExecContext(ctx, stmt,
		skill.Name,
		skill.Description,
		skill.ID,
	)

	if err != nil {
		return fmt.Errorf("error with UpdateSkill query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with UpdateSkill query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

// DeleteSkillByID deletes skill with selected id
func (db *PGSkillProvider) DeleteSkillByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM skills WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteSkillByID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteSkillByID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
