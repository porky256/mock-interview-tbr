package userrepo

import (
	"context"
	"fmt"
	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
)

// InsertUserSkill inserts a user's skill into the database
func (db *PGUserProvider) InsertUserSkill(skill repomodels.UserSkillRepo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `INSERT INTO  users_skills (skill_id, user_id, score) 
			 VALUES
			 ($1, $2, $3)`

	res, err := db.DB.ExecContext(ctx, stmt,
		skill.SkillID,
		skill.UserID,
		skill.Score,
	)

	if err != nil {
		return 0, fmt.Errorf("error with InsertUserSkill query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("error with InsertUserSkill query: %w", err)
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

// GetUserSkillByID scans for user's skill with selected id
func (db *PGUserProvider) GetUserSkillByID(id int) (*repomodels.UserSkillRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	skill := new(repomodels.UserSkillRepo)
	stmt := `SELECT id, skill_id, user_id, score, created_at, updated_at  FROM users_skills WHERE id=$1`
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&skill.ID,
		&skill.SkillID,
		&skill.UserID,
		&skill.Score,
		&skill.CreatedAt,
		&skill.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error with GetUserSkillByID query: %w", err)
	}

	return skill, nil
}

// GetUsersSkillsByUserID scans for all user's skills with selected user id
func (db *PGUserProvider) GetUsersSkillsByUserID(userID int) ([]repomodels.UserSkillRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	skills := make([]repomodels.UserSkillRepo, 0)
	stmt := `SELECT id, skill_id, user_id, score, created_at, updated_at  FROM users_skills WHERE user_id=$1`

	rows, err := db.DB.QueryContext(ctx, stmt, userID)

	if err != nil {
		return nil, fmt.Errorf("error with GetUsersSkillsByUserID query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var skill repomodels.UserSkillRepo
		err := rows.Scan(
			&skill.ID,
			&skill.SkillID,
			&skill.UserID,
			&skill.Score,
			&skill.CreatedAt,
			&skill.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error with GetUsersSkillsByUserID query: %w", err)
		}

		skills = append(skills, skill)
	}

	if err != nil {
		return nil, fmt.Errorf("error with GetUsersSkillsByUserID query: %w", err)
	}

	return skills, nil
}

// UpdateUserSkill updates a user's skill with selected id
func (db *PGUserProvider) UpdateUserSkill(skill repomodels.UserSkillRepo) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE users_skills SET skill_id = $1, user_id = $2 score = $3 WHERE id=$4`

	res, err := db.DB.ExecContext(ctx, stmt,
		skill.SkillID,
		skill.UserID,
		skill.Score,
		skill.ID,
	)

	if err != nil {
		return fmt.Errorf("error with UpdateUserSkill query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with UpdateUserSkill query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

// DeleteUserSkillByID deletes user's skill with selected id
func (db *PGUserProvider) DeleteUserSkillByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM users_skills WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteUserSkillByID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteUserSkillByID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}

// DeleteUserSkillByUserID deletes user's skill with selected user id
func (db *PGUserProvider) DeleteUserSkillByUserID(userID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM users_skills WHERE user_id=$1", userID)

	if err != nil {
		return fmt.Errorf("error with DeleteUserSkillByUserID query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error with DeleteUserSkillByUserID query: %w", err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
