package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/porky256/mock-interview-tbr/internal/database"

	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
)

// PGUserProvider implements GlobalDatabaseProvider
type PGUserProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGUserProvider creates a new postgres DB entity
func NewPGUserProvider(db *sql.DB, timeout time.Duration) *PGUserProvider {
	return &PGUserProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}

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
		return nil, fmt.Errorf("error with GetUserMatchByID query: %w", err)
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
		return nil, fmt.Errorf("error with GetUserMatchByUserAskerID query: %w", err)
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

// DeleteUserByID deletes user with selected id
func (db *PGUserProvider) DeleteUserByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)

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
		return 0, fmt.Errorf("error with InsertUserMatch query: %w", err)
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
