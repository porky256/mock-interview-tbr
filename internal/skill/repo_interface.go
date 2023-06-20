package skill

import "github.com/porky256/mock-interview-tbr/internal/models/repomodels"

// DatabaseSkillProvider provider for db related functions for skill
type DatabaseSkillProvider interface {
	InsertSkill(skill repomodels.SkillRepo) (int, error)
	GetSkillByID(id int) (*repomodels.SkillRepo, error)
	GetSkillByName(name string) (*repomodels.SkillRepo, error)
	UpdateSkill(skill repomodels.SkillRepo) error
	DeleteSkillByID(id int) error
}
