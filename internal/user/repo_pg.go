package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
	"time"
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

	var newID int

	stmt := `INSERT INTO  users (username, first_name, last_name, email, password, phone, user_status, description) 
			 VALUES
			 ($1, $2, $3, $4, $5, $6, $7,$8) RETURNING id`

	err := db.DB.QueryRowContext(ctx, stmt,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Phone,
		user.UserStatus,
		user.Description,
	).Scan(&newID)

	if err != nil {
		return 0, fmt.Errorf("error with InsertUser query: %w", err)
	}

	return newID, nil
}

// GetUserByID scans for user with selected id
func (db *PGUserProvider) GetUserByID(id int) (*repomodels.UserRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	user := new(repomodels.UserRepo)
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

// GetUserByUsername scans for user with selected username
func (db *PGUserProvider) GetUserByUsername(username string) (*repomodels.UserRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	user := new(repomodels.UserRepo)
	row := db.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE username=$1", username)
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

	_, err := db.DB.ExecContext(ctx, stmt,
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

	return nil
}

// DeleteUserByID deletes user with selected id
func (db *PGUserProvider) DeleteUserByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	_, err := db.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteUserByID query: %w", err)
	}

	return nil
}

// InsertUserSkill inserts a user's skill into the database
func (db *PGUserProvider) InsertUserSkill(skill repomodels.UserSkillRepo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	var newID int

	stmt := `INSERT INTO  users_skills (skill_id, user_id, score) 
			 VALUES
			 ($1, $2, $3) RETURNING id`

	err := db.DB.QueryRowContext(ctx, stmt,
		skill.SkillID,
		skill.UserID,
		skill.Score,
	).Scan(&newID)

	if err != nil {
		return 0, fmt.Errorf("error with InsertUserSkill query: %w", err)
	}

	return newID, nil
}

// GetUserSkillByID scans for user's skill with selected id
func (db *PGUserProvider) GetUserSkillByID(id int) (*repomodels.UserSkillRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	skill := new(repomodels.UserSkillRepo)
	row := db.DB.QueryRowContext(ctx, "SELECT * FROM users_skills WHERE id=$1", id)
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
	rows, err := db.DB.QueryContext(ctx, "SELECT * FROM users_skills WHERE user_id=$1", userID)

	if err != nil {
		return nil, fmt.Errorf("error with GetUsersSkillsByUserID query: %w", err)
	}

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

	_, err := db.DB.ExecContext(ctx, stmt,
		skill.SkillID,
		skill.UserID,
		skill.Score,
		skill.ID,
	)

	if err != nil {
		return fmt.Errorf("error with UpdateUserSkill query: %w", err)
	}

	return nil
}

// DeleteUserSkillByID deletes user's skill with selected id
func (db *PGUserProvider) DeleteUserSkillByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	_, err := db.DB.ExecContext(ctx, "DELETE FROM users_skills WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error with DeleteUserSkillByID query: %w", err)
	}

	return nil
}

// DeleteUserSkillByUserID deletes user's skill with selected user id
func (db *PGUserProvider) DeleteUserSkillByUserID(userID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	_, err := db.DB.ExecContext(ctx, "DELETE FROM users_skills WHERE user_id=$1", userID)

	if err != nil {
		return fmt.Errorf("error with DeleteUserSkillByUserID query: %w", err)
	}

	return nil
}
