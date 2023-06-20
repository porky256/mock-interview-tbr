package repomodels

import "time"

type SkillRepo struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
