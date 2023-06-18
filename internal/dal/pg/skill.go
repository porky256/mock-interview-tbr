package pg

import (
	"github.com/porky256/mock-interview-tbr/internal/dal/models"

	"context"
	"fmt"
)

// GetSkillByID scans for user with selected id
func (db *PostgresDB) GetSkillByID(id int) (*dalmodel.Skill, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	skill := new(dalmodel.Skill)
	row := db.DB.QueryRowContext(ctx, "SELECT * FROM skills WHERE id=$1", id)
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
